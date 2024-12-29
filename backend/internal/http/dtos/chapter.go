package dto

import (
	"time"

	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type ChapterDTOManager struct{}

func NewChapterDTOManager() ChapterDTOManager {
	return ChapterDTOManager{}
}

type ChapterDTO struct {
	ID               uuid.UUID  `json:"id"`
	CourseID         uuid.UUID  `json:"courseID"`
	LanguageID       uuid.UUID  `json:"languageID"`
	RewardID         *uuid.UUID `json:"rewardID"`
	RewardAmount     int32      `json:"rewardAmount"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	Content          string     `json:"content"`
	FuncName         string     `json:"fundName"`
	FrontendTmp      string     `json:"frontendTemplate"`
	DockerTmp        string     `json:"dockerTemplate"`
	CheckTmp         string     `json:"check_template"`
	GrantsExperience bool       `json:"grantsExperience"`
	Active           bool       `json:"active"`
	Tests            []TestDTO  `json:"tests,omitempty"`
	CreatedAt        time.Time  `json:"createdAt"`
	DeletedAt        *time.Time `json:"deletedAt"`
}

func (d *ChapterDTOManager) ToChapterDTO(appModel repo.TChapter) ChapterDTO {
	// testManager := new(TestDTOManager)

	var rewardID *uuid.UUID
	if appModel.RewardID.Valid {
		r := uuid.MustParse(appModel.RewardID.UUID.String())
		rewardID = &r
	} else {
		rewardID = nil
	}

	return ChapterDTO{
		ID:               appModel.ID,
		CourseID:         appModel.CourseID,
		LanguageID:       appModel.LanguageID,
		RewardID:         rewardID,
		RewardAmount:     appModel.RewardAmount,
		Title:            appModel.Title,
		Description:      appModel.Description,
		Content:          appModel.Content,
		FuncName:         appModel.FuncName,
		FrontendTmp:      appModel.FrontendTemplate,
		DockerTmp:        appModel.DockerTemplate,
		CheckTmp:         appModel.CheckTemplate,
		GrantsExperience: appModel.GrantsExperience,
		Active:           appModel.Active,
		// Tests:            testManager.ToTestDTOs(appModel.GetTests()),
		CreatedAt: appModel.CreatedAt.Time,
		DeletedAt: &appModel.DeletedAt.Time,
	}
}

func (d *ChapterDTOManager) ToChapterDTOs(appModels []repo.TChapter) []ChapterDTO {
	var chapterDTOs []ChapterDTO
	for _, model := range appModels {
		chapterDTOs = append(chapterDTOs, d.ToChapterDTO(model))
	}
	return chapterDTOs
}

type AddChapterDTO struct {
	CourseID         string `json:"courseID"`
	LanguageID       string `json:"languageID"`
	RewardID         string `json:"rewardID"`
	RewardAmount     int    `json:"rewardAmount" validate:"gte=1"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	Content          string `json:"content"`
	FuncName         string `json:"fundName"`
	FrontendTmp      string `json:"frontendTemplate"`
	DockerTmp        string `json:"dockerTemplate"`
	CheckTmp         string `json:"checkTemplate"`
	GrantsExperience bool   `json:"grantsExperience"`
	Active           bool   `json:"active"`
}

type UpdateChapterDTO struct {
	ID               string `json:"id"`
	CourseID         string `json:"courseID"`
	LanguageID       string `json:"languageID"`
	RewardID         string `json:"rewardID"`
	RewardAmount     int    `json:"rewardAmount" validate:"gte=1"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	Content          string `json:"content"`
	FuncName         string `json:"fundName"`
	FrontendTmp      string `json:"frontendTemplate"`
	DockerTmp        string `json:"dockerTemplate"`
	CheckTmp         string `json:"checkTemplate"`
	GrantsExperience bool   `json:"grantsExperience"`
	Active           bool   `json:"active"`
}
