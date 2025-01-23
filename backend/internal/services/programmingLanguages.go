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

type pLanguageService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func newPLanguageService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *pLanguageService {
	return &pLanguageService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *pLanguageService) GetProgrammingLanguages(ctx context.Context,
	id, languageID, name, page, limit string,
) ([]domains.PLanguage, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = s.utilService.D().Limits.DefaultProgrammingLanguageLimit
	}

	programmingLanguages, err := s.queries.GetPLanguages(ctx, repo.GetPLanguagesParams{
		ID:         s.utilService.ParseNullUUID(id),
		LanguageID: s.utilService.ParseNullUUID(languageID),
		Name:       s.utilService.ParseString(name),
		Lim:        int32(limitNum),
		Off:        (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringProgrammingLanguages,
			err,
		)
	}
	domainPLang := domains.NewPLanguages(programmingLanguages)

	return domainPLang, nil
}

func (s *pLanguageService) GetProgrammingLanguage(
	ctx context.Context,
	id string,
) (*domains.PLanguage, error) {

	idUUID, err := s.utilService.ParseUUID(id)
	if err != nil {
		return nil, err
	}

	programmingLanguage, err := s.queries.GetPLanguageByID(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(
				serviceErrors.StatusBadRequest,
				serviceErrors.ErrProgrammingLanguageNotFound,
			)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringProgrammingLanguages,
			err,
		)
	}
	domainPLang := domains.NewPLanguage(&programmingLanguage)

	return domainPLang, nil
}

func (s *pLanguageService) AddProgrammingLanguage(
	ctx context.Context,
	languageID, name, description, imagePath, fileExtention, monacoEditor string,
) (uuid.UUID, error) {
	languageUUID, err := s.utilService.ParseUUID(languageID)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := s.queries.CreatePLanguage(ctx, repo.CreatePLanguageParams{
		LanguageID:    languageUUID,
		Name:          name,
		Description:   description,
		ImagePath:     imagePath,
		FileExtention: fileExtention,
		MonacoEditor:  monacoEditor,
	})
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s *pLanguageService) UpdateProgrammingLanguage(
	ctx context.Context,
	id, languageID, name, description, imagePath, fileExtention, monacoEditor string,
) error {
	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}

	languageUUID, err := s.utilService.ParseUUID(languageID)
	if err != nil {
		return err
	}

	if ok, err := s.queries.CheckPLanguageByID(ctx, idUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringProgrammingLanguages)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrProgrammingLanguageNotFound)
	}

	if languageID != "" {
		if ok, err := s.queries.CheckLanguageByID(ctx, languageUUID); err != nil {
			return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringLanguages)
		} else if !ok {
			return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrLanguageNotFound)
		}
	}

	if err := s.queries.UpdatePLanguage(ctx, repo.UpdatePLanguageParams{
		ProgrammingLanguageID: idUUID,
		LanguageID:            s.utilService.ParseNullUUID(languageID),
		Name:                  s.utilService.ParseString(name),
		Description:           s.utilService.ParseString(description),
		FileExtention:         s.utilService.ParseString(fileExtention),
		MonacoEditor:          s.utilService.ParseString(monacoEditor),
		ImagePath:             s.utilService.ParseString(imagePath),
	}); err != nil {
		return err
	}

	return nil
}

func (s *pLanguageService) DeleteProgrammingLanguage(
	ctx context.Context,
	id string,
) (err error) {
	idUUID, err := s.utilService.ParseUUID(id)
	if err != nil {
		return err
	}

	if ok, err := s.queries.CheckPLanguageByID(ctx, idUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringProgrammingLanguages)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrProgrammingLanguageNotFound)
	}

	if err := s.queries.DeletePLanguage(ctx, idUUID); err != nil {
		return err
	}

	return nil
}
