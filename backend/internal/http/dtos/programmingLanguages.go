package dto

import (
	"time"

	repo "github.com/C-dexTeam/codex/internal/repos/out"
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

func (m *ProgrammingLanguageDTOManager) ToPLanguageDTO(appModel *repo.TProgrammingLanguage) ProgrammingLanguageDTO {
	return ProgrammingLanguageDTO{
		ID:            appModel.ID,
		LanguageID:    appModel.LanguageID,
		Name:          appModel.Name,
		Description:   appModel.Description,
		DownloadCMD:   appModel.DownloadCmd,
		CompileCMD:    appModel.CompileCmd,
		FileExtention: appModel.FileExtention,
		MonacoEditor:  appModel.MonacoEditor,
		CreatedAt:     appModel.CreatedAt,
	}
}

func (m *ProgrammingLanguageDTOManager) ToPLanguageDTOs(appModels []repo.TProgrammingLanguage) []ProgrammingLanguageDTO {
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
