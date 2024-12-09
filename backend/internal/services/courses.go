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
	chapterRepsitory domains.IChapterRepository
}

func newCourseService(
	courseRepository domains.ICourseRepository,
	chapterRepsitory domains.IChapterRepository,
) domains.ICourseService {
	return &courseService{
		courseRepository: courseRepository,
		chapterRepsitory: chapterRepsitory,
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

func (s *courseService) GetCourse(
	ctx context.Context,
	id, page, limit string,
) (course *domains.Course, err error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = domains.DefaultChapterLimit
	}

	idUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}
	courses, _, err := s.courseRepository.Filter(ctx, domains.CourseFilter{
		ID: idUUID,
	}, 1, 1)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringRewards, err)
	}
	if len(courses) != 1 {
		return nil, serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusNotFound, errorDomains.ErrRewardNotFound)
	}
	course = &courses[0]

	courseChapters, _, err := s.chapterRepsitory.Filter(ctx, domains.ChapterFilter{
		CourseID: course.GetID(),
	}, int64(limitNum), int64(pageNum))
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringChapter, err)
	}
	course.SetChapters(courseChapters)

	return course, nil
}

func (s *courseService) AddCourse(
	ctx context.Context,
	languageID, pLanguageID, rewardID, title, description, imagePath string,
	rewardAmount int,
) (uuid.UUID, error) {
	newCourse, err := domains.NewCourse(
		"",
		languageID,
		pLanguageID,
		rewardID,
		rewardAmount,
		title,
		description,
		imagePath,
		nil,
	)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := s.courseRepository.Add(ctx, newCourse)
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
	var idUUID uuid.UUID
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	courses, _, err := s.courseRepository.Filter(ctx, domains.CourseFilter{
		ID: idUUID,
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringCourse, err)
	}
	if len(courses) != 1 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusNotFound, errorDomains.ErrCourseNotFound)
	}

	updateCourse, err := domains.NewCourse(
		id,
		languageID,
		pLanguageID,
		rewardID,
		rewardAmount,
		title,
		description,
		imagePath,
		nil,
	)
	if err != nil {
		return err
	}

	if err := s.courseRepository.Update(ctx, updateCourse); err != nil {
		return err
	}

	return nil
}

func (s *courseService) DeleteCourse(
	ctx context.Context,
	id string,
) (err error) {
	var idUUID uuid.UUID
	idUUID, err = uuid.Parse(id)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	courses, _, err := s.courseRepository.Filter(ctx, domains.CourseFilter{
		ID: idUUID,
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringCourse, err)
	}
	if len(courses) != 1 {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrCourseNotFound, err)
	}

	if err = s.courseRepository.SoftDelete(ctx, idUUID); err != nil {
		return
	}
	return
}
