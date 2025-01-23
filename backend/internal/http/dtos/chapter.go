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

type UserChapterView struct {
	ID               uuid.UUID  `json:"id"`
	CourseID         uuid.UUID  `json:"courseID"`
	RewardID         *uuid.UUID `json:"rewardID"`
	RewardAmount     int32      `json:"rewardAmount"`
	RewardImage      string     `json:"rewardImage"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	Content          string     `json:"content"`
	GrantsExperience bool       `json:"grantsExperience"`
	Active           bool       `json:"active"`
	Tests            []TestView `json:"tests,omitempty"`
	CreatedAt        time.Time  `json:"createdAt"`
}

func (d *ChapterDTOManager) ToChapterDTO(appModel *domains.Chapter) UserChapterView {
	testManager := new(TestDTOManager)

	return UserChapterView{
		ID:               appModel.ID,
		CourseID:         appModel.CourseID,
		RewardID:         appModel.RewardID,
		RewardAmount:     appModel.RewardAmount,
		Title:            appModel.Title,
		Description:      appModel.Description,
		Content:          appModel.Content,
		GrantsExperience: appModel.GrantsExp,
		RewardImage:      appModel.Reward.ImagePath,
		Active:           appModel.Active,
		Tests:            testManager.ToTestDTOs(appModel.Tests),
		CreatedAt:        appModel.CreatedAt,
	}
}

func (d *ChapterDTOManager) ToChapterDTOs(appModels []domains.Chapter) []UserChapterView {
	var chapterDTOs []UserChapterView
	for _, model := range appModels {
		chapterDTOs = append(chapterDTOs, d.ToChapterDTO(&model))
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
	RewardAmount     int    `json:"rewardAmount"`
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

type RunChapter struct {
	ChapterID string `json:"chapterID" validate:"required,uuid4"`
	CourseID  string `json:"courseID" validate:"required,uuid4"`
	UserCode  string `json:"userCode"`
}
