package repositories

import (
	"context"
	"database/sql"

	"github.com/C-dexTeam/codex/internal/domains"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AttributesRepository struct {
	db *sqlx.DB
}

type dbModelAttribute struct {
	ID        sql.NullString `db:"id"`
	RewardID  sql.NullString `db:"reward_id"`
	TraitType sql.NullString `db:"trait_type"`
	Value     sql.NullString `db:"value"`
}

func (r *AttributesRepository) dbModelToAppModel(dbModel dbModelAttribute) (attribute domains.Attribute) {
	attribute.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		uuid.MustParse(dbModel.RewardID.String),
		dbModel.TraitType.String,
		dbModel.Value.String,
	)

	return
}

func (r *AttributesRepository) dbModelFromAppFilter(filter domains.AttributeFilter) (dbModel dbModelAttribute) {
	if filter.ID != uuid.Nil {
		dbModel.ID.String = filter.ID.String()
		dbModel.ID.Valid = true
	}
	if filter.RewardID != uuid.Nil {
		dbModel.RewardID.String = filter.RewardID.String()
		dbModel.RewardID.Valid = true
	}
	if filter.TraitType != "" {
		dbModel.TraitType.String = filter.TraitType
		dbModel.TraitType.Valid = true
	}

	return
}

func NewAttributesRepository(db *sqlx.DB) domains.IAttributeRepository {
	return &AttributesRepository{db: db}
}

func (r *AttributesRepository) Filter(ctx context.Context, filter domains.AttributeFilter, limit, page int64) (rewards []domains.Attribute, dataCount int64, err error) {
	dbFilter := r.dbModelFromAppFilter(filter)
	dbResult := []dbModelAttribute{}

	query := `
	SELECT
		r.*
	FROM 
		t_attributes r
	WHERE
		($1::uuid IS NULL OR r.id = $1::uuid) AND
		($2::uuid IS NULL OR r.reward_id = $2::uuid) AND
		($3::text IS NULL OR r.trait_type LIKE '%' || $3::text || '%') AND
		($4::text IS NULL OR r.value LIKE '%' || $4::text || '%')
	LIMIT $5 OFFSET $6
	`

	// Execute the query with the extracted fields
	err = r.db.SelectContext(ctx, &dbResult, query, dbFilter.ID, dbFilter.RewardID, dbFilter.TraitType, dbFilter.Value, limit, (page-1)*limit)
	if err != nil {
		return
	}
	for _, dbModel := range dbResult {
		rewards = append(rewards, r.dbModelToAppModel(dbModel))
	}
	return
}
