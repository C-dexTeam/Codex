package dto

import (
	"time"

	"github.com/C-dexTeam/codex/internal/domains"
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
	RewardAmount     int        `json:"rewardAmount"`
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

func (d *ChapterDTOManager) ToChapterDTO(appModel domains.Chapter) ChapterDTO {
	testManager := new(TestDTOManager)

	return ChapterDTO{
		ID:               appModel.GetID(),
		CourseID:         appModel.GetCourseID(),
		LanguageID:       appModel.GetLanguageID(),
		RewardID:         appModel.GetRewardID(),
		RewardAmount:     appModel.GetRewardAmount(),
		Title:            appModel.GetTitle(),
		Description:      appModel.GetDescription(),
		Content:          appModel.GetContent(),
		FuncName:         appModel.GetFuncName(),
		FrontendTmp:      appModel.GetFrontendTmp(),
		DockerTmp:        appModel.GetDockerTmp(),
		CheckTmp:         appModel.GetCheckTmp(),
		GrantsExperience: appModel.GetGrantsExperience(),
		Active:           appModel.GetActive(),
		Tests:            testManager.ToTestDTOs(appModel.GetTests()),
		CreatedAt:        appModel.GetCreatedAt(),
		DeletedAt:        appModel.GetDeletedAt(),
	}
}

func (d *ChapterDTOManager) ToChapterDTOs(appModels []domains.Chapter) []ChapterDTO {
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
