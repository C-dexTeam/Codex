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
}

type ICourseService interface {
	GetCourses(ctx context.Context, courseID, langugeID, pLanguageID, title, page, limit string) (courses []Course, err error)
}

const (
	DefaultCourseLimit = 10
)

type Course struct {
	id           uuid.UUID
	languageID   uuid.UUID
	pLanguageID  uuid.UUID
	rewardID     uuid.UUID
	rewardAmount int
	title        string
	description  string
	imagePath    string
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
	id, languageID, pLanguageID, rewardID uuid.UUID,
	rewardAmount int,
	title, description, imagePath string,
	createdAt, deletedAt time.Time,
) (course *Course, err error) {
	if err = course.SetTitle(title); err != nil {
		return nil, err
	}
	if err = course.SetImagePath(imagePath); err != nil {
		return nil, err
	}
	if err = course.SetRewardAmount(rewardAmount); err != nil {
		return nil, err
	}

	course.SetLanguageID(languageID)
	course.SetPLanguageID(pLanguageID)
	course.SetRewardID(rewardID)
	course.SetDescription(description)

	return
}

func (d *Course) Unmarshal(
	id, languageID, pLanguageID, rewardID uuid.UUID,
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

func (d *Course) GetRewardID() uuid.UUID {
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

func (d *Course) GetCreatedAt() time.Time {
	return d.createdAt
}

func (d *Course) GetDeletedAt() time.Time {
	return d.deletedAt
}

// Setter
func (d *Course) SetLanguageID(languageID uuid.UUID) {
	d.languageID = languageID
}

func (d *Course) SetPLanguageID(pLanguageID uuid.UUID) {
	d.pLanguageID = pLanguageID
}

func (d *Course) SetRewardID(rewardID uuid.UUID) {
	d.rewardID = rewardID
}

func (d *Course) SetRewardAmount(rewardAmount int) error {
	if rewardAmount > 0 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrCourseRewardAmountCannotBeNegative)
	}
	d.rewardAmount = rewardAmount

	return nil
}

func (d *Course) SetTitle(title string) error {
	if title == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrCourseTitleCannotBeEmpty)
	}
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
	if imagePath == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrCourseImagePathCannotBeEmpty)
	}
	if len(imagePath) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrCourseImagePathTooLong)

	}
	d.title = imagePath

	return nil
}
