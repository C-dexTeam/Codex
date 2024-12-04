package services

import (
	"context"
	"strconv"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type courseService struct {
	courseRepository domains.ICourseRepository
}

func newCourseService(
	courseRepository domains.ICourseRepository,
) domains.ICourseService {
	return &courseService{
		courseRepository: courseRepository,
	}
}

func (s *courseService) GetCourses(
	ctx context.Context,
	id, langugeID, pLanguageID, title, page, limit string,
) (courses []domains.Course, err error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = domains.DefaultCourseLimit
	}

	var (
		courseUUID    uuid.UUID
		languageUUID  uuid.UUID
		pLanguageUUID uuid.UUID
	)
	if id != "" {
		courseUUID, err = uuid.Parse(id)
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
	if pLanguageID != "" {
		pLanguageUUID, err = uuid.Parse(pLanguageID)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}

	courses, _, err = s.courseRepository.Filter(ctx, domains.CourseFilter{
		ID:          courseUUID,
		LanguageID:  languageUUID,
		PLanguageID: pLanguageUUID,
		Title:       title,
	}, int64(limitNum), int64(pageNum))
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringCourse, err)
	}

	return courses, nil
}
