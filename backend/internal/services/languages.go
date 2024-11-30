package services

import (
	"context"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
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

func (s *languageService) GetLanguages(ctx context.Context, languageID, value string) (languages []domains.Language, err error) {
	var languageUUID uuid.UUID
	if languageID != "" {
		languageUUID, err = uuid.Parse(languageID)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}

	languages, _, err = s.languageRepository.Filter(ctx, domains.LanguageFilter{
		ID:    languageUUID,
		Value: value,
	}, domains.DefaultLanguageLimit, 1)
	return
}

func (s *languageService) GetDefault(ctx context.Context) (language *domains.Language, err error) {
	langauges, _, err := s.languageRepository.Filter(ctx, domains.LanguageFilter{
		Value: domains.DefaultLanguage,
	}, 1, 1)
	if len(langauges) != 1 {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusNotFound, errorDomains.ErrLanguageDefaultNotFound, err)
	}
	language = &langauges[0]
	return
}
