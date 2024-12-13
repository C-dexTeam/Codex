package domains

import (
	"context"

	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type ITestRepository interface {
	FilterTest(ctx context.Context, filter TestFilter, limit, page int64) (tests []Test, dataCount int64, err error)
	FilterInput(ctx context.Context, filter GeneralFilter, limit, page int64) (inputs []Input, dataCount int64, err error)
	FilterOutput(ctx context.Context, filter GeneralFilter, limit, page int64) (outputs []Output, dataCount int64, err error)

	AddTest(ctx context.Context, input *Test) (uuid.UUID, error)
	AddInput(ctx context.Context, input *Input) (uuid.UUID, error)
	AddOutput(ctx context.Context, output *Output) (uuid.UUID, error)
}
type ITestService interface {
	GetTests(
		ctx context.Context,
		id, chapterID, page, limit string,
	) (tests []Test, err error)
}

const (
	DefaultTestLimit = 10
)

type Test struct {
	id        uuid.UUID
	chapterID uuid.UUID
	inputs    []Input
	outputs   []Output
}

type TestFilter struct {
	ID        uuid.UUID
	ChapterID uuid.UUID
}

type Input struct {
	id     uuid.UUID
	testID uuid.UUID
	value  string
}

type GeneralFilter struct {
	ID     uuid.UUID
	TestID uuid.UUID
}

type Output struct {
	id     uuid.UUID
	testID uuid.UUID
	value  string
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
	id, chapterID uuid.UUID,
	inputs []Input,
	outputs []Output,
) {
	t.id = id
	t.chapterID = chapterID
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
	testID, value string,
) (output *Output, err error) {
	if err := output.SetTestID(testID); err != nil {
		return nil, err
	}
	output.SetValue(value)

	return
}

func (o *Output) Unmarshal(
	id, testID uuid.UUID,
	value string,
) {
	o.id = id
	o.testID = testID
	o.value = value
}

// FOR TEST - Getter
func (t *Test) GetID() uuid.UUID {
	return t.id
}

func (t *Test) ChapterID() uuid.UUID {
	return t.chapterID
}

func (t *Test) GetInputs() []Input {
	return t.inputs
}

func (t *Test) GetOutputs() []Output {
	return t.outputs
}

// FOR TEST - Setter
func (t *Test) SetID(id string) error {
	if id != "" {
		idUUID, err := uuid.Parse(id)
		if err != nil {
			return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
		}
		t.id = idUUID
	}

	return nil
}

func (t *Test) SetChapterID(id string) error {
	if id != "" {
		idUUID, err := uuid.Parse(id)
		if err != nil {
			return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
		}
		t.id = idUUID
	}

	return nil
}

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
func (o *Output) GetID() uuid.UUID {
	return o.id
}

func (o *Output) GetTestID() uuid.UUID {
	return o.testID
}

func (o *Output) GetValue() string {
	return o.value
}

// FOR OUTPUT - Setter
func (o *Output) SetTestID(testID string) error {
	if testID != "" {
		idUUID, err := uuid.Parse(testID)
		if err != nil {
			return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID)
		}
		o.testID = idUUID
	}

	return nil
}

func (o *Output) SetValue(value string) {
	o.value = value
}
