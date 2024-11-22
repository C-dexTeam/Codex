package dto

type IDTOManager interface {
	UserManager() *UserDTOManager
	AdminManager() *AdminDTOManager
}

type DTOManager struct {
	userDTOManager  *UserDTOManager
	adminDTOManager *AdminDTOManager
}

func CreateNewDTOManager() IDTOManager {
	userDTOManager := NewUserDTOManager()
	adminDTOManager := NewAdminDTOManager()

	return &DTOManager{
		userDTOManager:  &userDTOManager,
		adminDTOManager: &adminDTOManager,
	}
}

func (m *DTOManager) UserManager() *UserDTOManager {
	return m.userDTOManager
}

func (m *DTOManager) AdminManager() *AdminDTOManager {
	return m.adminDTOManager
}
