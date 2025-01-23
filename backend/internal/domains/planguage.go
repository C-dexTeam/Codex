package domains

import (
	"time"

	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type PLanguage struct {
	ID            uuid.UUID
	LanguageID    uuid.UUID
	Name          string
	Description   string
	ImagePath     string
	FileExtention string
	MonacoEditor  string
	CreatedAt     time.Time
}

func NewPLanguage(
	pLanguage *repo.TProgrammingLanguage,
) *PLanguage {
	if pLanguage == nil {
		return nil
	}

	return &PLanguage{
		ID:            pLanguage.ID,
		LanguageID:    pLanguage.LanguageID,
		Name:          pLanguage.Name,
		Description:   pLanguage.Description,
		ImagePath:     pLanguage.ImagePath,
		FileExtention: pLanguage.FileExtention,
		MonacoEditor:  pLanguage.MonacoEditor,
		CreatedAt:     pLanguage.CreatedAt,
	}
}

func NewPLanguages(pLanguages []repo.TProgrammingLanguage) []PLanguage {
	var domainPLanguages []PLanguage
	for _, plang := range pLanguages {
		domainPLanguages = append(domainPLanguages, *NewPLanguage(&plang))
	}

	return domainPLanguages
}
