package repositories

import (
	"context"
	"database/sql"

	"github.com/C-dexTeam/codex/internal/domains"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PLanguageRepository struct {
	db *sqlx.DB
}

type dbModelPLanguages struct {
	ID            sql.NullString `db:"id"`
	LanguageID    sql.NullString `db:"language_id"`
	Name          sql.NullString `db:"name"`
	Description   sql.NullString `db:"description"`
	DownloadCMD   sql.NullString `db:"download_cmd"`
	CompileCMD    sql.NullString `db:"compileCMD"`
	ImagePath     sql.NullString `db:"image_path"`
	FileExtention sql.NullString `db:"file_extention"`
	MonacoEditor  sql.NullString `db:"monaco_editor"`
	CreatedAt     sql.NullTime   `db:"created_at"`
}

func (r *PLanguageRepository) dbModelToAppModel(dbModel dbModelPLanguages) (appModel domains.ProgrammingLanguage) {
	appModel.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		uuid.MustParse(dbModel.LanguageID.String),
		dbModel.Name.String,
		dbModel.Description.String,
		dbModel.DownloadCMD.String,
		dbModel.CompileCMD.String,
		dbModel.ImagePath.String,
		dbModel.FileExtention.String,
		dbModel.MonacoEditor.String,
		dbModel.CreatedAt.Time,
	)
	return
}

func (r *PLanguageRepository) dbModelFromAppFilter(filter domains.ProgrammingLanguageFilter) (dbModel dbModelPLanguages) {
	if filter.ID != uuid.Nil {
		dbModel.ID.String = filter.ID.String()
		dbModel.ID.Valid = true
	}
	if filter.LanguageID != uuid.Nil {
		dbModel.LanguageID.String = filter.LanguageID.String()
		dbModel.LanguageID.Valid = true
	}
	if filter.Name != "" {
		dbModel.Name.String = filter.Name
		dbModel.Name.Valid = true
	}
	return
}

func NewPLanguageRepository(db *sqlx.DB) domains.IPLanguagesRepository {
	return &PLanguageRepository{db: db}
}

func (r *PLanguageRepository) Filter(ctx context.Context, filter domains.ProgrammingLanguageFilter, limit, page int64) (pLanguages []domains.ProgrammingLanguage, dataCount int64, err error) {
	dbFilter := r.dbModelFromAppFilter(filter)
	dbResult := []dbModelPLanguages{}

	query := `
	SELECT
		*
	FROM t_programming_languages
	WHERE
		($1::uuid IS NULL OR id = $1::uuid) AND
		($2::uuid IS NULL OR language_id = $2::uuid) AND
		($3::text IS NULL OR name LIKE '%' || $3::text || '%')
	LIMIT $4 OFFSET $5
	`

	// Execute the query with the extracted fields
	if err = r.db.SelectContext(ctx, &dbResult, query, dbFilter.ID, dbFilter.LanguageID, dbFilter.Name, limit, (page-1)*limit); err != nil {
		return
	}
	for _, dbModel := range dbResult {
		pLanguages = append(pLanguages, r.dbModelToAppModel(dbModel))
	}
	return
}
