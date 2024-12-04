package dto

import (
	"github.com/C-dexTeam/codex/internal/domains"
	"github.com/google/uuid"
)

type RewardDTOManager struct{}

func NewRewardDTOManager() RewardDTOManager {
	return RewardDTOManager{}
}

type RewardDTO struct {
	ID          uuid.UUID      `json:"id"`
	RewardType  string         `json:"rewardType"`
	Name        string         `json:"name"`
	Symbol      string         `json:"symbol"`
	Description string         `json:"Description"`
	ImagePath   string         `json:"imagePath"`
	URI         string         `json:"uri"`
	Attributes  []AttributeDTO `json:"attributes,omitempty"`
}

func (m *RewardDTOManager) ToRewardDTO(appModel *domains.Reward) RewardDTO {
	return RewardDTO{
		ID:          appModel.GetID(),
		RewardType:  appModel.GetRewardType(),
		Name:        appModel.GetName(),
		Symbol:      appModel.GetSymbol(),
		Description: appModel.GetDescription(),
		ImagePath:   appModel.GetImagePath(),
		URI:         appModel.GetURI(),
		Attributes:  m.ToAttributeDTOs(appModel.GetAttribute()),
	}
}

func (m *RewardDTOManager) ToRewardDTOs(appModels []domains.Reward) []RewardDTO {
	var rewardDTOs []RewardDTO
	for _, model := range appModels {
		rewardDTOs = append(rewardDTOs, m.ToRewardDTO(&model))
	}
	return rewardDTOs
}

type AddRewardDTO struct {
	RewardType  string `json:"rewardType" validate:"required,max=30"`
	Name        string `json:"name" validate:"required,max=30"`
	Symbol      string `json:"symbol" validate:"required,max=30"`
	Description string `json:"Description"`
	ImagePath   string `json:"imagePath" validate:"required,max=60"`
	URI         string `json:"uri" validate:"required,max=120"`
}

type UpdateRewardDTO struct {
	ID          string `json:"id" validate:"required"`
	RewardType  string `json:"rewardType" validate:"max=30"`
	Name        string `json:"name" validate:"max=30"`
	Symbol      string `json:"symbol" validate:"max=30"`
	Description string `json:"Description"`
	ImagePath   string `json:"imagePath" validate:"max=60"`
	URI         string `json:"uri" validate:"max=120"`
}

// ATTRIBUTES
type AttributeDTO struct {
	ID        uuid.UUID  `json:"id"`
	RewardID  *uuid.UUID `json:"rewardID"`
	TraitType string     `json:"traitType" validate:"required,max=30"`
	Value     string     `json:"value" validate:"required,max=30"`
}

func (m *RewardDTOManager) ToAttributeDTO(appModel *domains.Attribute) AttributeDTO {
	return AttributeDTO{
		ID:        appModel.GetID(),
		RewardID:  appModel.GetRewardID(),
		TraitType: appModel.GetTraitType(),
		Value:     appModel.GetValue(),
	}
}

func (m *RewardDTOManager) ToAttributeDTOs(appModels []domains.Attribute) []AttributeDTO {
	var attributeDTOs []AttributeDTO
	for _, model := range appModels {
		attributeDTOs = append(attributeDTOs, m.ToAttributeDTO(&model))
	}
	return attributeDTOs
}

type AddAttributeDTO struct {
	RewardID  string `json:"rewardID"`
	TraitType string `json:"traitType" validate:"required,max=30"`
	Value     string `json:"value" validate:"required,max=30"`
}

type UpdateAttributeDTO struct {
	ID        string `json:"id"`
	RewardID  string `json:"rewardID"`
	TraitType string `json:"traitType" validate:"max=30"`
	Value     string `json:"value" validate:"max=30"`
}
