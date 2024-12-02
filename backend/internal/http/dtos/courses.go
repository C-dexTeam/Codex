package dto

import (
	"time"

	"github.com/C-dexTeam/codex/internal/domains"
	"github.com/google/uuid"
)

type CourseDTOManager struct{}

func NewCourseDTOManager() CourseDTOManager {
	return CourseDTOManager{}
}

type CourseDTO struct {
	ID           uuid.UUID  `json:"id"`
	LanguageID   *uuid.UUID `json:"languageID"`
	PLanguageID  *uuid.UUID `json:"programmingLanguageID"`
	RewardID     *uuid.UUID `json:"rewardID"`
	RewardAmount int        `json:"rewardAmount"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	ImagePath    string     `json:"imagePath"`
	CreatedAt    time.Time  `json:"createdAt"`
	DeletedAt    time.Time  `json:"deletedAt"`
}

func (d *CourseDTOManager) ToCourseDTO(appModel domains.Course) CourseDTO {
	return CourseDTO{
		ID:           appModel.GetID(),
		LanguageID:   appModel.GetLanguageID(),
		PLanguageID:  appModel.GetPLanguageID(),
		RewardID:     appModel.GetRewardID(),
		RewardAmount: appModel.GetRewardAmount(),
		Title:        appModel.GetTitle(),
		Description:  appModel.GetDescription(),
		ImagePath:    appModel.GetImagePath(),
		CreatedAt:    appModel.GetCreatedAt(),
		DeletedAt:    appModel.GetDeletedAt(),
	}
}

func (d *CourseDTOManager) ToCourseDTOs(appModels []domains.Course) []CourseDTO {
	var courseDTOs []CourseDTO
	for _, model := range appModels {
		courseDTOs = append(courseDTOs, d.ToCourseDTO(model))
	}

	return courseDTOs
}
