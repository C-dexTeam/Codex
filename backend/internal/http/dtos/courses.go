package dto

import (
	"time"

	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type CourseDTOManager struct{}

func NewCourseDTOManager() CourseDTOManager {
	return CourseDTOManager{}
}

type CourseView struct {
	ID           uuid.UUID    `json:"id"`
	LanguageID   uuid.UUID    `json:"languageID"`
	PLanguageID  uuid.UUID    `json:"programmingLanguageID"`
	RewardID     *uuid.UUID   `json:"rewardID"`
	RewardAmount int32        `json:"rewardAmount"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	ImagePath    string       `json:"imagePath"`
	CreatedAt    time.Time    `json:"createdAt"`
	DeletedAt    *time.Time   `json:"deletedAt"`
	Chapters     []ChapterDTO `json:"chapters,omitempty"`
}

func (d *CourseDTOManager) ToCourseDTO(courseModel *repo.TCourse, chapterModels []repo.TChapter) CourseView {
	chapterDTOManager := new(ChapterDTOManager)
	var rewardID *uuid.UUID
	if courseModel.RewardID.Valid {
		r := uuid.MustParse(courseModel.RewardID.UUID.String())
		rewardID = &r
	} else {
		rewardID = nil
	}
	var deletedAt *time.Time
	if courseModel.DeletedAt.Valid {
		deletedAt = &courseModel.DeletedAt.Time
	} else {
		deletedAt = nil
	}

	return CourseView{
		ID:           courseModel.ID,
		LanguageID:   courseModel.LanguageID,
		PLanguageID:  courseModel.ProgrammingLanguageID.UUID,
		RewardID:     rewardID,
		RewardAmount: courseModel.RewardAmount,
		Title:        courseModel.Title,
		Description:  courseModel.Description,
		ImagePath:    courseModel.ImagePath,
		CreatedAt:    courseModel.CreatedAt.Time,
		DeletedAt:    deletedAt,
		Chapters:     chapterDTOManager.ToChapterDTOs(chapterModels),
	}
}

func (d *CourseDTOManager) ToCourseDTOs(courseModels []repo.TCourse) []CourseView {
	var courseDTOs []CourseView
	for _, model := range courseModels {
		courseDTOs = append(courseDTOs, d.ToCourseDTO(&model, nil))
	}

	return courseDTOs
}

type AddCourseDTO struct {
	LanguageID   string `json:"languageID"`
	PLanguageID  string `json:"programmingLanguageID" validate:"required,uuid4"`
	RewardID     string `json:"rewardID"`
	RewardAmount int    `json:"rewardAmount" validate:"gte=1"`
	Title        string `json:"title" validate:"required,max=60"`
	Description  string `json:"description"`
	ImagePath    string `json:"imagePath"`
}

type UpdateCourseDTO struct {
	ID           string `json:"id"`
	LanguageID   string `json:"languageID"`
	PLanguageID  string `json:"programmingLanguageID"`
	RewardID     string `json:"rewardID"`
	RewardAmount int    `json:"rewardAmount" validate:"gte=1"`
	Title        string `json:"title" validate:"max=60"`
	Description  string `json:"description"`
	ImagePath    string `json:"imagePath"`
}

type StartCourseDTO struct {
	ID string `json:"id" validate:"required,uudi4"`
}
