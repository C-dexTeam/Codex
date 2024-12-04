package domains

import (
	"context"

	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type IRewardRepository interface {
	Filter(ctx context.Context, filter RewardFilter, limit, page int64) (rewards []Reward, dataCount int64, err error)
	Add(ctx context.Context, reward *Reward) (uuid.UUID, error)
	Update(ctx context.Context, reward *Reward) (err error)
	Delete(ctx context.Context, rewardID uuid.UUID) (err error)
}

type IRewardService interface {
	GetRewards(ctx context.Context, rewardID, name, symbol, rewardType, page, limit string) (rewards []Reward, err error)
	GetReward(ctx context.Context, rewardID, page, limit string) (reward *Reward, err error)
	AddReward(ctx context.Context, rewardType, symbol, name, description, imagePath, URI string) (uuid.UUID, error)
	UpdateReward(ctx context.Context, id, rewardType, symbol, name, description, imagePath, URI string) error
	DeleteReward(ctx context.Context, rewardID string) (err error)
}

const (
	DefaultRewardLimit = 10
)

type Reward struct {
	id          uuid.UUID
	rewardType  string
	name        string
	symbol      string
	description string
	imagePath   string
	uri         string
	attribute   []Attribute
}

type RewardFilter struct {
	ID         uuid.UUID
	Name       string
	Symbol     string
	RewardType string
}

func NewReward(
	id, rewardType, symbol, name, description, imagePath, uri string,
	attribute []Attribute,
) (*Reward, error) {
	reward := Reward{}
	if err := reward.SetID(id); err != nil {
		return nil, err
	}
	if err := reward.SetRewardType(rewardType); err != nil {
		return nil, err
	}
	if err := reward.SetSymbol(symbol); err != nil {
		return nil, err
	}
	if err := reward.SetName(name); err != nil {
		return nil, err
	}
	if err := reward.SetImagePath(imagePath); err != nil {
		return nil, err
	}
	if err := reward.SetURI(uri); err != nil {
		return nil, err
	}
	reward.SetDescription(description)
	reward.SetAttribute(attribute)

	return &reward, nil
}

func (d *Reward) Unmarshal(
	id uuid.UUID,
	rewardType, symbol, name, description, imagePath, uri string,
) {
	d.id = id
	d.rewardType = rewardType
	d.symbol = symbol
	d.name = name
	d.description = description
	d.imagePath = imagePath
	d.uri = uri
}

func (d *Reward) GetID() uuid.UUID {
	return d.id
}

func (d *Reward) SetID(id string) error {
	if id != "" {
		idUUID, err := uuid.Parse(id)
		if err != nil {
			return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
		}
		d.id = idUUID
	}

	return nil
}

func (d *Reward) GetRewardType() string {
	return d.rewardType
}

func (d *Reward) SetRewardType(rewardType string) error {
	if len(rewardType) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrRewardTypeTooLong)
	}
	d.rewardType = rewardType
	return nil
}

func (d *Reward) GetName() string {
	return d.name
}

func (d *Reward) SetName(name string) error {
	if len(name) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrRewardNameTooLong)
	}
	d.name = name
	return nil
}

func (d *Reward) GetSymbol() string {
	return d.symbol
}

func (d *Reward) SetSymbol(symbol string) error {
	if len(symbol) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrRewardSymbolTooLong)
	}
	d.symbol = symbol
	return nil
}

func (d *Reward) GetDescription() string {
	return d.description
}

func (d *Reward) SetDescription(description string) {
	d.description = description
}

func (d *Reward) GetImagePath() string {
	return d.imagePath
}

func (d *Reward) SetImagePath(imagePath string) error {
	if len(imagePath) > 60 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrRewardImagePathTooLong)
	}
	d.imagePath = imagePath
	return nil
}

func (d *Reward) GetURI() string {
	return d.uri
}

func (d *Reward) SetURI(uri string) error {
	if len(uri) > 120 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrRewardURITooLong)
	}
	d.uri = uri
	return nil
}

func (d *Reward) GetAttribute() []Attribute {
	return d.attribute
}

func (d *Reward) SetAttribute(attribute []Attribute) {
	d.attribute = attribute
}
