package domains

import (
	"context"
	"time"

	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type IPLanguagesRepository interface {
	Filter(ctx context.Context, filter ProgrammingLanguageFilter, limit, page int64) (pLanguages []ProgrammingLanguage, dataCount int64, err error)
}

type IPLanguagesService interface {
	GetProgrammingLanguages(ctx context.Context, programmingLanguageID, languageID, name, page, limit string) (programmingLanguages []ProgrammingLanguage, err error)
}

const (
	DefaultProgrammingLanguageLimit = 10
)

type ProgrammingLanguage struct {
	id            uuid.UUID
	languageID    uuid.UUID
	name          string
	description   string
	downloadCMD   string
	compileCMD    string
	imagePath     string
	fileExtention string
	monacoEditor  string
	createdAt     time.Time
}

type ProgrammingLanguageFilter struct {
	ID         uuid.UUID
	LanguageID uuid.UUID
	Name       string
}

func NewProgrammingLanguage(
	id, languageID uuid.UUID,
	name, description, downloadCMD, compileCMD, imagePath, fileExtention, monacoEditor string,
	createdAt time.Time,
) (*ProgrammingLanguage, error) {
	var pLanguage ProgrammingLanguage

	if err := pLanguage.SetCompileCMD(compileCMD); err != nil {
		return nil, err
	}
	if err := pLanguage.SetDownloadCMD(downloadCMD); err != nil {
		return nil, err
	}
	if err := pLanguage.SetFileExtention(fileExtention); err != nil {
		return nil, err
	}
	if err := pLanguage.SetMonacoEditor(monacoEditor); err != nil {
		return nil, err
	}
	if err := pLanguage.SetImagePath(imagePath); err != nil {
		return nil, err
	}
	if err := pLanguage.SetName(name); err != nil {
		return nil, err
	}

	pLanguage.SetDescription(description)

	return &pLanguage, nil
}

func (d *ProgrammingLanguage) Unmarshal(
	id, languageID uuid.UUID,
	name, description, downloadCMD, compileCMD, imagePath, fileExtention, monacoEditor string,
	createdAt time.Time,
) {
	d.id = id
	d.languageID = languageID
	d.name = name
	d.description = description
	d.downloadCMD = downloadCMD
	d.compileCMD = compileCMD
	d.imagePath = imagePath
	d.fileExtention = fileExtention
	d.monacoEditor = monacoEditor
	d.createdAt = createdAt
}

func (d *ProgrammingLanguage) GetID() uuid.UUID {
	return d.id
}

func (d *ProgrammingLanguage) GetLanguageID() uuid.UUID {
	return d.languageID
}

func (d *ProgrammingLanguage) GetName() string {
	return d.name
}

func (d *ProgrammingLanguage) SetName(name string) error {
	if name == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrPLanguageNameCannotBeEmpty)
	}
	if len(name) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrPLanguageNameTooLong)
	}
	d.name = name
	return nil
}

func (d *ProgrammingLanguage) GetDescription() string {
	return d.description
}

func (d *ProgrammingLanguage) SetDescription(description string) {
	d.description = description
}

func (d *ProgrammingLanguage) GetDownloadCMD() string {
	return d.downloadCMD
}

func (d *ProgrammingLanguage) SetDownloadCMD(downloadCMD string) error {
	if downloadCMD == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrPLanguageDownloadCMDCannotBeEmpty)
	}
	if len(downloadCMD) > 256 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrPLanguageDownloadCMDTooLong)
	}
	d.name = downloadCMD
	return nil
}

func (d *ProgrammingLanguage) GetCompileCMD() string {
	return d.compileCMD
}

func (d *ProgrammingLanguage) SetCompileCMD(compileCMD string) error {
	if compileCMD == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrPLanguageCompileCMDCannotBeEmpty)
	}
	if len(compileCMD) > 256 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrPLanguageCompileCMDTooLong)
	}
	d.name = compileCMD
	return nil
}

func (d *ProgrammingLanguage) GetImagePath() string {
	return d.imagePath
}

func (d *ProgrammingLanguage) SetImagePath(imagePath string) error {
	if imagePath == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrPLanguageImagePathCannotBeEmpty)
	}
	if len(imagePath) > 60 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrPLanguageImagePathTooLong)
	}
	d.name = imagePath
	return nil
}

func (d *ProgrammingLanguage) GetFileExtention() string {
	return d.fileExtention
}

func (d *ProgrammingLanguage) SetFileExtention(fileExtention string) error {
	if fileExtention == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrPLanguageFileExtentionCannotBeEmpty)
	}
	if len(fileExtention) > 10 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrPLanguageFileExtentionTooLong)
	}
	d.name = fileExtention
	return nil
}

func (d *ProgrammingLanguage) GetMonacoEditor() string {
	return d.monacoEditor
}

func (d *ProgrammingLanguage) SetMonacoEditor(monacoEditor string) error {
	if monacoEditor == "" {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrPLanguageFileExtentionCannotBeEmpty)
	}
	if len(monacoEditor) > 10 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrPLanguageFileExtentionTooLong)
	}
	d.name = monacoEditor
	return nil
}

func (d *ProgrammingLanguage) GetCreatedAt() time.Time {
	return d.createdAt
}
