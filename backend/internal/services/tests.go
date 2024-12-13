package services

import (
	"context"
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
	if chapterID != "" {
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

	for i := range tests {
		inputs, _, err := s.testRepository.FilterInput(ctx, domains.GeneralFilter{
			ID:     uuid.Nil,
			TestID: tests[i].GetID(),
		}, int64(limitNum), int64(pageNum))
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringInputs, err)
		}

		outputs, _, err := s.testRepository.FilterOutput(ctx, domains.GeneralFilter{
			ID:     uuid.Nil,
			TestID: tests[i].GetID(),
		}, int64(limitNum), int64(pageNum))
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringOutputs, err)
		}
		tests[i].SetInputs(inputs)
		tests[i].SetOutputs(outputs)
	}

	return tests, nil
}

func (s *testService) GetInputs(
	ctx context.Context,
	id, testID, page, limit string,
) (inputs []domains.Input, err error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = domains.DefaultTestLimit
	}

	var idUUID uuid.UUID
	var testUUID uuid.UUID
	if id != "" {
		idUUID, err = uuid.Parse(id)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}
	if testID != "" {
		testUUID, err = uuid.Parse(testID)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}

	inputs, _, err = s.testRepository.FilterInput(ctx, domains.GeneralFilter{
		ID:     idUUID,
		TestID: testUUID,
	}, int64(limitNum), int64(pageNum))
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringInputs, err)
	}

	return
}

func (s *testService) AddInput(
	ctx context.Context,
	testID, value string,
) (uuid.UUID, error) {
	newInput, err := domains.NewInput(
		"",
		testID,
		value,
	)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := s.testRepository.AddInput(ctx, newInput)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
