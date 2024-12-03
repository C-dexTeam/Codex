package domains

import (
	"time"

	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type IChapterRepository interface{}

type IChapterService interface{}

type Chapter struct {
	id               uuid.UUID
	languageID       *uuid.UUID
	courseID         *uuid.UUID
	rewardID         *uuid.UUID
	title            string
	description      string
	content          string
	funcName         string
	frontendTmp      string
	dockerTmp        string
	checkTmp         string
	grantsExperience bool
	active           bool
	createdAt        time.Time
	deletedAt        time.Time
}

type ChapterFilter struct {
	ID               uuid.UUID
	LanguageID       *uuid.UUID
	CourseID         *uuid.UUID
	RewardID         *uuid.UUID
	Title            string
	GrantsExperience bool
	Active           bool
}

func NewChapter(
	id uuid.UUID,
	languageID, courseID, rewardID *uuid.UUID,
	title, description, content, funcName, frontendTmp, dockerTmp, checkTmp string,
	grantsExperience, active bool,
	createdAt, deletedAt time.Time,
) (chapter *Chapter, err error) {
	if err = chapter.SetTitle(title); err != nil {
		return
	}
	if err = chapter.SetFuncName(funcName); err != nil {
		return
	}

	chapter.SetLanguageID(languageID)
	chapter.SetCourseID(courseID)
	chapter.SetRewardID(rewardID)
	chapter.SetDescription(description)
	chapter.SetContent(content)
	chapter.SetFrontendTmp(frontendTmp)
	chapter.SetDockerTmp(dockerTmp)
	chapter.SetCheckTmp(checkTmp)
	chapter.SetGrantsExperience(grantsExperience)
	chapter.SetActive(active)
	chapter.SetDeletedAt(deletedAt)

	return
}

func (c *Chapter) Unmarshal(
	id uuid.UUID,
	languageID, courseID, rewardID *uuid.UUID,
	title, description, content, funcName, frontendTmp, dockerTmp, checkTmp string,
	grantsExperience, active bool,
	createdAt, deletedAt time.Time,
) {
	c.id = id
	c.languageID = languageID
	c.courseID = courseID
	c.rewardID = rewardID
	c.title = title
	c.description = description
	c.content = content
	c.funcName = funcName
	c.frontendTmp = frontendTmp
	c.dockerTmp = dockerTmp
	c.checkTmp = checkTmp
	c.grantsExperience = grantsExperience
	c.active = active
	c.createdAt = createdAt
	c.deletedAt = deletedAt
}

// Getter
func (c *Chapter) GetID() uuid.UUID {
	return c.id
}

func (c *Chapter) GetCourseID() *uuid.UUID {
	return c.courseID
}

func (c *Chapter) GetLanguageID() *uuid.UUID {
	return c.languageID
}

func (c *Chapter) GetRewardID() *uuid.UUID {
	return c.rewardID
}

func (c *Chapter) GetTitle() string {
	return c.title
}

func (c *Chapter) GetDescription() string {
	return c.description
}

func (c *Chapter) GetContent() string {
	return c.content
}

func (c *Chapter) GetFuncName() string {
	return c.funcName
}

func (c *Chapter) GetFrontendTmp() string {
	return c.frontendTmp
}

func (c *Chapter) GetDockerTmp() string {
	return c.dockerTmp
}

func (c *Chapter) GetCheckTmp() string {
	return c.checkTmp
}

func (c *Chapter) GetGrantsExperience() bool {
	return c.grantsExperience
}

func (c *Chapter) GetActive() bool {
	return c.active
}

func (c *Chapter) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *Chapter) GetDeletedAt() time.Time {
	return c.deletedAt
}

// Setter
func (c *Chapter) SetCourseID(courseID *uuid.UUID) {
	c.courseID = courseID
}

func (c *Chapter) SetLanguageID(languageID *uuid.UUID) {
	c.languageID = languageID
}

func (c *Chapter) SetRewardID(rewardID *uuid.UUID) {
	c.rewardID = rewardID
}

func (c *Chapter) SetTitle(title string) error {
	if title == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrChapterTitleCannotBeEmpty)
	}
	if len(title) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrChapterTitleTooLong)
	}
	c.title = title
	return nil
}

func (c *Chapter) SetDescription(description string) {
	c.description = description
}

func (c *Chapter) SetContent(content string) {
	c.content = content
}

func (c *Chapter) SetFuncName(funcName string) error {
	if funcName == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrChapterFuncNameCannotBeEmpty)
	}
	if len(funcName) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrChapterFuncNameTooLong)
	}
	c.funcName = funcName
	return nil
}

func (c *Chapter) SetFrontendTmp(frontendTmp string) {
	c.frontendTmp = frontendTmp
}

func (c *Chapter) SetDockerTmp(dockerTmp string) {
	c.dockerTmp = dockerTmp
}

func (c *Chapter) SetCheckTmp(checkTmp string) {
	c.checkTmp = checkTmp
}

func (c *Chapter) SetGrantsExperience(grantsExperience bool) {
	c.grantsExperience = grantsExperience
}

func (c *Chapter) SetActive(active bool) {
	c.active = active
}

func (c *Chapter) SetCreatedAt(createdAt time.Time) {
	c.createdAt = createdAt
}

func (c *Chapter) SetDeletedAt(deletedAt time.Time) {
	c.deletedAt = deletedAt
}
