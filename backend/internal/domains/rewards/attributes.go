package rewardsDomains

import (
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type Attribute struct {
	id        uuid.UUID
	rewardID  uuid.UUID
	traitType string
	value     string
}

type AttributeFilter struct {
	ID        uuid.UUID
	RewardID  uuid.UUID
	TraitType string
}

func NewAttribute(
	id, rewardID uuid.UUID,
	traitType, value string,
) (*Attribute, error) {
	var attribute Attribute
	if err := attribute.SetTraitType(traitType); err != nil {
		return nil, err
	}
	if err := attribute.SetValue(value); err != nil {
		return nil, err
	}

	return &attribute, nil
}

func (d *Attribute) Unmarshal(
	id, rewardID uuid.UUID,
	traitType, value string,
) {
	d.id = id
	d.rewardID = rewardID
	d.traitType = traitType
	d.value = value
}

func (d *Attribute) GetID() uuid.UUID {
	return d.id
}

func (d *Attribute) GetRewardID() uuid.UUID {
	return d.rewardID
}

func (d *Attribute) GetTraitType() string {
	return d.traitType
}

func (d *Attribute) SetTraitType(traitType string) error {
	if traitType == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrTraitTypeCannotBeEmpty)
	}
	if len(traitType) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrTraitTypeTooLong)
	}
	d.traitType = traitType

	return nil
}

func (d *Attribute) GetValue() string {
	return d.value
}

func (d *Attribute) SetValue(value string) error {
	if value == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrValueCannotBeEmpty)
	}
	if len(value) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrValueTooLong)
	}
	d.value = value

	return nil
}
