package dto

type IDTOManager interface {
	UserManager() *UserDTOManager
	AdminManager() *AdminDTOManager
	LanguageManager() *LanguageDTOManager
	RewardManager() *RewardDTOManager
	ProgrammingManager() *ProgrammingLanguageDTOManager
	CourseManager() *CourseDTOManager
}

type DTOManager struct {
	userDTOManager      *UserDTOManager
	adminDTOManager     *AdminDTOManager
	languageDTOManager  *LanguageDTOManager
	rewardDTOManager    *RewardDTOManager
	pLanguageDTOManager *ProgrammingLanguageDTOManager
	courseDTOManager    *CourseDTOManager
}

func CreateNewDTOManager() IDTOManager {
	userDTOManager := NewUserDTOManager()
	adminDTOManager := NewAdminDTOManager()
	languageDTOManager := NewLanguageDTOManager()
	rewardDTOManager := NewRewardDTOManager()
	pLanguageDTOManager := NewProgrammingLanguageDTOManager()
	courseDTOManager := NewCourseDTOManager()

	return &DTOManager{
		userDTOManager:      &userDTOManager,
		adminDTOManager:     &adminDTOManager,
		languageDTOManager:  &languageDTOManager,
		rewardDTOManager:    &rewardDTOManager,
		pLanguageDTOManager: &pLanguageDTOManager,
		courseDTOManager:    &courseDTOManager,
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

func (m *DTOManager) ProgrammingManager() *ProgrammingLanguageDTOManager {
	return m.pLanguageDTOManager
}

func (m *DTOManager) CourseManager() *CourseDTOManager {
	return m.courseDTOManager
}
