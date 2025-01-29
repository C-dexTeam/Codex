package dto

import (
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type RewardDTOManager struct{}

func NewRewardDTOManager() RewardDTOManager {
	return RewardDTOManager{}
}

type RewardView struct {
	ID          uuid.UUID      `json:"id"`
	RewardType  string         `json:"rewardType" validate:"required"`
	Name        string         `json:"name" validate:"required,min=3,max=30"`
	Symbol      string         `json:"symbol" validate:"required,min=2,max=8"`
	Description string         `json:"description" validate:"required"`
	ImagePath   string         `json:"imagePath"`
	URI         string         `json:"uri" validate:"required"`
	Attributes  []AttributeDTO `json:"attributes,omitempty"`
}

func (m *RewardDTOManager) ToRewardDTO(appModel *repo.TReward, appAttributeModel []repo.TAttribute) RewardView {
	return RewardView{
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

func (m *RewardDTOManager) ToRewardDTOs(appModels []repo.TReward) []RewardView {
	var rewardDTOs []RewardView
	for _, model := range appModels {
		rewardDTOs = append(rewardDTOs, m.ToRewardDTO(&model, nil))
	}
	return rewardDTOs
}

type MetadataView struct {
	Name        string                  `json:"name"`
	Symbol      string                  `json:"symbol"`
	Description string                  `json:"description"`
	Image       string                  `json:"image"`
	URI         string                  `json:"uri"`
	Attributes  []MetadataAttributeView `json:"attributes,omitempty"`
	SellerFee   int                     `json:"seller_fee_basis_points"`
}

func (m *RewardDTOManager) ToMetadataView(appModel *repo.TReward, appAttributeModel []repo.TAttribute, URL string) MetadataView {
	return MetadataView{
		Name:        appModel.Name,
		Symbol:      appModel.Symbol,
		Description: appModel.Description,
		Image:       URL + appModel.ImagePath,
		URI:         appModel.Uri,
		Attributes:  m.ToMetadataAttributeDTOs(appAttributeModel),
	}
}

type MetadataAttributeView struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

func (m *RewardDTOManager) ToMetadataAttributeDTO(appModel *repo.TAttribute) MetadataAttributeView {
	return MetadataAttributeView{
		TraitType: appModel.TraitType,
		Value:     appModel.Value,
	}
}

func (m *RewardDTOManager) ToMetadataAttributeDTOs(appModels []repo.TAttribute) []MetadataAttributeView {
	var attributeDTOs []MetadataAttributeView
	for _, model := range appModels {
		attributeDTOs = append(attributeDTOs, m.ToMetadataAttributeDTO(&model))
	}
	return attributeDTOs
}

type AddRewardDTO struct {
	RewardType  string `json:"rewardType" validate:"required,max=30"`
	Name        string `json:"name" validate:"required,max=30"`
	Symbol      string `json:"symbol" validate:"required,max=30"`
	Description string `json:"Description"`
}

type UpdateRewardDTO struct {
	ID          string `json:"id" validate:"required"`
	RewardType  string `json:"rewardType" validate:"max=30"`
	Name        string `json:"name" validate:"max=30"`
	Symbol      string `json:"symbol" validate:"max=30"`
	Description string `json:"Description"`
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
