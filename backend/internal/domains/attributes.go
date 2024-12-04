package domains

import (
	"context"

	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type IAttributeRepository interface {
	Filter(ctx context.Context, filter AttributeFilter, limit, page int64) (rewards []Attribute, dataCount int64, err error)
	Add(ctx context.Context, attribute *Attribute) (uuid.UUID, error)
	Update(ctx context.Context, attribute *Attribute) (err error)
	Delete(ctx context.Context, attributeID uuid.UUID) (err error)
}

type IAttributeService interface {
	GetAttributes(ctx context.Context, attributeID, rewardID, traitType, page, limit string) (attributes []Attribute, err error)
	AddAttribute(ctx context.Context, rewardID, traitType, value string) (uuid.UUID, error)
	UpdateAttribute(ctx context.Context, id, rewardID, traitType, value string) error
	DeleteAttribute(ctx context.Context, attributeID string) (err error)
}

const (
	DefaultAttributeLimit = 10
)

type Attribute struct {
	id        uuid.UUID
	rewardID  *uuid.UUID
	traitType string
	value     string
}

type AttributeFilter struct {
	ID        uuid.UUID
	RewardID  uuid.UUID
	TraitType string
}

func NewAttribute(
	id, rewardID string,
	traitType, value string,
) (*Attribute, error) {
	var attribute Attribute
	if err := attribute.SetID(id); err != nil {
		return nil, err
	}
	if err := attribute.SetRewardID(rewardID); err != nil {
		return nil, err
	}
	if err := attribute.SetTraitType(traitType); err != nil {
		return nil, err
	}
	if err := attribute.SetValue(value); err != nil {
		return nil, err
	}

	return &attribute, nil
}

func (d *Attribute) Unmarshal(
	id uuid.UUID,
	rewardID *uuid.UUID,
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

func (d *Attribute) SetID(id string) error {
	if id != "" {
		idUUID, err := uuid.Parse(id)
		if err != nil {
			return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
		}
		d.id = idUUID
	}

	return nil
}

func (d *Attribute) GetRewardID() *uuid.UUID {
	return d.rewardID
}

func (d *Attribute) SetRewardID(rewardID string) (err error) {
	// Like this. bc rewardID can be null
	if rewardID != "" {
		id, err := uuid.Parse(rewardID)
		if err != nil {
			return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
		d.rewardID = &id
	} else {
		// Nil gitsin diye diğer türlü uuid.nil gidiyor.
		d.rewardID = nil
	}

	return nil
}

func (d *Attribute) GetTraitType() string {
	return d.traitType
}

func (d *Attribute) SetTraitType(traitType string) error {
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
	if len(value) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrValueTooLong)
	}
	d.value = value

	return nil
}
