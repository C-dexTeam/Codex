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

type courseService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func newCourseService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *courseService {
	return &courseService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *courseService) GetCourses(
	ctx context.Context,
	id, langugeID, pLanguageID, title, page, limit string,
) ([]repo.TCourse, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = s.utilService.D().Limits.DefaultCourseLimit
	}

	// Check if the uuid is wrong. If it's wrong return an error
	// Meybe this is not neccessariy. Because of nulluuid in sqlc.
	if _, err := s.utilService.ParseUUID(id); err != nil {
		return nil, err
	}
	if _, err := s.utilService.ParseUUID(langugeID); err != nil {
		return nil, err
	}
	if _, err := s.utilService.ParseUUID(pLanguageID); err != nil {
		return nil, err
	}

	courses, err := s.queries.GetCourses(ctx, repo.GetCoursesParams{
		ID:                    s.utilService.ParseNullUUID(id),
		LanguageID:            s.utilService.ParseNullUUID(langugeID),
		ProgrammingLanguageID: s.utilService.ParseNullUUID(pLanguageID),
		Title:                 s.utilService.ParseString(title),
		Lim:                   int32(limitNum),
		Off:                   (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringCourse,
			err,
		)
	}

	return courses, nil
}

func (s *courseService) GetCourse(
	ctx context.Context,
	id, page, limit string,
) (*repo.TCourse, []repo.TChapter, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = s.utilService.D().Limits.DefaultChapterLimit
	}

	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return nil, nil, err
	}

	course, err := s.queries.GetCourseByID(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, nil, serviceErrors.NewServiceErrorWithMessage(
				serviceErrors.StatusBadRequest,
				serviceErrors.ErrCourseNotFound,
			)
		}
		return nil, nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringCourse,
			err,
		)
	}

	courseChapters, err := s.queries.GetChapters(ctx, repo.GetChaptersParams{
		CourseID: s.utilService.ParseNullUUID(course.ID.String()),
		Lim:      int32(limitNum),
		Off:      (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringChapter,
			err,
		)
	}

	return &course, courseChapters, nil
}

func (s *courseService) AddCourse(
	ctx context.Context,
	languageID, pLanguageID, rewardID, title, description, imagePath string,
	rewardAmount int,
) (uuid.UUID, error) {
	languageUUID, err := s.utilService.NParseUUID(languageID)
	if err != nil {
		return uuid.Nil, err
	}
	if _, err := s.utilService.NParseUUID(pLanguageID); err != nil {
		return uuid.Nil, err
	}
	if _, err := s.utilService.ParseUUID(rewardID); err != nil {
		return uuid.Nil, err
	}

	id, err := s.queries.CreateCourse(ctx, repo.CreateCourseParams{
		LanguageID:            languageUUID,
		ProgrammingLanguageID: s.utilService.ParseNullUUID(pLanguageID),
		RewardID:              s.utilService.ParseNullUUID(rewardID),
		RewardAmount:          int32(rewardAmount),
		Title:                 title,
		Description:           description,
		ImagePath:             imagePath,
	})
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s *courseService) UpdateCourse(
	ctx context.Context,
	id, languageID, pLanguageID, rewardID, title, description, imagePath string,
	rewardAmount int,
) error {
	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}

	if ok, err := s.queries.CheckCourseByID(ctx, idUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringCourse, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrCourseNotFound)
	}

	var r sql.NullInt32
	if rewardAmount == 0 {
		r.Valid = false
	} else {
		r.Valid = true
		r.Int32 = int32(rewardAmount)
	}

	if err := s.queries.UpdateCourse(ctx, repo.UpdateCourseParams{
		CourseID:              idUUID,
		LanguageID:            s.utilService.ParseNullUUID(languageID),
		ProgrammingLanguageID: s.utilService.ParseNullUUID(pLanguageID),
		RewardID:              s.utilService.ParseNullUUID(rewardID),
		RewardAmount:          r,
		Title:                 s.utilService.ParseString(title),
		Description:           s.utilService.ParseString(description),
		ImagePath:             s.utilService.ParseString(imagePath),
	}); err != nil {
		return err
	}

	return nil
}

func (s *courseService) DeleteCourse(
	ctx context.Context,
	id string,
) (err error) {
	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}

	if ok, err := s.queries.CheckCourseByID(ctx, idUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringCourse, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrCourseNotFound)
	}

	if err := s.queries.SoftDeleteCourse(ctx, idUUID); err != nil {
		return err
	}
	return
}

func (s *courseService) StartCourse(
	ctx context.Context,
	id, userAuthID string,
) (uuid.UUID, error) {
	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return uuid.Nil, err
	}

	// Comes From Session
	userAuthUUID := uuid.MustParse(userAuthID)

	if ok, err := s.queries.CheckCourseByID(ctx, idUUID); err != nil {
		return uuid.Nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringCourse, err)
	} else if !ok {
		return uuid.Nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrCourseNotFound)
	}

	if err := s.queries.AddCourseToUser(ctx, repo.AddCourseToUserParams{
		CourseID:   idUUID,
		UserAuthID: userAuthUUID,
	}); err != nil {
		return uuid.Nil, err
	}

	return userAuthUUID, nil
}
