package services

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/C-dexTeam/codex/internal/domains"
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
) (*domains.Courses, error) {
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

	count, err := s.GetCourseCount(ctx)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringCourseCount)
	}

	domainCourses := domains.NewCourses(courses)

	return domains.NewCoursesS(domainCourses, count), nil
}

func (s *courseService) GetPopularCourses(
	ctx context.Context,
	page, limit string,
) ([]domains.Course, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = s.utilService.D().Limits.DefaultCourseLimit
	}

	courses, err := s.queries.GetTopCourses(ctx, repo.GetTopCoursesParams{
		Lim: int32(limitNum),
		Off: (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringCourse,
			err,
		)
	}
	domainCourses := domains.NewCourses(domains.ToGetCoursesRow(courses))

	return domainCourses, nil
}

func (s *courseService) GetCourse(
	ctx context.Context,
	id, page, limit string,
) (*domains.Course, error) {
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
		return nil, err
	}

	course, err := s.queries.GetCourse(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(
				serviceErrors.StatusBadRequest,
				serviceErrors.ErrCourseNotFound,
			)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
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
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringChapter,
			err,
		)
	}
	domainCourse := domains.NewCourse(&course, courseChapters, nil)

	return &domainCourse, nil
}

func (s *courseService) UserCourse(ctx context.Context, userAuthID, courseID string) (*repo.TUserCourse, error) {
	userAuthUUID := uuid.MustParse(userAuthID)
	courseUUID := uuid.MustParse(courseID)

	userCourse, err := s.queries.UserCourse(ctx, repo.UserCourseParams{
		CourseID:   courseUUID,
		UserAuthID: userAuthUUID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(
				serviceErrors.StatusBadRequest,
				serviceErrors.ErrUserCourseNotFound,
			)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringUserCourse,
			err,
		)
	}

	return &userCourse, nil
}

func (s *courseService) AddCourse(
	ctx context.Context,
	languageID, pLanguageID, rewardID, title, description, imagePath string,
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
		Title:                 title,
		Description:           description,
		ImagePath:             s.utilService.ParseString(imagePath),
	})
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s *courseService) UpdateCourse(
	ctx context.Context,
	id, languageID, pLanguageID, rewardID, title, description, imagePath string,
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

	if err := s.queries.UpdateCourse(ctx, repo.UpdateCourseParams{
		CourseID:              idUUID,
		LanguageID:            s.utilService.ParseNullUUID(languageID),
		ProgrammingLanguageID: s.utilService.ParseNullUUID(pLanguageID),
		RewardID:              s.utilService.ParseNullUUID(rewardID),
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

	if err := s.queries.DeleteCourse(ctx, idUUID); err != nil {
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

	// Check if the course already started. If its return an error.
	if ok, err := s.queries.CheckUserCourseByID(ctx, repo.CheckUserCourseByIDParams{
		CourseID:   idUUID,
		UserAuthID: userAuthUUID,
	}); err != nil {
		return uuid.Nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringCourse, err)
	} else if ok {
		return uuid.Nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrCourseAlreadyStarted)
	}

	if err := s.queries.AddCourseToUser(ctx, repo.AddCourseToUserParams{
		CourseID:   idUUID,
		UserAuthID: userAuthUUID,
	}); err != nil {
		return uuid.Nil, err
	}

	return userAuthUUID, nil
}

func (s *courseService) UpdateUserCourseProgress(
	ctx context.Context,
	userAuthID, courseID string,
) (int32, error) {
	userAuthUUD := uuid.MustParse(userAuthID)
	courseUUID := uuid.MustParse(courseID)

	progress, err := s.queries.UpdateUserCourseProgress(ctx, repo.UpdateUserCourseProgressParams{
		UserAuthID: userAuthUUD,
		CourseID:   courseUUID,
	})
	if err != nil {
		return 0, err
	}

	if progress.Valid {
		return progress.Int32, nil
	}

	return 0, nil
}

func (s *courseService) GetCourseCount(ctx context.Context) (int64, error) {
	count, err := s.queries.CourseCount(ctx)
	if err != nil {
		return 0, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringCourseCount)
	}

	return count, nil
}
