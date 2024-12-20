package services

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
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
		limitNum = domains.DefaultChapterLimit
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
	// if grantsExperience != "" {
	// 	grantsExpBoolValue, err := strconv.ParseBool(grantsExperience)
	// 	if err != nil {
	// 		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
	// 			errorDomains.StatusBadRequest,
	// 			errorDomains.ErrInvalidBoolean,
	// 			err,
	// 		)
	// 	}
	// 	grantsExpBool = &grantsExpBoolValue
	// }
	// if active != "" {
	// 	activeBoolValue, err := strconv.ParseBool(active)
	// 	if err != nil {
	// 		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
	// 			errorDomains.StatusBadRequest,
	// 			errorDomains.ErrInvalidBoolean,
	// 			err,
	// 		)
	// 	}
	// 	activeBool = &activeBoolValue
	// }

	chapters, err := s.queries.GetChapters(ctx, repo.GetChaptersParams{
		ID:         s.utilService.ParseNullUUID(id),
		LanguageID: s.utilService.ParseNullUUID(langugeID),
		RewardID:   s.utilService.ParseNullUUID(rewardID),
		CourseID:   s.utilService.ParseNullUUID(courseID),
		Title:      s.utilService.ParseString(title),
		Lim:        int32(limitNum),
		Off:        (int32(pageNum) - 1) + int32(limitNum),
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			errorDomains.StatusInternalServerError,
			errorDomains.ErrErrorWhileFilteringChapter,
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
			return nil, serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrUserNotFound)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}

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
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrUserNotFound)
	}

	var r sql.NullInt32
	if rewardAmount == 0 {
		r.Valid = false
	} else {
		r.Valid = true
		r.Int32 = int32(rewardAmount)
	}

	var g sql.NullBool
	if grantsExperience {
		g.Valid = true
		g.Bool = true
	} else {
		g.Valid = false
	}

	var a sql.NullBool
	if active {
		a.Valid = true
		a.Bool = true
	} else {
		a.Valid = false
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
		RewardAmount:     r,
		GrantsExperience: g,
		Active:           a,
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
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrUserNotFound)
	}

	if err = s.queries.SoftDeleteChapter(ctx, idUUID); err != nil {
		return
	}
	return
}
