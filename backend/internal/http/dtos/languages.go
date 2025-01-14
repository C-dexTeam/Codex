package dto

import (
	repo "github.com/C-dexTeam/codex/internal/repos/out"
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

func (m *LanguageDTOManager) ToLanguageDTO(appModel *repo.TLanguage) LanguageDTO {
	return LanguageDTO{
		ID:    appModel.ID,
		Value: appModel.Value,
	}
}

func (m *LanguageDTOManager) ToLanguageDTOs(appModels []repo.TLanguage) []LanguageDTO {
	var languagesDTOs []LanguageDTO
	for _, model := range appModels {
		languagesDTOs = append(languagesDTOs, m.ToLanguageDTO(&model))
	}
	return languagesDTOs
}
