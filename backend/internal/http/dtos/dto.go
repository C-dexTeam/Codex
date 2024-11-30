package dto

type IDTOManager interface {
	UserManager() *UserDTOManager
	AdminManager() *AdminDTOManager
	LanguageManager() *LanguageDTOManager
	RewardManager() *RewardDTOManager
}

type DTOManager struct {
	userDTOManager     *UserDTOManager
	adminDTOManager    *AdminDTOManager
	languageDTOManager *LanguageDTOManager
	rewardDTOManager   *RewardDTOManager
}

func CreateNewDTOManager() IDTOManager {
	userDTOManager := NewUserDTOManager()
	adminDTOManager := NewAdminDTOManager()
	languageDTOManager := NewLanguageDTOManager()
	rewardDTOManager := NewRewardDTOManager()

	return &DTOManager{
		userDTOManager:     &userDTOManager,
		adminDTOManager:    &adminDTOManager,
		languageDTOManager: &languageDTOManager,
		rewardDTOManager:   &rewardDTOManager,
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

func (m *DTOManager) RewardManager() *RewardDTOManager {
	return m.rewardDTOManager
}
