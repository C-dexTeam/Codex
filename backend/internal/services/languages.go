package services

import (
	"context"
	"database/sql"
	"strings"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
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
		return nil, serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
	}

	languages, err := s.queries.GetLanguages(ctx, repo.GetLanguagesParams{
		ID:    s.utilService.ParseNullUUID(id),
		Value: s.utilService.ParseString(value),
		Lim:   domains.DefaultLanguageLimit,
		Off:   0,
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringLanguages, err)
	}

	return languages, nil
}

func (s *languageService) GetLanguage(
	ctx context.Context,
	id string,
) (*repo.TLanguage, error) {
	languageID, err := s.utilService.ParseUUID(id)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
	}

	language, err := s.queries.GetLanguageByID(ctx, languageID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrUserNotFound)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
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
			return nil, serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrUserNotFound)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}

	return &language, nil
}

func (s *languageService) GetDefault(
	ctx context.Context,
) (*repo.TLanguage, error) {
	return s.GetByValue(ctx, domains.DefaultLanguage)
}
