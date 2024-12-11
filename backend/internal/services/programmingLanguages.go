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

func (s *pLanguageService) GetProgrammingLanguage(
	ctx context.Context,
	id string,
) (programmingLanguage *domains.ProgrammingLanguage, err error) {
	pLanguageUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	programmingLanguages, _, err := s.pLanguageRepository.Filter(ctx, domains.ProgrammingLanguageFilter{
		ID: pLanguageUUID,
	}, 1, 1)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringProgrammingLanguages, err)
	}
	if len(programmingLanguages) != 1 {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrProgrammingLanguageNotFound, err)
	}
	programmingLanguage = &programmingLanguages[0]

	return programmingLanguage, nil
}

func (s *pLanguageService) AddProgrammingLanguage(
	ctx context.Context,
	languageID, name, description, downloadCMD, compileCMD, imagePath, fileExtention, monacoEditor string,
) (uuid.UUID, error) {
	newProgrammingLanguage, err := domains.NewProgrammingLanguage(
		"",
		languageID,
		name,
		description,
		downloadCMD,
		compileCMD,
		imagePath,
		fileExtention,
		monacoEditor,
	)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := s.pLanguageRepository.Add(ctx, newProgrammingLanguage)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s *pLanguageService) UpdateProgrammingLanguage(
	ctx context.Context,
	id, languageID, name, description, downloadCMD, compileCMD, imagePath, fileExtention, monacoEditor string,
) error {
	var (
		idUUID       uuid.UUID
		languageUUID uuid.UUID
	)
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}
	pLanguages, _, err := s.pLanguageRepository.Filter(ctx, domains.ProgrammingLanguageFilter{
		ID: idUUID,
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringProgrammingLanguages, err)
	}
	if len(pLanguages) != 1 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusNotFound, errorDomains.ErrProgrammingLanguageNotFound)
	}

	if languageID != "" {
		languageUUID, err = uuid.Parse(languageID)
		if err != nil {
			return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
		languages, _, err := s.pLanguageRepository.Filter(ctx, domains.ProgrammingLanguageFilter{
			ID: languageUUID,
		}, 1, 1)
		if err != nil {
			return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringProgrammingLanguages, err)
		}
		if len(languages) != 1 {
			return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusNotFound, errorDomains.ErrLanguageNotFound)
		}
	}

	updateProgrammingLanguage, err := domains.NewProgrammingLanguage(
		id,
		languageID,
		name,
		description,
		downloadCMD,
		compileCMD,
		imagePath,
		fileExtention,
		monacoEditor,
	)
	if err != nil {
		return err
	}

	if err := s.pLanguageRepository.Update(ctx, updateProgrammingLanguage); err != nil {
		return err
	}

	return nil
}

func (s *pLanguageService) DeleteProgrammingLanguage(
	ctx context.Context,
	id string,
) (err error) {
	var idUUID uuid.UUID
	idUUID, err = uuid.Parse(id)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	pLanguages, _, err := s.pLanguageRepository.Filter(ctx, domains.ProgrammingLanguageFilter{
		ID: idUUID,
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringRewards, err)
	}
	if len(pLanguages) != 1 {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrProgrammingLanguageNotFound, err)
	}

	if err := s.pLanguageRepository.Delete(ctx, idUUID); err != nil {
		return err
	}

	return nil
}
