package dto

import (
	repo "github.com/C-dexTeam/codex/internal/repos/out"
)

type TestDTOManager struct{}

func NewTestDTOManager() TestDTOManager {
	return TestDTOManager{}
}

type TestView struct {
	ID        string `json:"id"`
	ChapterID string `json:"chapterID"`
	Input     string `json:"input"`
	Output    string `json:"output"`
}

func (t *TestDTOManager) ToTestDTO(appModel repo.TTest) TestView {
	return TestView{
		ID:        appModel.ID.String(),
		ChapterID: appModel.ChapterID.String(),
		Input:     appModel.InputValue,
		Output:    appModel.OutputValue,
	}
}

func (t *TestDTOManager) ToTestDTOs(appModels []repo.TTest) []TestView {
	var testDTOs []TestView
	for _, model := range appModels {
		testDTOs = append(testDTOs, t.ToTestDTO(model))
	}

	return testDTOs
}

type AddTestDTO struct {
	ChapterID   string `json:"chapterID" validate:"required,uuid4"`
	InputValue  string `json:"inputValue"`
	OutputValue string `json:"outputValue"`
}

type UpdateTestDTO struct {
	ID          string `json:"id" validate:"required,uuid4"`
	InputValue  string `json:"inputValue"`
	OutputValue string `json:"outputValue"`
}
