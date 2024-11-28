package services

import (
	"context"

	"github.com/C-dexTeam/codex/internal/domains"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type languageService struct {
	languageRepository domains.ILanguagesRepository
}

func newLanguageService(
	languageRepository domains.ILanguagesRepository,

) domains.ILanguagesService {
	return &languageService{
		languageRepository: languageRepository,
	}
}

func (s *languageService) GetLanguages(ctx context.Context, languageID, value string) (languages []domains.Languages, err error) {
	languageUUID, err := uuid.Parse(languageID)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(500, domains.ErrLanguageNotFound, err)
	}

	langauges, _, err := s.languageRepository.Filter(ctx, domains.LanguagesFilter{
		ID:    languageUUID,
		Value: value,
	}, 1, 1)
	if len(langauges) != 1 {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(domains.StatusNotFound, domains.ErrLanguageNotFound, err)
	}
	return
}

func (s *languageService) GetDefault(ctx context.Context) (language *domains.Languages, err error) {
	langauges, _, err := s.languageRepository.Filter(ctx, domains.LanguagesFilter{
		Value: domains.DefaultLanguage,
	}, 1, 1)
	if len(langauges) != 1 {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(domains.StatusNotFound, domains.ErrLanguageNotFound, err)
	}
	language = &langauges[0]
	return
}
