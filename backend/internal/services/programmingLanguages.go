package services

import "github.com/C-dexTeam/codex/internal/domains"

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
