package domains

import (
	"context"

	"github.com/google/uuid"
)

type ILanguagesRepository interface {
	Filter(ctx context.Context, filterModel LanguageFilter, limit, page int64) (languages []Language, dataCount int64, err error)
}

type ILanguagesService interface {
	GetLanguages(ctx context.Context, languageID, value string) (languages []Language, err error)
	GetDefault(ctx context.Context) (language *Language, err error)
}

type Language struct {
	id    uuid.UUID
	value string
}

type LanguageFilter struct {
	ID    uuid.UUID
	Value string
}

const (
	DefaultLanguage      = "EN"
	DefaultLanguageLimit = 8
)

func NewLanguage(value string) (*Language, error) {
	language := &Language{}
	language.SetValue(value)

	return language, nil
}

func (d *Language) Unmarshal(id uuid.UUID, value string) {
	d.id = id
	d.value = value
}

func (d *Language) GetID() uuid.UUID {
	return d.id
}

func (d *Language) GetValue() string {
	return d.value
}

func (d *Language) SetValue(value string) {
	d.value = value
}
