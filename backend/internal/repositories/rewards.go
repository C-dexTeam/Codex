package repositories

import (
	"context"
	"database/sql"

	"github.com/C-dexTeam/codex/internal/domains"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RewardsRepository struct {
	db *sqlx.DB
}

type dbModelReward struct {
	ID          sql.NullString `db:"id"`
	RewardType  sql.NullString `db:"reward_type"`
	Name        sql.NullString `db:"name"`
	Symbol      sql.NullString `db:"symbol"`
	Description sql.NullString `db:"description"`
	ImagePath   sql.NullString `db:"image_path"`
	URI         sql.NullString `db:"uri"`
}

func (r *RewardsRepository) dbModelToAppModel(dbModel dbModelReward) (reward domains.Reward) {
	reward.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		dbModel.RewardType.String,
		dbModel.Symbol.String,
		dbModel.Name.String,
		dbModel.Description.String,
		dbModel.ImagePath.String,
		dbModel.URI.String,
	)

	return
}

func (r *RewardsRepository) dbModelFromAppFilter(filter domains.RewardFilter) (dbModelReward dbModelReward) {
	if filter.ID != uuid.Nil {
		dbModelReward.ID.String = filter.ID.String()
		dbModelReward.ID.Valid = true
	}
	if filter.RewardType != "" {
		dbModelReward.RewardType.String = filter.RewardType
		dbModelReward.RewardType.Valid = true
	}
	if filter.Symbol != "" {
		dbModelReward.Symbol.String = filter.Symbol
		dbModelReward.Symbol.Valid = true
	}
	if filter.Name != "" {
		dbModelReward.Name.String = filter.Name
		dbModelReward.Name.Valid = true
	}

	return
}

func NewRewardsRepository(db *sqlx.DB) domains.IRewardRepository {
	return &RewardsRepository{db: db}
}

func (r *RewardsRepository) Filter(ctx context.Context, filter domains.RewardFilter, limit, page int64) (rewards []domains.Reward, dataCount int64, err error) {
	dbFilterReward := r.dbModelFromAppFilter(filter)
	dbResult := []dbModelReward{}

	query := `
	SELECT
		r.*
	FROM 
		t_rewards r
	WHERE
		($1::uuid IS NULL OR r.id = $1::uuid) AND
		($2::text IS NULL OR r.reward_type = $2::text) AND
		($3::text IS NULL OR r.name LIKE '%' || $3::text || '%') AND
		($4::text IS NULL OR r.symbol LIKE '%' || $4::text || '%')
	LIMIT $5 OFFSET $6
	`

	// Execute the query with the extracted fields
	err = r.db.SelectContext(ctx, &dbResult, query, dbFilterReward.ID, dbFilterReward.RewardType, dbFilterReward.Name, dbFilterReward.Symbol, limit, (page-1)*limit)
	if err != nil {
		return
	}
	for _, dbModel := range dbResult {
		rewards = append(rewards, r.dbModelToAppModel(dbModel))
	}
	return
}
