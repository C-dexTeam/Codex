package rewardsRepositories

import (
	"context"

	rewardsDomains "github.com/C-dexTeam/codex/internal/domains/rewards"
	"github.com/jmoiron/sqlx"
)

type RewardsRepository struct {
	db *sqlx.DB
}

func NewRewardsRepository(db *sqlx.DB) rewardsDomains.IRewardRepository {
	return &RewardsRepository{db: db}
}

func (r *RewardsRepository) Filter(ctx context.Context, filter rewardsDomains.RewardFilter, limit, page int64) (rewards []rewardsDomains.Reward, dataCount int64, err error) {
	dbFilter := r.dbModelFromAppFilter(filter)
	dbResult := []dbModelReward{}

	query := `
		SELECT
			*
		FROM
			t_rewards r
		LEFT JOIN
			t_attributes a
		ON
			r.id = a.reward_id
		WHERE
			($1::uuid IS NULL OR r.id = $1::uuid) AND
			($2::text IS NULL OR r.reward_type = $2::text) AND
			($3::text IS NULL OR r.name LIKE '%' || $3::text || '%') AND
			($4::text IS NULL OR r.symbol LIKE '%' || $4::text || '%') AND
			($5::uuid IS NULL OR a.id = $5::uuid) AND
			($6::uuid IS NULL OR a.reward_id = $6::uuid) AND
			($7::text IS NULL OR a.trait_type LIKE '%' || $7::text || '%')
		LIMIT $8 OFFSET $9;
	`

	if err = r.db.SelectContext(ctx, &dbResult, query, dbFilter.ID, dbFilter.RewardType, dbFilter.Name, dbFilter.Symbol, dbFilter.Attributes[0].ID, dbFilter.Attributes[0].RewardID, dbFilter.Attributes[0].TraitType, limit, (page-1)*limit); err != nil {
		return
	}
	for _, dbModel := range dbResult {
		rewards = append(rewards, r.dbModelToAppModel(dbModel))
	}

	return
}
