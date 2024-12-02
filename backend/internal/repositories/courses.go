package repositories

import (
	"context"
	"database/sql"

	"github.com/C-dexTeam/codex/internal/domains"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CourseRepository struct {
	db *sqlx.DB
}

type dbModelCourse struct {
	ID           sql.NullString `db:"id"`
	LanguageID   sql.NullString `db:"language_id"`
	PLanguageID  sql.NullString `db:"programming_language_id"`
	RewardID     sql.NullString `db:"reward_id"`
	RewardAmount sql.NullInt64  `db:"reward_amount"`
	Title        sql.NullString `db:"title"`
	Description  sql.NullString `db:"description"`
	ImagePath    sql.NullString `db:"image_path"`
	CreatedAt    sql.NullTime   `db:"created_at"`
	DeletedAt    sql.NullTime   `db:"deleted_at"`
}

func (r *CourseRepository) dbModelToAppModel(dbModel dbModelCourse) (appModel domains.Course) {
	var languageID, planguageID, rewardID *uuid.UUID

	// Dil, programlama dili, ödül ID'leri için aynı şekilde kontrol yapılır
	if parsedLanguageID, err := uuid.Parse(dbModel.LanguageID.String); err == nil {
		languageID = &parsedLanguageID
	} else {
		languageID = nil
	}

	if parsedPLanguageID, err := uuid.Parse(dbModel.PLanguageID.String); err == nil {
		planguageID = &parsedPLanguageID
	} else {
		planguageID = nil
	}

	if parsedRewardID, err := uuid.Parse(dbModel.RewardID.String); err == nil {
		rewardID = &parsedRewardID
	} else {
		rewardID = nil
	}

	// Verileri AppModel'e aktar
	appModel.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		languageID,
		planguageID,
		rewardID,
		int(dbModel.RewardAmount.Int64),
		dbModel.Title.String,
		dbModel.Description.String,
		dbModel.ImagePath.String,
		dbModel.CreatedAt.Time,
		dbModel.DeletedAt.Time,
	)
	return
}

func (r *CourseRepository) dbModelFromAppModel(appModel domains.Course) (dbModel dbModelCourse) {
	if appModel.GetID() != uuid.Nil {
		dbModel.ID.String = appModel.GetID().String()
		dbModel.ID.Valid = true
	}
	if appModel.GetLanguageID() != nil {
		dbModel.LanguageID.String = appModel.GetLanguageID().String()
		dbModel.LanguageID.Valid = true
	}
	if appModel.GetPLanguageID() != nil {
		dbModel.PLanguageID.String = appModel.GetPLanguageID().String()
		dbModel.PLanguageID.Valid = true
	}
	if appModel.GetRewardID() != nil {
		dbModel.RewardID.String = appModel.GetRewardID().String()
		dbModel.RewardID.Valid = true
	}
	if appModel.GetRewardAmount() != 0 {
		dbModel.RewardAmount.Int64 = int64(appModel.GetRewardAmount())
		dbModel.RewardAmount.Valid = true
	}
	if appModel.GetTitle() != "" {
		dbModel.Title.String = appModel.GetTitle()
		dbModel.Title.Valid = true
	}
	if appModel.GetDescription() != "" {
		dbModel.Description.String = appModel.GetDescription()
		dbModel.Description.Valid = true
	}
	if appModel.GetImagePath() != "" {
		dbModel.ImagePath.String = appModel.GetImagePath()
		dbModel.ImagePath.Valid = true
	}
	if !appModel.GetCreatedAt().IsZero() {
		dbModel.CreatedAt.Time = appModel.GetCreatedAt()
		dbModel.CreatedAt.Valid = true
	}
	if !appModel.GetDeletedAt().IsZero() {
		dbModel.DeletedAt.Time = appModel.GetDeletedAt()
		dbModel.DeletedAt.Valid = true
	}

	return
}

func (d *CourseRepository) dbModelFromAppFilter(appFilter domains.CourseFilter) (dbModel dbModelCourse) {
	if appFilter.ID != uuid.Nil {
		dbModel.ID.String = appFilter.ID.String()
		dbModel.ID.Valid = true
	}
	if appFilter.LanguageID != uuid.Nil {
		dbModel.LanguageID.String = appFilter.LanguageID.String()
		dbModel.LanguageID.Valid = true
	}
	if appFilter.PLanguageID != uuid.Nil {
		dbModel.PLanguageID.String = appFilter.PLanguageID.String()
		dbModel.PLanguageID.Valid = true
	}
	if appFilter.Title != "" {
		dbModel.Title.String = appFilter.Title
		dbModel.Title.Valid = true
	}
	if !appFilter.CreatedAt.IsZero() {
		dbModel.CreatedAt.Time = appFilter.CreatedAt
		dbModel.CreatedAt.Valid = true
	}

	return
}

func NewCourseRepository(db *sqlx.DB) domains.ICourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (r *CourseRepository) Filter(ctx context.Context, filter domains.CourseFilter, limit, page int64) (courses []domains.Course, dataCount int64, err error) {
	dbFilter := r.dbModelFromAppFilter(filter)
	dbResult := []dbModelCourse{}

	query := `
	SELECT
		*
	FROM t_courses
	WHERE
		($1::uuid IS NULL OR id = $1::uuid) AND
		($2::uuid IS NULL OR language_id = $2::uuid) AND
		($3::uuid IS NULL OR programming_language_id = $3::uuid) AND
		($4::text IS NULL OR title LIKE '%' || $4::text || '%') AND
		deleted_at IS NULL
	LIMIT $5 OFFSET $6
	`

	if err = r.db.SelectContext(ctx, &dbResult, query, dbFilter.ID, dbFilter.LanguageID, dbFilter.PLanguageID, dbFilter.Title, limit, (page-1)*limit); err != nil {
		return
	}
	for _, dbModel := range dbResult {
		courses = append(courses, r.dbModelToAppModel(dbModel))
	}
	return
}
