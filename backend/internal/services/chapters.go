package services

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type chapterService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func NewChapterService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *chapterService {
	return &chapterService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *chapterService) GetChapters(
	ctx context.Context,
	id, langugeID, courseID, rewardID, title, grantsExperience, active, page, limit string,
) ([]repo.TChapter, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = s.utilService.D().Limits.DefaultChapterLimit
	}

	// Hata var ise dönsün diye
	if _, err := s.utilService.ParseUUID(id); err != nil {
		return nil, err
	}
	if _, err := s.utilService.ParseUUID(langugeID); err != nil {
		return nil, err
	}
	if _, err := s.utilService.ParseUUID(courseID); err != nil {
		return nil, err
	}
	if _, err := s.utilService.ParseUUID(rewardID); err != nil {
		return nil, err
	}

	chapters, err := s.queries.GetChapters(ctx, repo.GetChaptersParams{
		ID:         s.utilService.ParseNullUUID(id),
		LanguageID: s.utilService.ParseNullUUID(langugeID),
		RewardID:   s.utilService.ParseNullUUID(rewardID),
		CourseID:   s.utilService.ParseNullUUID(courseID),
		Title:      s.utilService.ParseString(title),
		Lim:        int32(limitNum),
		Off:        (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringChapter,
			err,
		)
	}

	return chapters, nil
}

func (s *chapterService) GetChapter(
	ctx context.Context,
	id, page, limit string,
) (*repo.TChapter, error) {
	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return nil, err
	}

	chapter, err := s.queries.GetChapterByID(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrUserNotFound)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUsers, err)
	}

	// TODO: Return tests with input and output by chapter id

	return &chapter, nil
}

func (s *chapterService) AddChapter(
	ctx context.Context,
	courseID, languageID, rewardID, title, description, content, funcName string,
	frontendTmp, dockerTmp, checkTmp string,
	grantsExperience, active bool,
	rewardAmount int,
) (uuid.UUID, error) {
	languageUUID, err := s.utilService.NParseUUID(languageID)
	if err != nil {
		return uuid.Nil, err
	}
	courseUUID, err := s.utilService.ParseUUID(courseID)
	if err != nil {
		return uuid.Nil, err
	}
	if _, err := s.utilService.ParseUUID(rewardID); err != nil {
		return uuid.Nil, err
	}

	id, err := s.queries.CreateChapter(ctx, repo.CreateChapterParams{
		LanguageID:       languageUUID,
		CourseID:         courseUUID,
		RewardID:         s.utilService.ParseNullUUID(rewardID),
		Title:            title,
		Description:      description,
		FuncName:         funcName,
		FrontendTemplate: frontendTmp,
		DockerTemplate:   dockerTmp,
		CheckTemplate:    checkTmp,
		RewardAmount:     int32(rewardAmount),
		GrantsExperience: grantsExperience,
		Active:           active,
	})
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
	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}

	if ok, err := s.queries.CheckChapterByID(ctx, idUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUsers, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrUserNotFound)
	}

	var rewAmountNullInt sql.NullInt32
	if rewardAmount == 0 {
		rewAmountNullInt.Valid = false
	} else {
		rewAmountNullInt.Valid = true
		rewAmountNullInt.Int32 = int32(rewardAmount)
	}

	var grantsExpNullBool sql.NullBool
	if grantsExperience {
		grantsExpNullBool.Valid = true
		grantsExpNullBool.Bool = true
	} else {
		grantsExpNullBool.Valid = false
	}

	var validNulBool sql.NullBool
	if active {
		validNulBool.Valid = true
		validNulBool.Bool = true
	} else {
		validNulBool.Valid = false
	}

	if err := s.queries.UpdateChapter(ctx, repo.UpdateChapterParams{
		ChapterID:        idUUID,
		LanguageID:       s.utilService.ParseNullUUID(languageID),
		CourseID:         s.utilService.ParseNullUUID(courseID),
		RewardID:         s.utilService.ParseNullUUID(rewardID),
		Title:            s.utilService.ParseString(title),
		Description:      s.utilService.ParseString(description),
		Content:          s.utilService.ParseString(content),
		FuncName:         s.utilService.ParseString(funcName),
		FrontendTemplate: s.utilService.ParseString(frontendTmp),
		DockerTemplate:   s.utilService.ParseString(dockerTmp),
		CheckTemplate:    s.utilService.ParseString(checkTmp),
		RewardAmount:     rewAmountNullInt,
		GrantsExperience: grantsExpNullBool,
		Active:           validNulBool,
	}); err != nil {
		return err
	}

	return nil
}

func (s *chapterService) DeleteChapter(
	ctx context.Context,
	id string,
) (err error) {
	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}

	if ok, err := s.queries.CheckChapterByID(ctx, idUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUsers, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrUserNotFound)
	}

	if err = s.queries.SoftDeleteChapter(ctx, idUUID); err != nil {
		return
	}
	return
}
