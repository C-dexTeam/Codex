package dto

import (
	"github.com/C-dexTeam/codex/internal/domains"
	"github.com/google/uuid"
)

type LanguageDTOManager struct{}

func NewLanguageDTOManager() LanguageDTOManager {
	return LanguageDTOManager{}
}

type LanguageDTO struct {
	ID    uuid.UUID `json:"id"`
	Value string    `json:"value"`
}

func (m *LanguageDTOManager) ToLanguageDTO(appModel *domains.Languages) LanguageDTO {
	return LanguageDTO{
		ID:    appModel.GetID(),
		Value: appModel.GetValue(),
	}
}

func (m *LanguageDTOManager) ToLanguageDTOs(appModels []domains.Languages) []LanguageDTO {
	var languagesDTOs []LanguageDTO
	for _, model := range appModels {
		languagesDTOs = append(languagesDTOs, m.ToLanguageDTO(&model))
	}
	return languagesDTOs
}
