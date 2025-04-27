package dto

import (
	"github.com/C-dexTeam/codex/internal/domains"
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type RewardDTOManager struct{}

func NewRewardDTOManager() RewardDTOManager {
	return RewardDTOManager{}
}

type UserRewardView struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" validate:"required,min=3,max=30"`
	Symbol      string    `json:"symbol" validate:"required,min=2,max=8"`
	Description string    `json:"description" validate:"required"`
	ImagePath   string    `json:"imagePath"`
	URI         string    `json:"uri" validate:"required"`
}

func (m *RewardDTOManager) ToUserRewardDTO(appModel *repo.UserRewardsRow) UserRewardView {
	return UserRewardView{
		ID:          appModel.ID,
		Name:        appModel.Name,
		Symbol:      appModel.Symbol,
		Description: appModel.Description,
		ImagePath:   appModel.ImagePath.String,
		URI:         appModel.Uri.String,
	}
}

func (m *RewardDTOManager) ToUserRewardDTOs(appModels []repo.UserRewardsRow) []UserRewardView {
	var rewardDTOs []UserRewardView
	for _, model := range appModels {
		rewardDTOs = append(rewardDTOs, m.ToUserRewardDTO(&model))
	}
	return rewardDTOs
}

type RewardView struct {
	ID          uuid.UUID      `json:"id"`
	Name        string         `json:"name" validate:"required,min=3,max=30"`
	Symbol      string         `json:"symbol" validate:"required,min=2,max=8"`
	Description string         `json:"description" validate:"required"`
	ImagePath   string         `json:"imagePath"`
	URI         string         `json:"uri" validate:"required"`
	Attributes  []AttributeDTO `json:"attributes,omitempty"`
}

func (m *RewardDTOManager) ToRewardDTO(appModel *domains.Reward) RewardView {
	return RewardView{
		ID:          appModel.ID,
		Name:        appModel.Name,
		Symbol:      appModel.Symbol,
		Description: appModel.Description,
		ImagePath:   appModel.ImagePath,
		URI:         appModel.URI,
		Attributes:  m.ToAttributeDTOs(appModel.Attributes),
	}
}

func (m *RewardDTOManager) ToRewardDTOs(appModels []domains.Reward) []RewardView {
	var rewardDTOs []RewardView
	for _, model := range appModels {
		rewardDTOs = append(rewardDTOs, m.ToRewardDTO(&model))
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

func (m *RewardDTOManager) ToMetadataView(appModel *domains.Reward, URL string) MetadataView {
	return MetadataView{
		Name:        appModel.Name,
		Symbol:      appModel.Symbol,
		Description: appModel.Description,
		Image:       URL + appModel.ImagePath,
		URI:         appModel.URI,
		Attributes:  m.ToMetadataAttributeDTOs(appModel.Attributes),
	}
}

type MetadataAttributeView struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

func (m *RewardDTOManager) ToMetadataAttributeDTO(appModel *domains.Attribute) MetadataAttributeView {
	return MetadataAttributeView{
		TraitType: appModel.TraitType,
		Value:     appModel.Value,
	}
}

func (m *RewardDTOManager) ToMetadataAttributeDTOs(appModels []domains.Attribute) []MetadataAttributeView {
	var attributeDTOs []MetadataAttributeView
	for _, model := range appModels {
		attributeDTOs = append(attributeDTOs, m.ToMetadataAttributeDTO(&model))
	}
	return attributeDTOs
}

type AddRewardDTO struct {
	Name        string `json:"name" validate:"required,max=30"`
	Symbol      string `json:"symbol" validate:"required,max=30"`
	Description string `json:"Description"`
	SellerFee   int    `json:"sellerFee" `
}

type UpdateRewardDTO struct {
	ID          string `json:"id" validate:"required"`
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

func (m *RewardDTOManager) ToAttributeDTO(appModel *domains.Attribute) AttributeDTO {
	return AttributeDTO{
		ID:        appModel.ID,
		RewardID:  appModel.RewardID,
		TraitType: appModel.TraitType,
		Value:     appModel.Value,
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
	RewardID  string `json:"rewardID" validate:"required"`
	TraitType string `json:"traitType" validate:"required,max=30"`
	Value     string `json:"value" validate:"required,max=30"`
}

type UpdateAttributeDTO struct {
	ID        string `json:"id"`
	RewardID  string `json:"rewardID"`
	TraitType string `json:"traitType"`
	Value     string `json:"value"`
}
