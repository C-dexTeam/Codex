package domains

import (
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type Reward struct {
	ID          uuid.UUID
	Name        string
	Symbol      string
	Description string
	ImagePath   string
	URI         string
	Attributes  *Attributes
	SellerFee   int
}

type Attribute struct {
	ID        uuid.UUID
	RewardID  uuid.UUID
	TraitType string
	Value     string
}

type Attributes struct {
	Attributes     []Attribute
	TotalAttribute int64
}

func NewReward(reward *repo.TReward, attributes []repo.TAttribute) *Reward {
	if reward == nil {
		return nil
	}

	return &Reward{
		ID:          reward.ID,
		Name:        reward.Name,
		Symbol:      reward.Symbol,
		Description: reward.Description,
		ImagePath:   reward.ImagePath.String,
		URI:         reward.Uri.String,
		Attributes:  NewAttributes(attributes, 0),
		SellerFee:   int(reward.SellerFee),
	}
}

func NewRewards(rewards []repo.TReward, attributes []repo.TAttribute) []Reward {
	var result []Reward

	for _, reward := range rewards {
		result = append(result, *NewReward(&reward, attributes))
	}

	return result
}

func NewAttribute(attribute *repo.TAttribute) *Attribute {
	if attribute == nil {
		return nil
	}

	return &Attribute{
		ID:        attribute.ID,
		RewardID:  attribute.RewardID,
		TraitType: attribute.TraitType,
		Value:     attribute.Value,
	}
}

func NewAttributes(attributes []repo.TAttribute, count int64) *Attributes {
	var result []Attribute

	for _, attribute := range attributes {
		result = append(result, *NewAttribute(&attribute))
	}

	return &Attributes{Attributes: result, TotalAttribute: count}
}
