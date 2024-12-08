package domains

import (
	"context"
	"time"

	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type ICourseRepository interface {
	Filter(ctx context.Context, filter CourseFilter, limit, page int64) (courses []Course, dataCount int64, err error)
	Add(ctx context.Context, course *Course) (uuid.UUID, error)
	Update(ctx context.Context, course *Course) (err error)
	SoftDelete(ctx context.Context, id uuid.UUID) (err error)
}

type ICourseService interface {
	GetCourses(ctx context.Context, courseID, langugeID, pLanguageID, title, page, limit string) (courses []Course, err error)
	GetCourse(ctx context.Context, id, page, limit string) (course *Course, err error)
	AddCourse(ctx context.Context, languageID, pLanguageID, rewardID, title, description, imagePath string, rewardAmount int) (uuid.UUID, error)
	UpdateCourse(ctx context.Context, id, languageID, pLanguageID, rewardID, title, description, imagePath string, rewardAmount int) error
	DeleteCourse(ctx context.Context, id string) (err error)
}

const (
	DefaultCourseLimit = 10
)

type Course struct {
	id           uuid.UUID
	languageID   uuid.UUID
	pLanguageID  uuid.UUID
	rewardID     *uuid.UUID
	rewardAmount int
	title        string
	description  string
	imagePath    string
	chapters     []Chapter
	createdAt    time.Time
	deletedAt    time.Time
}

type CourseFilter struct {
	ID          uuid.UUID
	LanguageID  uuid.UUID
	PLanguageID uuid.UUID
	Title       string
	CreatedAt   time.Time
}

func NewCourse(
	id, languageID, pLanguageID, rewardID string,
	rewardAmount int,
	title, description, imagePath string,
	chapters []Chapter,
) (course *Course, err error) {
	course = &Course{}
	if err := course.SetID(id); err != nil {
		return nil, err
	}
	if err := course.SetLanguageID(languageID); err != nil {
		return nil, err
	}
	if err := course.SetPLanguageID(pLanguageID); err != nil {
		return nil, err
	}
	if err := course.SetRewardID(rewardID); err != nil {
		return nil, err
	}
	if err = course.SetTitle(title); err != nil {
		return nil, err
	}
	if err = course.SetImagePath(imagePath); err != nil {
		return nil, err
	}
	if err = course.SetRewardAmount(rewardAmount); err != nil {
		return nil, err
	}

	course.SetDescription(description)
	course.SetChapters(chapters)

	return
}

func (d *Course) Unmarshal(
	id, languageID, pLanguageID uuid.UUID,
	rewardID *uuid.UUID,
	rewardAmount int,
	title, description, imagePath string,
	createdAt, deletedAt time.Time,
) {
	d.id = id
	d.languageID = languageID
	d.pLanguageID = pLanguageID
	d.rewardID = rewardID
	d.rewardAmount = rewardAmount
	d.title = title
	d.description = description
	d.imagePath = imagePath
	d.createdAt = createdAt
	d.deletedAt = deletedAt
}

// Getter
func (d *Course) GetID() uuid.UUID {
	return d.id
}

func (d *Course) GetLanguageID() uuid.UUID {
	return d.languageID
}

func (d *Course) GetPLanguageID() uuid.UUID {
	return d.pLanguageID
}

func (d *Course) GetRewardID() *uuid.UUID {
	return d.rewardID
}

func (d *Course) GetRewardAmount() int {
	return d.rewardAmount
}

func (d *Course) GetTitle() string {
	return d.title
}

func (d *Course) GetDescription() string {
	return d.description
}

func (d *Course) GetImagePath() string {
	return d.imagePath
}

func (d *Course) GetChapters() []Chapter {
	return d.chapters
}

func (d *Course) GetCreatedAt() time.Time {
	return d.createdAt
}

func (d *Course) GetDeletedAt() time.Time {
	return d.deletedAt
}

// Setter
func (d *Course) SetID(id string) error {
	if id != "" {
		idUUID, err := uuid.Parse(id)
		if err != nil {
			return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
		}
		d.id = idUUID
	}

	return nil
}

func (d *Course) SetLanguageID(languageID string) error {
	idUUID, err := uuid.Parse(languageID)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
	}
	d.languageID = idUUID

	return nil
}

func (d *Course) SetPLanguageID(pLanguageID string) error {
	idUUID, err := uuid.Parse(pLanguageID)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
	}
	d.pLanguageID = idUUID

	return nil
}

func (d *Course) SetRewardID(rewardID string) error {
	if rewardID == "" {
		d.rewardID = nil
	} else {
		idUUID, err := uuid.Parse(rewardID)
		if err != nil {
			return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
		}
		d.rewardID = &idUUID
	}

	return nil
}

func (d *Course) SetRewardAmount(rewardAmount int) (err error) {
	if rewardAmount < 0 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrCourseRewardAmountCannotBeNegative)
	}
	d.rewardAmount = rewardAmount

	return nil
}

func (d *Course) SetTitle(title string) error {
	if len(title) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrCourseTitleTooLong)

	}
	d.title = title

	return nil
}

func (d *Course) SetDescription(description string) {
	d.description = description
}

func (d *Course) SetImagePath(imagePath string) error {
	if len(imagePath) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrCourseImagePathTooLong)

	}
	d.imagePath = imagePath

	return nil
}

func (d *Course) SetChapters(chapters []Chapter) {
	d.chapters = chapters
}
