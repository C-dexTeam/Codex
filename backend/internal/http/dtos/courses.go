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

type UserCourseView struct {
	ID           uuid.UUID          `json:"id"`
	RewardID     *uuid.UUID         `json:"rewardID"`
	Title        string             `json:"title"`
	Description  string             `json:"description"`
	ImagePath    string             `json:"imagePath"`
	ChapterCount int64              `json:"chapterCount"`
	Chapters     []UserChapterView  `json:"chapters,omitempty"`
	PLanguage    *UserPLanguageView `json:"programmingLanguage,omitempty"`
	CreatedAt    time.Time          `json:"createdAt"`
}

type UserCourseViews struct {
	Courses     []UserCourseView `json:"courses"`
	CourseCount int64            `json:"courseCount"`
}

func (d *CourseDTOManager) ToCourseDTO(courseModel *domains.Course) UserCourseView {
	chapterDTOManager := new(ChapterDTOManager)
	pLangDTOManager := new(ProgrammingLanguageDTOManager)

	return UserCourseView{
		ID:           courseModel.ID,
		RewardID:     courseModel.RewardID,
		Title:        courseModel.Title,
		Description:  courseModel.Description,
		ImagePath:    courseModel.ImagePath,
		ChapterCount: courseModel.ChapterCount,
		Chapters:     chapterDTOManager.ToChapterDTOs(courseModel.Chapters),
		PLanguage:    pLangDTOManager.ToPLanguageDTO(courseModel.PLanguage),
		CreatedAt:    courseModel.CreatedAt,
	}
}

func (d *CourseDTOManager) ToCourseDTOs(courseModels []domains.Course) []UserCourseView {
	var courseDTOs []UserCourseView
	for _, model := range courseModels {
		courseDTOs = append(courseDTOs, d.ToCourseDTO(&model))
	}

	return courseDTOs
}

func (d *CourseDTOManager) ToCourseDTOCount(courseModels *domains.Courses) UserCourseViews {
	var viewsCourse UserCourseViews
	var courseDTOs []UserCourseView
	for _, model := range courseModels.Courses {
		courseDTOs = append(courseDTOs, d.ToCourseDTO(&model))
	}

	viewsCourse.Courses = courseDTOs
	viewsCourse.CourseCount = courseModels.TotalCourse

	return viewsCourse
}

type AddCourseDTO struct {
	LanguageID            string
	ProgrammingLanguageID string `validate:"required,uuid4"`
	RewardID              string
	Title                 string `validate:"required,max=60"`
	Description           string `validte:"required"`
}

type UpdateCourseDTO struct {
	ID          string `json:"id"`
	LanguageID  string `json:"languageID"`
	PLanguageID string `json:"programmingLanguageID"`
	RewardID    string `json:"rewardID"`
	Title       string `json:"title" validate:"max=60"`
	Description string `json:"description"`
}

type StartCourseDTO struct {
	ID string `json:"id" validate:"required"`
}
