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

type testService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func newTestService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *testService {
	return &testService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *testService) GetTests(
	ctx context.Context,
	id, chapterID, page, limit string,
) ([]domains.Test, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = s.utilService.D().Limits.DefaultTestLimit
	}

	if _, err := s.utilService.ParseUUID(id); err != nil {
		return nil, err
	}
	if _, err := s.utilService.ParseUUID(chapterID); err != nil {
		return nil, err
	}

	tests, err := s.queries.GetTests(ctx, repo.GetTestsParams{
		ID:        s.utilService.ParseNullUUID(id),
		ChapterID: s.utilService.ParseNullUUID(chapterID),
		Lim:       int32(limitNum),
		Off:       (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringRewards,
			err,
		)
	}
	domainTests := domains.NewTests(tests)

	return domainTests, nil
}

func (s *testService) GetTestByID(
	ctx context.Context,
	id string,
) (*domains.Test, error) {

	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return nil, err
	}

	test, err := s.queries.GetTest(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrTestNotFound)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringTests, err)
	}
	domainTest := domains.NewTest(&test)

	return domainTest, nil
}

func (s *testService) AddTest(
	ctx context.Context,
	chapterID, inputValue, outputValue string,
) (uuid.UUID, error) {
	chapterUUID, err := s.utilService.NParseUUID(chapterID)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := s.queries.CreateTest(ctx, repo.CreateTestParams{
		ChapterID:   chapterUUID,
		InputValue:  inputValue,
		OutputValue: outputValue,
	})
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s *testService) UpdateTest(
	ctx context.Context,
	id, inputValue, outputValue string,
) error {
	idUUID, err := s.utilService.ParseUUID(id)
	if err != nil {
		return err
	}
	if ok, err := s.queries.CheckTestByID(ctx, idUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringTests, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrTestNotFound)
	}

	if err := s.queries.UpdateTest(ctx, repo.UpdateTestParams{
		TestID:      idUUID,
		InputValue:  s.utilService.ParseString(inputValue),
		OutputValue: s.utilService.ParseString(outputValue),
	}); err != nil {
		return err
	}

	return nil
}

func (s *testService) DeleteTest(
	ctx context.Context,
	id string,
) (err error) {
	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}

	_, err = s.queries.GetTest(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return serviceErrors.NewServiceErrorWithMessage(
				serviceErrors.StatusBadRequest,
				serviceErrors.ErrTestNotFound,
			)
		}
		return serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringTests,
			err,
		)
	}

	if err := s.queries.DeleteTest(ctx, idUUID); err != nil {
		return err
	}
	return
}
