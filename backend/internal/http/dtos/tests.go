package dto

import "github.com/C-dexTeam/codex/internal/domains"

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

func (t *TestDTOManager) ToInputDTO(appModel domains.Input) InputDTO {
	return InputDTO{
		ID:    appModel.GetID().String(),
		Value: appModel.GetValue(),
	}
}

func (t *TestDTOManager) ToOutputDTO(appModel domains.Output) OutputDTO {
	return OutputDTO{
		ID:    appModel.GetTestID().String(),
		Value: appModel.GetValue(),
	}
}

func (t *TestDTOManager) ToInputDTOs(appModels []domains.Input) []InputDTO {
	var inputDTOs []InputDTO
	for _, model := range appModels {
		inputDTOs = append(inputDTOs, t.ToInputDTO(model))
	}

	return inputDTOs
}

func (t *TestDTOManager) ToOutputDTOs(appModels []domains.Output) []OutputDTO {
	var outputDTOs []OutputDTO
	for _, model := range appModels {
		outputDTOs = append(outputDTOs, t.ToOutputDTO(model))
	}

	return outputDTOs
}

func (t *TestDTOManager) ToTestDTO(appModel domains.Test) TestDTO {
	return TestDTO{
		ID:      appModel.GetID().String(),
		Inputs:  t.ToInputDTOs(appModel.GetInputs()),
		Outputs: t.ToOutputDTOs(appModel.GetOutputs()),
	}
}

func (t *TestDTOManager) ToTestDTOs(appModels []domains.Test) []TestDTO {
	var testDTOs []TestDTO
	for _, model := range appModels {
		testDTOs = append(testDTOs, t.ToTestDTO(model))
	}

	return testDTOs
}

type AddGeneralDTO struct {
	TestID string `json:"value" validate:"required,uuid4"`
	Value  string `json:"value"`
}
