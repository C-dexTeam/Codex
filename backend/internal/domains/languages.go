package domains

import "github.com/google/uuid"

type ILanguagesRepository interface {
}

type ILanguagesService interface {
}

type Languages struct {
	id    uuid.UUID
	value string
}

const (
	DefaultLanguage = "EN"
)

func NewLanguage(value string) (*Languages, error) {
	language := &Languages{}
	language.SetValue(value)

	return language, nil
}

func (d *Languages) Unmarshal(id uuid.UUID, value string) {
	d.id = id
	d.value = value
}

func (d *Languages) GetID() uuid.UUID {
	return d.id
}

func (d *Languages) GetValue() string {
	return d.value
}

func (d *Languages) SetValue(value string) {
	d.value = value
}
