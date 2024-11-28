package domains

import (
	"context"

	"github.com/google/uuid"
)

type ILanguagesRepository interface {
	Filter(ctx context.Context, filterModel LanguagesFilter, limit, page int64) (languages []Languages, dataCount int64, err error)
}

type ILanguagesService interface {
	GetLanguages(ctx context.Context, languageID, value string) (languages []Languages, err error)
	GetDefault(ctx context.Context) (language *Languages, err error)
}

type Languages struct {
	id    uuid.UUID
	value string
}

type LanguagesFilter struct {
	ID    uuid.UUID
	Value string
}

const (
	DefaultLanguage      = "EN"
	DefaultLanguageLimit = 8
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
