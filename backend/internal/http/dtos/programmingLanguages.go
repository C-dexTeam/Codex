package dto

import (
	"time"

	"github.com/C-dexTeam/codex/internal/domains"
	"github.com/google/uuid"
)

type ProgrammingLanguageDTOManager struct{}

func NewProgrammingLanguageDTOManager() ProgrammingLanguageDTOManager {
	return ProgrammingLanguageDTOManager{}
}

type UserPLanguageView struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	FileExtention string    `json:"fileExtention"`
	MonacoEditor  string    `json:"monacoEditor"`
	CreatedAt     time.Time `json:"createdAt"`
}

func (m *ProgrammingLanguageDTOManager) ToPLanguageDTO(appModel *domains.PLanguage) *UserPLanguageView {
	if appModel == nil {
		return nil
	}

	return &UserPLanguageView{
		ID:            appModel.ID,
		Name:          appModel.Name,
		Description:   appModel.Description,
		FileExtention: appModel.FileExtention,
		MonacoEditor:  appModel.MonacoEditor,
		CreatedAt:     appModel.CreatedAt,
	}
}

func (m *ProgrammingLanguageDTOManager) ToPLanguageDTOs(appModels []domains.PLanguage) []UserPLanguageView {
	var pLanguagesDTOs []UserPLanguageView
	for _, model := range appModels {
		pLanguagesDTOs = append(pLanguagesDTOs, *m.ToPLanguageDTO(&model))
	}
	return pLanguagesDTOs
}

type AddPLanguageDTO struct {
	LanguageID    string `json:"languageID"`
	Name          string `json:"name" validate:"required,max=30"`
	Description   string `json:"description"`
	ImagePath     string `json:"imagePath" validate:"required,max=60"`
	FileExtention string `json:"fileExtention" validate:"required,max=30"`
	MonacoEditor  string `json:"monacoEditor" validate:"required,max=30"`
}

type UpdatePLanguageDTO struct {
	ID            string `json:"id"`
	LanguageID    string `json:"languageID"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ImagePath     string `json:"imagePath"`
	FileExtention string `json:"fileExtention"`
	MonacoEditor  string `json:"monacoEditor"`
}
