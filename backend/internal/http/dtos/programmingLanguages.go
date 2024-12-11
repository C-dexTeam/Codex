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

type ProgrammingLanguageDTO struct {
	ID            uuid.UUID `json:"id"`
	LanguageID    uuid.UUID `json:"languageID"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	DownloadCMD   string    `json:"downloadCMD"`
	CompileCMD    string    `json:"compileCMD"`
	FileExtention string    `json:"fileExtention"`
	MonacoEditor  string    `json:"monacoEditor"`
	CreatedAt     time.Time `json:"createdAt"`
}

func (m *ProgrammingLanguageDTOManager) ToPLanguageDTO(appModel *domains.ProgrammingLanguage) ProgrammingLanguageDTO {
	return ProgrammingLanguageDTO{
		ID:            appModel.GetID(),
		LanguageID:    appModel.GetLanguageID(),
		Name:          appModel.GetName(),
		Description:   appModel.GetDescription(),
		DownloadCMD:   appModel.GetDownloadCMD(),
		CompileCMD:    appModel.GetCompileCMD(),
		FileExtention: appModel.GetFileExtention(),
		MonacoEditor:  appModel.GetMonacoEditor(),
		CreatedAt:     appModel.GetCreatedAt(),
	}
}

func (m *ProgrammingLanguageDTOManager) ToPLanguageDTOs(appModels []domains.ProgrammingLanguage) []ProgrammingLanguageDTO {
	var pLanguagesDTOs []ProgrammingLanguageDTO
	for _, model := range appModels {
		pLanguagesDTOs = append(pLanguagesDTOs, m.ToPLanguageDTO(&model))
	}
	return pLanguagesDTOs
}

type AddPLanguageDTO struct {
	LanguageID    string `json:"languageID"`
	Name          string `json:"name" validate:"required,max=30"`
	Description   string `json:"description"`
	DownloadCMD   string `json:"downloadCMD"`
	CompileCMD    string `json:"compileCMD"`
	ImagePath     string `json:"imagePath" validate:"required,max=60"`
	FileExtention string `json:"fileExtention" validate:"required,max=30"`
	MonacoEditor  string `json:"monacoEditor" validate:"required,max=30"`
}

type UpdatePLanguageDTO struct {
	ID            string `json:"id"`
	LanguageID    string `json:"languageID"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	DownloadCMD   string `json:"downloadCMD"`
	CompileCMD    string `json:"compileCMD"`
	ImagePath     string `json:"imagePath"`
	FileExtention string `json:"fileExtention"`
	MonacoEditor  string `json:"monacoEditor"`
}
