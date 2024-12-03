package services

import (
	"context"
	"strconv"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type chapterService struct {
	chapterRepository domains.IChapterRepository
}

func NewChapterService(
	chapterRepository domains.IChapterRepository,
) domains.IChapterService {
	return &chapterService{
		chapterRepository: chapterRepository,
	}
}

func (s *chapterService) GetChapters(ctx context.Context, chapterID, langugeID, courseID, rewardID, title, grantsExperience, active, page, limit string) (chapters []domains.Chapter, err error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = domains.DefaultChapterLimit
	}

	var (
		chapterUUID   uuid.UUID
		languageUUID  uuid.UUID
		courseUUID    uuid.UUID
		rewardUUID    uuid.UUID
		grantsExpBool bool
		activeBool    bool
	)
	if chapterID != "" {
		chapterUUID, err = uuid.Parse(courseID)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}
	if langugeID != "" {
		languageUUID, err = uuid.Parse(langugeID)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}
	if courseID != "" {
		courseUUID, err = uuid.Parse(courseID)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}
	if rewardID != "" {
		rewardUUID, err = uuid.Parse(rewardID)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}
	if grantsExperience != "" {
		grantsExpBool, err = strconv.ParseBool(grantsExperience)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidBoolean, err)
		}
	}
	if active != "" {
		activeBool, err = strconv.ParseBool(active)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidBoolean, err)
		}
	}

	chapters, _, err = s.chapterRepository.Filter(ctx, domains.ChapterFilter{
		ID:               chapterUUID,
		LanguageID:       languageUUID,
		CourseID:         courseUUID,
		RewardID:         rewardUUID,
		Title:            title,
		GrantsExperience: grantsExpBool,
		Active:           activeBool,
	}, int64(limitNum), int64(pageNum))
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringChapter, err)
	}

	return
}
