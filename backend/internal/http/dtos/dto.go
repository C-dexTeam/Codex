package dto

type IDTOManager interface {
	UserManager() *UserDTOManager
	AdminManager() *AdminDTOManager
	LanguageManager() *LanguageDTOManager
}

type DTOManager struct {
	userDTOManager     *UserDTOManager
	adminDTOManager    *AdminDTOManager
	languageDTOManager *LanguageDTOManager
}

func CreateNewDTOManager() IDTOManager {
	userDTOManager := NewUserDTOManager()
	adminDTOManager := NewAdminDTOManager()
	languageDTOManager := NewLanguageDTOManager()

	return &DTOManager{
		userDTOManager:     &userDTOManager,
		adminDTOManager:    &adminDTOManager,
		languageDTOManager: &languageDTOManager,
	}
}

func (m *DTOManager) UserManager() *UserDTOManager {
	return m.userDTOManager
}

func (m *DTOManager) AdminManager() *AdminDTOManager {
	return m.adminDTOManager
}

func (m *DTOManager) LanguageManager() *LanguageDTOManager {
	return m.languageDTOManager
}
