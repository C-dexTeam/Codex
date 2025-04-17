package domains

import (
	"time"

	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type Course struct {
	ID           uuid.UUID
	LanguageID   uuid.UUID
	PLanguageID  uuid.UUID
	RewardID     *uuid.UUID
	Title        string
	Description  string
	ImagePath    string
	Chapters     []Chapter
	ChapterCount int64
	PLanguage    *PLanguage
	CreatedAt    time.Time
	DeletedAt    *time.Time
}

type Courses struct {
	Courses     []Course
	TotalCourse int64
}

func NewCoursesS(
	courses []Course,
	count int64,
) *Courses {
	return &Courses{
		Courses:     courses,
		TotalCourse: count,
	}
}

func NewCourse(
	course *repo.GetCourseRow,
	chapters []repo.TChapter,
	pLanguage *repo.TProgrammingLanguage,
) Course {
	var rewardID *uuid.UUID
	if course.RewardID.Valid {
		r := uuid.MustParse(course.RewardID.UUID.String())
		rewardID = &r
	} else {
		rewardID = nil
	}
	var deletedAt *time.Time
	if course.DeletedAt.Valid {
		deletedAt = &course.DeletedAt.Time
	} else {
		deletedAt = nil
	}

	return Course{
		ID:           course.ID,
		LanguageID:   course.LanguageID,
		PLanguageID:  course.ProgrammingLanguageID.UUID,
		RewardID:     rewardID,
		Title:        course.Title,
		Description:  course.Description,
		ImagePath:    course.ImagePath.String,
		PLanguage:    NewPLanguage(pLanguage),
		Chapters:     NewChapters(chapters),
		ChapterCount: course.ChapterCount,
		CreatedAt:    course.CreatedAt.Time,
		DeletedAt:    deletedAt,
	}
}

func NewGetCoursesRow(course repo.GetTopCoursesRow) *repo.GetCoursesRow {
	return &repo.GetCoursesRow{
		ID:                    course.ID,
		LanguageID:            course.LanguageID,
		ProgrammingLanguageID: course.ProgrammingLanguageID,
		RewardID:              course.RewardID,
		Title:                 course.Title,
		Description:           course.Description,
		ImagePath:             course.ImagePath,
		ChapterCount:          course.ChapterCount,
		CreatedAt:             course.CreatedAt,
		DeletedAt:             course.DeletedAt,
	}
}

func ToGetCoursesRow(courses []repo.GetTopCoursesRow) []repo.GetCoursesRow {
	result := make([]repo.GetCoursesRow, len(courses))
	for i, course := range courses {
		result[i] = *NewGetCoursesRow(course)
	}
	return result
}

func NewCourses(
	courses []repo.GetCoursesRow,
) []Course {
	var domainCourses []Course
	for _, course := range courses {
		var rewardID *uuid.UUID
		if course.RewardID.Valid {
			r := uuid.MustParse(course.RewardID.UUID.String())
			rewardID = &r
		} else {
			rewardID = nil
		}
		var deletedAt *time.Time
		if course.DeletedAt.Valid {
			deletedAt = &course.DeletedAt.Time
		} else {
			deletedAt = nil
		}

		newCourse := Course{
			ID:           course.ID,
			LanguageID:   course.LanguageID,
			PLanguageID:  course.ProgrammingLanguageID.UUID,
			RewardID:     rewardID,
			Title:        course.Title,
			Description:  course.Description,
			ImagePath:    course.ImagePath.String,
			ChapterCount: course.ChapterCount,
			CreatedAt:    course.CreatedAt.Time,
			DeletedAt:    deletedAt,
		}

		domainCourses = append(domainCourses, newCourse)
	}

	return domainCourses
}
