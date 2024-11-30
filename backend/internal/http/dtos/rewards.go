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
	ID          uuid.UUID `json:"id"`
	RewardType  string    `json:"rewardType"`
	Name        string    `json:"name"`
	Symbol      string    `json:"symbol"`
	Description string    `json:"Description"`
	ImagePath   string    `json:"imagePath"`
	URI         string    `json:"uri"`
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
	}
}

func (m *RewardDTOManager) ToRewardDTOs(appModels []domains.Reward) []RewardDTO {
	var rewardDTOs []RewardDTO
	for _, model := range appModels {
		rewardDTOs = append(rewardDTOs, m.ToRewardDTO(&model))
	}
	return rewardDTOs
}
