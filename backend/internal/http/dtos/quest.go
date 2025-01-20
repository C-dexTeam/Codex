package dto

import (
	repo "github.com/C-dexTeam/codex/internal/repos/out"
)

type QuestDTOManager struct{}

func NewQuestDTOManager() QuestDTOManager {
	return QuestDTOManager{}
}

type QuestView struct {
	Chapter                QuestChapter         `json:"chapter"`
	Tests                  []QuestTest          `json:"tests"`
	ProgrammingLanguageDTO QuestProgrammingLang `json:"programmingLanguage"`
}

type QuestChapter struct {
	UserCode    string `json:"userCode"`
	FuncName    string `json:"funcname"`
	FrontendTmp string `json:"frontendTmp"`
	DockerTmp   string `json:"dockerTmp"`
	CheckTmp    string `json:"checkTmp"`
}

func (q *QuestDTOManager) QuestChapterDTO(chapter *repo.TChapter, userCode string) QuestChapter {
	return QuestChapter{
		UserCode:    userCode,
		FuncName:    chapter.FuncName,
		FrontendTmp: chapter.FrontendTemplate,
		DockerTmp:   chapter.DockerTemplate,
		CheckTmp:    chapter.CheckTemplate,
	}
}

type QuestProgrammingLang struct {
	Name          string `json:"name"`
	FileExtention string `json:"fileExtention"`
}

func (q *QuestDTOManager) QuestPLangDTO(pLang *repo.TProgrammingLanguage) QuestProgrammingLang {
	return QuestProgrammingLang{
		Name:          pLang.Name,
		FileExtention: pLang.FileExtention,
	}
}

type QuestTest struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

func (q *QuestDTOManager) QuestTestDTO(test *repo.TTest) QuestTest {
	return QuestTest{
		Input:  test.InputValue,
		Output: test.OutputValue,
	}
}

func (q *QuestDTOManager) QuestTestDTOs(tests []repo.TTest) []QuestTest {
	var testDTOs []QuestTest
	for _, model := range tests {
		testDTOs = append(testDTOs, q.QuestTestDTO(&model))
	}

	return testDTOs
}

func (q *QuestDTOManager) ToQuestDTO(chapter *repo.TChapter, tests []repo.TTest, pLanguage *repo.TProgrammingLanguage, userCode string) *QuestView {
	return &QuestView{
		Chapter:                q.QuestChapterDTO(chapter, userCode),
		Tests:                  q.QuestTestDTOs(tests),
		ProgrammingLanguageDTO: q.QuestPLangDTO(pLanguage),
	}
}
