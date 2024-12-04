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
	CourseID         *uuid.UUID `json:"course_id"`
	LanguageID       *uuid.UUID `json:"language_id"`
	RewardID         *uuid.UUID `json:"reward_id"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	Content          string     `json:"content"`
	FuncName         string     `json:"func_name"`
	FrontendTmp      string     `json:"frontend_template"`
	DockerTmp        string     `json:"docker_template"`
	CheckTmp         string     `json:"check_template"`
	GrantsExperience bool       `json:"grants_experience"`
	Active           bool       `json:"active"`
	CreatedAt        time.Time  `json:"created_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
}

func (d *ChapterDTOManager) ToChapterDTO(appModel domains.Chapter) ChapterDTO {
	return ChapterDTO{
		ID:               appModel.GetID(),
		CourseID:         appModel.GetCourseID(),
		LanguageID:       appModel.GetLanguageID(),
		RewardID:         appModel.GetRewardID(),
		Title:            appModel.GetTitle(),
		Description:      appModel.GetDescription(),
		Content:          appModel.GetContent(),
		FuncName:         appModel.GetFuncName(),
		FrontendTmp:      appModel.GetFrontendTmp(),
		DockerTmp:        appModel.GetDockerTmp(),
		CheckTmp:         appModel.GetCheckTmp(),
		GrantsExperience: appModel.GetGrantsExperience(),
		Active:           appModel.GetActive(),
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
