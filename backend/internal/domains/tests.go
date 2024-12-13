package domains

import (
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type Test struct {
	inputs  []Input
	outputs []Output
}

type Input struct {
	id     uuid.UUID
	testID uuid.UUID
	value  string
}

type InputFilter struct {
	ID      uuid.UUID
	InputID uuid.UUID
}

type Output struct {
	inputID uuid.UUID
	value   string
}

type OutputFilter struct {
	InputID uuid.UUID
}

func NewTest(
	inputs []Input,
	outputs []Output,
) (test *Test, err error) {
	test.SetInputs(inputs)
	test.SetOutputs(outputs)

	return
}

func (t *Test) Unmarshal(
	inputs []Input,
	outputs []Output,
) {
	t.inputs = inputs
	t.outputs = outputs
}

func NewInput(
	id, testID, value string,
) (input *Input, err error) {
	if err := input.SetID(id); err != nil {
		return nil, err
	}
	if err := input.SetTestID(testID); err != nil {
		return nil, err
	}
	input.SetValue(value)

	return
}

func (i *Input) Unmarshal(
	id, testID uuid.UUID,
	value string,
) {
	i.id = id
	i.testID = testID
	i.value = value
}

func NewOutput(
	inputID, value string,
) (output *Output, err error) {
	if err := output.SetInputID(inputID); err != nil {
		return nil, err
	}
	output.SetValue(value)

	return
}

func (o *Output) Unmarshal(
	inputID uuid.UUID,
	value string,
) {
	o.inputID = inputID
	o.value = value
}

// FOR TEST - Getter
func (t *Test) GetInputs() []Input {
	return t.inputs
}

func (t *Test) GetOutputs() []Output {
	return t.outputs
}

// FOR TEST - Setter
func (t *Test) SetInputs(inputs []Input) {
	t.inputs = inputs
}

func (t *Test) SetOutputs(outputs []Output) {
	t.outputs = outputs
}

// FOR INPUT - Getter
func (i *Input) GetID() uuid.UUID {
	return i.id
}

func (i *Input) GetTestID() uuid.UUID {
	return i.testID
}

func (i *Input) GetValue() string {
	return i.value
}

// FOR INPUT - Setter
func (i *Input) SetID(id string) error {
	if id != "" {
		idUUID, err := uuid.Parse(id)
		if err != nil {
			return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
		}
		i.id = idUUID
	}

	return nil
}

func (i *Input) SetTestID(testID string) error {
	if testID != "" {
		idUUID, err := uuid.Parse(testID)
		if err != nil {
			return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
		}
		i.testID = idUUID
	}

	return nil
}

func (i *Input) SetValue(value string) {
	i.value = value
}

// FOR OUTPUT - Getter
func (o *Output) GetInputID() uuid.UUID {
	return o.inputID
}

func (o *Output) GetValue() string {
	return o.value
}

// FOR OUTPUT - Setter
func (o *Output) SetInputID(inputID string) error {
	if inputID != "" {
		idUUID, err := uuid.Parse(inputID)
		if err != nil {
			return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
		}
		o.inputID = idUUID
	}

	return nil
}

func (o *Output) SetValue(value string) {
	o.value = value
}
