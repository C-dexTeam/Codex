package rewardsDomains

import (
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

const (
	DefaultRewardLimit = 10
)

type Reward struct {
	id          uuid.UUID
	rewardType  string
	symbol      string
	name        string
	description string
	imagePath   string
	uri         string
}

func NewReward(
	id uuid.UUID,
	rewardType, symbol, name, description, imagePath, uri string,
) (*Reward, error) {
	reward := Reward{}
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

func (d *Reward) GetRewardType() string {
	return d.rewardType
}

func (d *Reward) SetRewardType(rewardType string) error {
	if rewardType == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrRewardTypeCannotBeEmpty)
	}
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
	if name == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrRewardNameCannotBeEmpty)
	}
	if len(name) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrRewardNameTooLong)
	}
	d.name = name
	return nil
}

func (d *Reward) GetSymbol() string {
	return d.name
}

func (d *Reward) SetSymbol(symbol string) error {
	if symbol == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrRewardSymbolCannotBeEmpty)
	}
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
	if imagePath == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrRewardImagePathCannotBeEmpty)
	}
	if len(imagePath) > 60 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrRewardImagePathTooLong)
	}
	d.imagePath = imagePath
	return nil
}

func (d *Reward) GetURI() string {
	return d.imagePath
}

func (d *Reward) SetURI(uri string) error {
	if uri == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrRewardURICannotBeEmpty)
	}
	if len(uri) > 120 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrRewardURITooLong)
	}
	d.uri = uri
	return nil
}
