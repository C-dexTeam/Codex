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

func (s *chapterService) GetChapters(
	ctx context.Context,
	id, langugeID, courseID, rewardID, title, grantsExperience, active, page, limit string,
) (chapters []domains.Chapter, err error) {
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
		grantsExpBool *bool
		activeBool    *bool
	)
	if id != "" {
		chapterUUID, err = uuid.Parse(id)
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
		grantsExpBoolValue, err := strconv.ParseBool(grantsExperience)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidBoolean, err)
		}
		grantsExpBool = &grantsExpBoolValue
	}
	if active != "" {
		activeBoolValue, err := strconv.ParseBool(active)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidBoolean, err)
		}
		activeBool = &activeBoolValue
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

func (s *chapterService) AddChapter(
	ctx context.Context,
	courseID, languageID, rewardID, title, description, content, funcName string,
	frontendTmp, dockerTmp, checkTmp string,
	grantsExperience, active bool,
	rewardAmount int,
) (uuid.UUID, error) {
	newChapter, err := domains.NewChapter(
		"",
		languageID,
		courseID,
		rewardID,
		title,
		description,
		content,
		funcName,
		frontendTmp,
		dockerTmp,
		checkTmp,
		rewardAmount,
		grantsExperience,
		active,
	)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := s.chapterRepository.Add(ctx, newChapter)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s *chapterService) UpdateChapter(
	ctx context.Context,
	id, courseID, languageID, rewardID, title, description, content, funcName string,
	frontendTmp, dockerTmp, checkTmp string,
	grantsExperience, active bool,
	rewardAmount int,
) error {
	var idUUID uuid.UUID
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	rewards, _, err := s.chapterRepository.Filter(ctx, domains.ChapterFilter{
		ID: idUUID,
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringChapter, err)
	}
	if len(rewards) != 1 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusNotFound, errorDomains.ErrChapterNotFound)
	}

	updateChapter, err := domains.NewChapter(
		id,
		languageID,
		courseID,
		rewardID,
		title,
		description,
		content,
		funcName,
		frontendTmp,
		dockerTmp,
		checkTmp,
		rewardAmount,
		grantsExperience,
		active,
	)
	if err != nil {
		return err
	}

	if err := s.chapterRepository.Update(ctx, updateChapter); err != nil {
		return err
	}

	return nil
}

func (s *chapterService) DeleteChapter(
	ctx context.Context,
	id string,
) (err error) {
	var idUUID uuid.UUID
	idUUID, err = uuid.Parse(id)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	chapters, _, err := s.chapterRepository.Filter(ctx, domains.ChapterFilter{
		ID: idUUID,
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringChapter, err)
	}
	if len(chapters) != 1 {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrChapterNotFound, err)
	}

	if err = s.chapterRepository.SoftDelete(ctx, idUUID); err != nil {
		return
	}
	return
}
