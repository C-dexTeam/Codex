package repositories

import (
	"context"
	"database/sql"

	"github.com/C-dexTeam/codex/internal/domains"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type LanguagesRepository struct {
	db *sqlx.DB
}

type dbModelLanguages struct {
	ID    sql.NullString `db:"id"`
	Value sql.NullString `db:"value"`
}

func (r *LanguagesRepository) dbModelToAppModel(dbModel dbModelLanguages) (language domains.Languages) {
	language.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		dbModel.Value.String,
	)
	return
}

func (r *LanguagesRepository) appModelToDBModel(appModel domains.Languages) (dbModel dbModelLanguages) {
	if appModel.GetID() != uuid.Nil {
		dbModel.ID.String = appModel.GetID().String()
		dbModel.ID.Valid = true
	}
	if appModel.GetValue() != "" {
		dbModel.Value.String = appModel.GetValue()
		dbModel.Value.Valid = true
	}

	return
}

func NewLanguageRepository(db *sqlx.DB) domains.ILanguagesRepository {
	return &LanguagesRepository{db: db}
}

func (r *LanguagesRepository) Filter(ctx context.Context, appModel domains.Languages, limit, page int64) (languages []domains.Languages, dataCount int64, err error) {
	dbModel := r.appModelToDBModel(appModel)
	dbResult := []dbModelLanguages{}

	query := `
	SELECT
		*
	FROM t_languages
	WHERE
		($1::uuid IS NULL OR id = $1::uuid) AND
		($2::text IS NULL OR value LIKE '%' || $2::text || '%')
	LIMIT $3 OFFSET $4
	`
	if err = r.db.SelectContext(ctx, &dbResult, query, dbModel.ID, dbModel.Value, limit, (page-1)*limit); err != nil {
		return
	}
	for _, dbModel := range dbResult {
		languages = append(languages, r.dbModelToAppModel(dbModel))
	}
	return
}
