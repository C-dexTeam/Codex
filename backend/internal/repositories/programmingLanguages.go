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
	CompileCMD    sql.NullString `db:"compile_cmd"`
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

func (r *PLanguageRepository) dbModelFromAppModel(appModel domains.ProgrammingLanguage) (dbModel dbModelPLanguages) {
	if appModel.GetID() != uuid.Nil {
		dbModel.ID.String = appModel.GetID().String()
		dbModel.ID.Valid = true
	}
	if appModel.GetLanguageID() != uuid.Nil {
		dbModel.LanguageID.String = appModel.GetLanguageID().String()
		dbModel.LanguageID.Valid = true
	}
	if appModel.GetName() != "" {
		dbModel.Name.String = appModel.GetName()
		dbModel.Name.Valid = true
	}
	if appModel.GetDescription() != "" {
		dbModel.Description.String = appModel.GetDescription()
		dbModel.Description.Valid = true
	}
	if appModel.GetDownloadCMD() != "" {
		dbModel.DownloadCMD.String = appModel.GetDownloadCMD()
		dbModel.DownloadCMD.Valid = true
	}
	if appModel.GetCompileCMD() != "" {
		dbModel.CompileCMD.String = appModel.GetCompileCMD()
		dbModel.CompileCMD.Valid = true
	}
	if appModel.GetImagePath() != "" {
		dbModel.ImagePath.String = appModel.GetImagePath()
		dbModel.ImagePath.Valid = true
	}
	if appModel.GetFileExtention() != "" {
		dbModel.FileExtention.String = appModel.GetFileExtention()
		dbModel.FileExtention.Valid = true
	}
	if appModel.GetMonacoEditor() != "" {
		dbModel.MonacoEditor.String = appModel.GetMonacoEditor()
		dbModel.MonacoEditor.Valid = true
	}
	if !appModel.GetCreatedAt().IsZero() {
		dbModel.CreatedAt.Time = appModel.GetCreatedAt()
		dbModel.CreatedAt.Valid = true
	}
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

func (r *PLanguageRepository) Add(ctx context.Context, pLanguage *domains.ProgrammingLanguage) (uuid.UUID, error) {
	dbModel := r.dbModelFromAppModel(*pLanguage)
	query := `
		INSERT INTO
			t_programming_languages (language_id, name, description, download_cmd, compile_cmd, image_path, file_extention, monaco_editor)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	var id uuid.UUID
	err := r.db.QueryRowxContext(
		ctx,
		query,
		dbModel.LanguageID,
		dbModel.Name,
		dbModel.Description,
		dbModel.DownloadCMD,
		dbModel.CompileCMD,
		dbModel.ImagePath,
		dbModel.FileExtention,
		dbModel.MonacoEditor,
	).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *PLanguageRepository) Update(ctx context.Context, pLanguage *domains.ProgrammingLanguage) (err error) {
	dbModel := r.dbModelFromAppModel(*pLanguage)
	query := `
		UPDATE
			t_programming_languages
		SET
			language_id = COALESCE(:language_id, language_id),
			name = COALESCE(:name, name),
			description = COALESCE(:description, description),
			download_cmd =  COALESCE(:download_cmd, download_cmd),
			compile_cmd =  COALESCE(:compile_cmd, compile_cmd),
			image_path =  COALESCE(:image_path, image_path),
			file_extention =  COALESCE(:file_extention, file_extention),
			monaco_editor =  COALESCE(:monaco_editor, monaco_editor)
		WHERE
			id = :id
	`
	_, err = r.db.NamedExecContext(ctx, query, dbModel)
	if err != nil {
		return
	}
	return
}

func (r *PLanguageRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	query := `
		DELETE FROM
			t_programming_languages
		WHERE 
			id = $1
	`
	if _, err = r.db.ExecContext(ctx, query, id); err != nil {
		return
	}

	return
}
