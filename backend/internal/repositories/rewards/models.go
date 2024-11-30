package rewardsRepositories

import (
	"database/sql"

	rewardsDomains "github.com/C-dexTeam/codex/internal/domains/rewards"
	"github.com/google/uuid"
)

type dbModelReward struct {
	ID          sql.NullString           `db:"id"`
	RewardType  sql.NullString           `db:"reward_type"`
	Name        sql.NullString           `db:"name"`
	Symbol      sql.NullString           `db:"symbol"`
	Description sql.NullString           `db:"description"`
	ImagePath   sql.NullString           `db:"image_path"`
	URI         sql.NullString           `db:"uri"`
	Attributes  []dbModelRewardAttribute `db:"-"`
}

type dbModelRewardAttribute struct {
	ID        sql.NullString `db:"id"`
	RewardID  sql.NullString `db:"reward_id"`
	TraitType sql.NullString `db:"trait_type"`
	Value     sql.NullString `db:"value"`
}

// TODO: dbModelFromAppModel

func (r *RewardsRepository) dbModelToAppModelReward(dbModel dbModelReward) (reward rewardsDomains.Reward) {
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

func (r *RewardsRepository) dbModelFromAppFilterReward(filter rewardsDomains.RewardFilter) (dbModel dbModelReward) {
	if filter.ID != uuid.Nil {
		dbModel.ID.String = filter.ID.String()
		dbModel.ID.Valid = true
	}
	if filter.RewardType != "" {
		dbModel.RewardType.String = filter.RewardType
		dbModel.RewardType.Valid = true
	}
	if filter.Symbol != "" {
		dbModel.Symbol.String = filter.Symbol
		dbModel.Symbol.Valid = true
	}
	if filter.Name != "" {
		dbModel.Name.String = filter.Name
		dbModel.Name.Valid = true
	}
	return
}

func (r *RewardsRepository) dbModelToAppModelAttribute(dbModel dbModelRewardAttribute) (attribute rewardsDomains.Attribute) {
	attribute.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		uuid.MustParse(dbModel.RewardID.String),
		dbModel.TraitType.String,
		dbModel.Value.String,
	)

	return
}

func (r *RewardsRepository) dbModelFromAppFilterAttribute(filter rewardsDomains.AttributeFilter) (dbModel dbModelRewardAttribute) {
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

func (r *RewardsRepository) dbModelToAppModel(dbModel dbModelReward) (reward rewardsDomains.Reward) {
	var attributes []rewardsDomains.Attribute

	for _, att := range dbModel.Attributes {
		var attribute rewardsDomains.Attribute
		attribute.Unmarshal(
			uuid.MustParse(att.ID.String),
			uuid.MustParse(att.RewardID.String),
			att.TraitType.String,
			att.Value.String,
		)

		attributes = append(attributes, attribute)
	}

	reward.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		dbModel.RewardType.String,
		dbModel.Symbol.String,
		dbModel.Name.String,
		dbModel.Description.String,
		dbModel.ImagePath.String,
		dbModel.URI.String,
	)
	reward.SetAttribute(attributes)

	return
}

func (r *RewardsRepository) dbModelFromAppFilter(filter rewardsDomains.RewardFilter) (dbModel dbModelReward) {
	var reward dbModelReward
	var attributes []dbModelRewardAttribute

	if filter.ID != uuid.Nil {
		reward.ID.String = filter.ID.String()
		reward.ID.Valid = true
	}
	if filter.RewardType != "" {
		reward.RewardType.String = filter.RewardType
		reward.RewardType.Valid = true
	}
	if filter.Symbol != "" {
		reward.Symbol.String = filter.Symbol
		reward.Symbol.Valid = true
	}
	if filter.Name != "" {
		reward.Name.String = filter.Name
		reward.Name.Valid = true
	}

	for _, attributeFilter := range filter.Attributes {
		var attribute dbModelRewardAttribute

		if attributeFilter.ID != uuid.Nil {
			attribute.ID.String = attributeFilter.ID.String()
			attribute.ID.Valid = true
		}
		if attributeFilter.RewardID != uuid.Nil {
			attribute.RewardID.String = attributeFilter.RewardID.String()
			attribute.RewardID.Valid = true
		}
		if attributeFilter.TraitType != "" {
			attribute.TraitType.String = attributeFilter.TraitType
			attribute.TraitType.Valid = true
		}

		attributes = append(attributes, attribute)
	}
	reward.Attributes = attributes

	return reward
}
