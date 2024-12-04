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
	var rewardID *uuid.UUID

	if parsedRewardID, err := uuid.Parse(dbModel.RewardID.String); err == nil {
		rewardID = &parsedRewardID
	} else {
		rewardID = nil
	}

	attribute.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		rewardID,
		dbModel.TraitType.String,
		dbModel.Value.String,
	)

	return
}

func (r *AttributesRepository) dbModelFromAppModel(appModel domains.Attribute) (dbModel dbModelAttribute) {
	if appModel.GetID() != uuid.Nil {
		dbModel.ID.String = appModel.GetID().String()
		dbModel.ID.Valid = true
	}
	if appModel.GetRewardID() != nil {
		dbModel.RewardID.String = appModel.GetRewardID().String()
		dbModel.RewardID.Valid = true
	}
	if appModel.GetTraitType() != "" {
		dbModel.TraitType.String = appModel.GetTraitType()
		dbModel.TraitType.Valid = true
	}
	if appModel.GetValue() != "" {
		dbModel.Value.String = appModel.GetValue()
		dbModel.Value.Valid = true
	}

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

func (r *AttributesRepository) Filter(ctx context.Context, filter domains.AttributeFilter, limit, page int64) (attributes []domains.Attribute, dataCount int64, err error) {
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
		attributes = append(attributes, r.dbModelToAppModel(dbModel))
	}
	return
}

func (r *AttributesRepository) Add(ctx context.Context, attribute *domains.Attribute) (uuid.UUID, error) {
	dbModel := r.dbModelFromAppModel(*attribute)
	query := `
		INSERT INTO
			t_attributes (reward_id, trait_type, value)
		VALUES
			($1, $2, $3)
		RETURNING id
	`

	var id uuid.UUID
	err := r.db.QueryRowxContext(ctx, query, dbModel.RewardID, dbModel.TraitType, dbModel.Value).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *AttributesRepository) Update(ctx context.Context, attribute *domains.Attribute) (err error) {
	dbModel := r.dbModelFromAppModel(*attribute)
	query := `
		UPDATE
			t_attributes
		SET
			reward_id = COALESCE(:reward_id, reward_id),
			trait_type = COALESCE(:trait_type, trait_type),
			value = COALESCE(:value, value)
		WHERE
			id = :id
	`
	_, err = r.db.NamedExecContext(ctx, query, dbModel)
	if err != nil {
		return
	}
	return
}

func (r *AttributesRepository) Delete(ctx context.Context, attributeID uuid.UUID) (err error) {
	query := `
		DELETE FROM
			t_attributes
		WHERE 
			id = $1
	`
	if _, err = r.db.ExecContext(ctx, query, attributeID); err != nil {
		return
	}

	return
}
