package services

import (
	"context"
	"strconv"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type pLanguageService struct {
	pLanguageRepository domains.IPLanguagesRepository
}

func newPLanguageService(
	pLanguageRepository domains.IPLanguagesRepository,
) domains.IPLanguagesService {
	return &pLanguageService{
		pLanguageRepository: pLanguageRepository,
	}
}

func (s *pLanguageService) GetProgrammingLanguages(ctx context.Context,
	id, languageID, name, page, limit string,
) (programmingLanguages []domains.ProgrammingLanguage, err error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = domains.DefaultProgrammingLanguageLimit
	}

	var (
		pLanguageUUID uuid.UUID
		languageUUID  uuid.UUID
	)
	if id != "" {
		pLanguageUUID, err = uuid.Parse(id)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}
	if languageID != "" {
		languageUUID, err = uuid.Parse(languageID)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}

	programmingLanguages, _, err = s.pLanguageRepository.Filter(ctx, domains.ProgrammingLanguageFilter{
		ID:         pLanguageUUID,
		LanguageID: languageUUID,
		Name:       name,
	}, int64(limitNum), int64(pageNum))
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringProgrammingLanguages, err)
	}

	return programmingLanguages, nil
}
