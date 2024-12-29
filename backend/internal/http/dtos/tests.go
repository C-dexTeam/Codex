package dto

import (
	repo "github.com/C-dexTeam/codex/internal/repos/out"
)

type TestDTOManager struct{}

func NewTestDTOManager() TestDTOManager {
	return TestDTOManager{}
}

type TestDTO struct {
	ID      string      `json:"id"`
	Inputs  []InputDTO  `json:"inputs"`
	Outputs []OutputDTO `json:"outputs"`
}

type InputDTO struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type OutputDTO struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

func (t *TestDTOManager) ToInputDTO(appModel repo.TInput) InputDTO {
	return InputDTO{
		ID:    appModel.ID.String(),
		Value: appModel.Value.String,
	}
}

func (t *TestDTOManager) ToOutputDTO(appModel repo.TOutput) OutputDTO {
	return OutputDTO{
		ID:    appModel.ID.String(),
		Value: appModel.Value.String,
	}
}

func (t *TestDTOManager) ToInputDTOs(appModels []repo.TInput) []InputDTO {
	var inputDTOs []InputDTO
	for _, model := range appModels {
		inputDTOs = append(inputDTOs, t.ToInputDTO(model))
	}

	return inputDTOs
}

func (t *TestDTOManager) ToOutputDTOs(appModels []repo.TOutput) []OutputDTO {
	var outputDTOs []OutputDTO
	for _, model := range appModels {
		outputDTOs = append(outputDTOs, t.ToOutputDTO(model))
	}

	return outputDTOs
}

func (t *TestDTOManager) ToTestDTO(appModel repo.TTest, inputModels []repo.TInput, outputModels []repo.TOutput) TestDTO {
	return TestDTO{
		ID:      appModel.ID.String(),
		Inputs:  t.ToInputDTOs(inputModels),
		Outputs: t.ToOutputDTOs(outputModels),
	}
}

func (t *TestDTOManager) ToTestDTOs(appModels []repo.TTest) []TestDTO {
	var testDTOs []TestDTO
	for _, model := range appModels {
		testDTOs = append(testDTOs, t.ToTestDTO(model, nil, nil))
	}

	return testDTOs
}

type AddGeneralDTO struct {
	TestID string `json:"testID" validate:"required,uuid4"`
	Value  string `json:"value"`
}
