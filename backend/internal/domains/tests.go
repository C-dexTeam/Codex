package domains

import "github.com/google/uuid"

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
	id, testID uuid.UUID,
	value string,
) (input *Input, err error) {
	input.SetID(id)
	input.SetTestID(testID)
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
	inputID uuid.UUID,
	value string,
) (output *Output, err error) {
	output.SetInputID(inputID)
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
func (i *Input) SetID(id uuid.UUID) {
	i.id = id
}

func (i *Input) SetTestID(testID uuid.UUID) {
	i.testID = testID
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
func (o *Output) SetInputID(inputID uuid.UUID) {
	o.inputID = inputID
}

func (o *Output) SetValue(value string) {
	o.value = value
}
