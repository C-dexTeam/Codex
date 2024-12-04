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

func (r *RewardsRepository) dbModelFromAppModel(appModel domains.Reward) (dbModel dbModelReward) {
	if appModel.GetID() != uuid.Nil {
		dbModel.ID.String = appModel.GetID().String()
		dbModel.ID.Valid = true
	}
	if appModel.GetRewardType() != "" {
		dbModel.RewardType.String = appModel.GetRewardType()
		dbModel.RewardType.Valid = true
	}
	if appModel.GetSymbol() != "" {
		dbModel.Symbol.String = appModel.GetSymbol()
		dbModel.Symbol.Valid = true
	}
	if appModel.GetName() != "" {
		dbModel.Name.String = appModel.GetName()
		dbModel.Name.Valid = true
	}
	if appModel.GetDescription() != "" {
		dbModel.Description.String = appModel.GetDescription()
		dbModel.Description.Valid = true
	}
	if appModel.GetImagePath() != "" {
		dbModel.ImagePath.String = appModel.GetImagePath()
		dbModel.ImagePath.Valid = true
	}
	if appModel.GetURI() != "" {
		dbModel.URI.String = appModel.GetURI()
		dbModel.URI.Valid = true
	}
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
	dbFilter := r.dbModelFromAppFilter(filter)
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
	err = r.db.SelectContext(ctx, &dbResult, query, dbFilter.ID, dbFilter.RewardType, dbFilter.Name, dbFilter.Symbol, limit, (page-1)*limit)
	if err != nil {
		return
	}
	for _, dbModel := range dbResult {
		rewards = append(rewards, r.dbModelToAppModel(dbModel))
	}
	return
}

func (r *RewardsRepository) Add(ctx context.Context, reward *domains.Reward) (uuid.UUID, error) {
	dbModel := r.dbModelFromAppModel(*reward)
	query := `
		INSERT INTO
			t_rewards (reward_type, symbol, name, description, image_path, uri)
		VALUES
			($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var id uuid.UUID
	err := r.db.QueryRowxContext(ctx, query, dbModel.RewardType, dbModel.Symbol, dbModel.Name, dbModel.Description, dbModel.ImagePath, dbModel.URI).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *RewardsRepository) Update(ctx context.Context, reward *domains.Reward) (err error) {
	dbModel := r.dbModelFromAppModel(*reward) // thanks to this. This func behaves like patch.
	query := `
		UPDATE
			t_rewards
		SET
			reward_type = COALESCE(:reward_type, reward_type),
			symbol = COALESCE(:symbol, symbol),
			name = COALESCE(:name, name),
			description =  COALESCE(:description, description),
			image_path =  COALESCE(:image_path, image_path),
			uri =  COALESCE(:uri, uri)
		WHERE
			id = :id
	`
	_, err = r.db.NamedExecContext(ctx, query, dbModel)
	if err != nil {
		return
	}
	return
}

func (r *RewardsRepository) Delete(ctx context.Context, rewardID uuid.UUID) (err error) {
	query := `
		DELETE FROM
			t_rewards
		WHERE 
			id = $1
	`
	if _, err = r.db.ExecContext(ctx, query, rewardID); err != nil {
		return
	}

	return
}
