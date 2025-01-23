package dto

import (
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type RewardDTOManager struct{}

func NewRewardDTOManager() RewardDTOManager {
	return RewardDTOManager{}
}

type RewardDTO struct {
	ID          uuid.UUID      `json:"id"`
	RewardType  string         `json:"rewardType" validate:"required"`
	Name        string         `json:"name" validate:"required,min=3,max=30"`
	Symbol      string         `json:"symbol" validate:"required,min=2,max=8"`
	Description string         `json:"description" validate:"required"`
	ImagePath   string         `json:"imagePath"`
	URI         string         `json:"uri" validate:"required"`
	Attributes  []AttributeDTO `json:"attributes,omitempty"`
}

func (m *RewardDTOManager) ToRewardDTO(appModel *repo.TReward, appAttributeModel []repo.TAttribute) RewardDTO {
	return RewardDTO{
		ID:          appModel.ID,
		RewardType:  appModel.RewardType,
		Name:        appModel.Name,
		Symbol:      appModel.Symbol,
		Description: appModel.Description,
		ImagePath:   appModel.ImagePath,
		URI:         appModel.Uri,
		Attributes:  m.ToAttributeDTOs(appAttributeModel),
	}
}

func (m *RewardDTOManager) ToRewardDTOs(appModels []repo.TReward) []RewardDTO {
	var rewardDTOs []RewardDTO
	for _, model := range appModels {
		rewardDTOs = append(rewardDTOs, m.ToRewardDTO(&model, nil))
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
	ID        uuid.UUID `json:"id"`
	RewardID  uuid.UUID `json:"rewardID"`
	TraitType string    `json:"traitType" validate:"required,max=30"`
	Value     string    `json:"value" validate:"required,max=30"`
}

func (m *RewardDTOManager) ToAttributeDTO(appModel *repo.TAttribute) AttributeDTO {
	return AttributeDTO{
		ID:        appModel.ID,
		RewardID:  appModel.RewardID,
		TraitType: appModel.TraitType,
		Value:     appModel.Value,
	}
}

func (m *RewardDTOManager) ToAttributeDTOs(appModels []repo.TAttribute) []AttributeDTO {
	var attributeDTOs []AttributeDTO
	for _, model := range appModels {
		attributeDTOs = append(attributeDTOs, m.ToAttributeDTO(&model))
	}
	return attributeDTOs
}

type AddAttributeDTO struct {
	RewardID  string `json:"rewardID" validate:"required"`
	TraitType string `json:"traitType" validate:"required,max=30"`
	Value     string `json:"value" validate:"required,max=30"`
}

type UpdateAttributeDTO struct {
	ID        string `json:"id"`
	RewardID  string `json:"rewardID" validate:"required"`
	TraitType string `json:"traitType" validate:"max=30"`
	Value     string `json:"value" validate:"max=30"`
}
