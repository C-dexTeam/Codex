package services

import (
	"context"
	"database/sql"
	"strings"

	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	repo "github.com/C-dexTeam/codex/internal/repos/out"
)

type questService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func NewQuestService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *questService {
	return &questService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *questService) GetQuest(ctx context.Context, chapterID, courseID string) (*repo.TChapter, []repo.TTest, *repo.TProgrammingLanguage, error) {
	chapterUUID, err := s.utilService.NParseUUID(chapterID)
	if err != nil {
		return nil, nil, nil, err
	}
	courseUUID, err := s.utilService.NParseUUID(courseID)
	if err != nil {
		return nil, nil, nil, err
	}

	chapter, err := s.queries.GetChapterByID(ctx, chapterUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, nil, nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrChapterNotFound)
		}
		return nil, nil, nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringChapter, err)
	}

	// TODO: Return tests with input and output by chapter id
	chapterTests, err := s.queries.GetTests(ctx, repo.GetTestsParams{
		ChapterID: s.utilService.ParseNullUUID(chapterID),
		Lim:       100,
		Off:       0,
	})
	if err != nil {
		return nil, nil, nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringTests,
			err,
		)
	}

	course, err := s.queries.GetCourseByID(ctx, courseUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, nil, nil, serviceErrors.NewServiceErrorWithMessage(
				serviceErrors.StatusBadRequest,
				serviceErrors.ErrCourseNotFound,
			)
		}
		return nil, nil, nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringCourse, err)
	}

	pLanguage, err := s.queries.GetPLanguageByID(ctx, course.ProgrammingLanguageID.UUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, nil, nil, serviceErrors.NewServiceErrorWithMessage(
				serviceErrors.StatusBadRequest,
				serviceErrors.ErrProgrammingLanguageNotFound,
			)
		}
		return nil, nil, nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringProgrammingLanguages, err)
	}

	return &chapter, chapterTests, &pLanguage, nil
}
