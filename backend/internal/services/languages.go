package services

import (
	"context"
	"database/sql"
	"strings"

	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	repo "github.com/C-dexTeam/codex/internal/repos/out"
)

type languageService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func newLanguageService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *languageService {
	return &languageService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *languageService) GetLanguages(
	ctx context.Context,
	id, value string,
) ([]repo.TLanguage, error) {
	if _, err := s.utilService.ParseUUID(id); err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrInvalidID)
	}

	languages, err := s.queries.GetLanguages(ctx, repo.GetLanguagesParams{
		ID:    s.utilService.ParseNullUUID(id),
		Value: s.utilService.ParseString(value),
		Lim:   int32(s.utilService.D().Limits.DefaultLanguageLimit),
		Off:   0,
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringLanguages, err)
	}

	return languages, nil
}

func (s *languageService) GetLanguage(
	ctx context.Context,
	id string,
) (*repo.TLanguage, error) {
	languageID, err := s.utilService.ParseUUID(id)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrInvalidID)
	}

	language, err := s.queries.GetLanguageByID(ctx, languageID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrLanguageNotFound)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringLanguages, err)
	}
	return &language, nil
}

func (s *languageService) GetByValue(
	ctx context.Context,
	value string,
) (*repo.TLanguage, error) {
	language, err := s.queries.GetLanguageByValue(ctx, value)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrLanguageNotFound)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringLanguages, err)
	}

	return &language, nil
}

func (s *languageService) GetDefault(
	ctx context.Context,
) (*repo.TLanguage, error) {
	defaultLanguage, err := s.queries.GetDefaultLanguage(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrLanguageNotFound)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringLanguages, err)
	}

	return &defaultLanguage, nil
}
