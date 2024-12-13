package dto

type IDTOManager interface {
	UserManager() *UserDTOManager
	AdminManager() *AdminDTOManager
	LanguageManager() *LanguageDTOManager
	RewardManager() *RewardDTOManager
	ProgrammingManager() *ProgrammingLanguageDTOManager
	CourseManager() *CourseDTOManager
	ChapterManager() *ChapterDTOManager
	TestManager() *TestDTOManager
}

type DTOManager struct {
	userDTOManager      *UserDTOManager
	adminDTOManager     *AdminDTOManager
	languageDTOManager  *LanguageDTOManager
	rewardDTOManager    *RewardDTOManager
	pLanguageDTOManager *ProgrammingLanguageDTOManager
	courseDTOManager    *CourseDTOManager
	chapterDTOManager   *ChapterDTOManager
	testDTOManager      *TestDTOManager
}

func CreateNewDTOManager() IDTOManager {
	userDTOManager := NewUserDTOManager()
	adminDTOManager := NewAdminDTOManager()
	languageDTOManager := NewLanguageDTOManager()
	rewardDTOManager := NewRewardDTOManager()
	pLanguageDTOManager := NewProgrammingLanguageDTOManager()
	courseDTOManager := NewCourseDTOManager()
	chapterDTOManager := NewChapterDTOManager()
	testDTOManager := NewTestDTOManager()

	return &DTOManager{
		userDTOManager:      &userDTOManager,
		adminDTOManager:     &adminDTOManager,
		languageDTOManager:  &languageDTOManager,
		rewardDTOManager:    &rewardDTOManager,
		pLanguageDTOManager: &pLanguageDTOManager,
		courseDTOManager:    &courseDTOManager,
		chapterDTOManager:   &chapterDTOManager,
		testDTOManager:      &testDTOManager,
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

func (m *DTOManager) ChapterManager() *ChapterDTOManager {
	return m.chapterDTOManager
}

func (m DTOManager) TestManager() *TestDTOManager {
	return m.testDTOManager
}
