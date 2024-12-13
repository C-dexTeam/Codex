package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type testService struct {
	testRepository domains.ITestRepository
}

func newTestService(
	testRepository domains.ITestRepository,
) domains.ITestService {
	return &testService{
		testRepository: testRepository,
	}
}

func (s *testService) GetTests(
	ctx context.Context,
	id, chapterID, page, limit string,
) (tests []domains.Test, err error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = domains.DefaultTestLimit
	}

	var idUUID uuid.UUID
	var chapterUUID uuid.UUID
	if id != "" {
		idUUID, err = uuid.Parse(id)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}
	if id != "" {
		chapterUUID, err = uuid.Parse(chapterID)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}

	tests, _, err = s.testRepository.FilterTest(ctx, domains.TestFilter{
		ID:        idUUID,
		ChapterID: chapterUUID,
	}, int64(limitNum), int64(pageNum))
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringTests, err)
	}

	for _, test := range tests {
		inputs, _, err := s.testRepository.FilterInput(ctx, domains.GeneralFilter{
			ID:     uuid.Nil,
			TestID: test.GetID(),
		}, int64(limitNum), int64(pageNum))
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringInputs, err)
		}

		fmt.Println("inputs:", inputs)

		outputs, _, err := s.testRepository.FilterOutput(ctx, domains.GeneralFilter{
			ID:     uuid.Nil,
			TestID: test.GetID(),
		}, int64(limitNum), int64(pageNum))
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringOutputs, err)
		}
		fmt.Println("outputs:", outputs)

		test.SetInputs(inputs)
		test.SetOutputs(outputs)
	}

	return
}
