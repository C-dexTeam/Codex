package dto

type IDTOManager interface {
	UserManager() *UserDTOManager
	LanguageManager() *LanguageDTOManager
	RewardManager() *RewardDTOManager
	ProgrammingManager() *ProgrammingLanguageDTOManager
	CourseManager() *CourseDTOManager
	ChapterManager() *ChapterDTOManager
	TestManager() *TestDTOManager
	QuestManager() *QuestDTOManager
}

type DTOManager struct {
	userDTOManager      *UserDTOManager
	languageDTOManager  *LanguageDTOManager
	rewardDTOManager    *RewardDTOManager
	pLanguageDTOManager *ProgrammingLanguageDTOManager
	courseDTOManager    *CourseDTOManager
	chapterDTOManager   *ChapterDTOManager
	testDTOManager      *TestDTOManager
	questDTOManager     *QuestDTOManager
}

func CreateNewDTOManager() IDTOManager {
	userDTOManager := NewUserDTOManager()
	languageDTOManager := NewLanguageDTOManager()
	rewardDTOManager := NewRewardDTOManager()
	pLanguageDTOManager := NewProgrammingLanguageDTOManager()
	courseDTOManager := NewCourseDTOManager()
	chapterDTOManager := NewChapterDTOManager()
	testDTOManager := NewTestDTOManager()
	questDTOManager := NewQuestDTOManager()

	return &DTOManager{
		userDTOManager:      &userDTOManager,
		languageDTOManager:  &languageDTOManager,
		rewardDTOManager:    &rewardDTOManager,
		pLanguageDTOManager: &pLanguageDTOManager,
		courseDTOManager:    &courseDTOManager,
		chapterDTOManager:   &chapterDTOManager,
		testDTOManager:      &testDTOManager,
		questDTOManager:     &questDTOManager,
	}
}

func (m *DTOManager) QuestManager() *QuestDTOManager {
	return m.questDTOManager
}

func (m *DTOManager) UserManager() *UserDTOManager {
	return m.userDTOManager
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
