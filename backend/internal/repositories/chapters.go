package repositories

import (
	"context"
	"database/sql"

	"github.com/C-dexTeam/codex/internal/domains"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ChapterRepository struct {
	db *sqlx.DB
}

type dbModelChapter struct {
	ID               sql.NullString `db:"id"`
	CourseID         sql.NullString `db:"course_id"`
	LanguageID       sql.NullString `db:"language_id"`
	RewardID         sql.NullString `db:"reward_id"`
	RewardAmount     sql.NullInt64  `db:"reward_amount"`
	Title            sql.NullString `db:"title"`
	Description      sql.NullString `db:"description"`
	Content          sql.NullString `db:"content"`
	FuncName         sql.NullString `db:"func_name"`
	FrontendTmp      sql.NullString `db:"frontend_template"`
	DockerTmp        sql.NullString `db:"docker_template"`
	CheckTmp         sql.NullString `db:"check_template"`
	GrantsExperience sql.NullBool   `db:"grants_experience"`
	Active           sql.NullBool   `db:"active"`
	CreatedAt        sql.NullTime   `db:"created_at"`
	DeletedAt        sql.NullTime   `db:"deleted_at"`
}

func (r *ChapterRepository) dbModelToAppModel(dbModel dbModelChapter) (appModel domains.Chapter) {
	var rewardID *uuid.UUID

	if parsedRewardID, err := uuid.Parse(dbModel.RewardID.String); err == nil {
		rewardID = &parsedRewardID
	} else {
		rewardID = nil
	}

	appModel.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		uuid.MustParse(dbModel.LanguageID.String),
		uuid.MustParse(dbModel.CourseID.String),
		rewardID,
		int(dbModel.RewardAmount.Int64),
		dbModel.Title.String,
		dbModel.Description.String,
		dbModel.Content.String,
		dbModel.FuncName.String,
		dbModel.FrontendTmp.String,
		dbModel.DockerTmp.String,
		dbModel.CheckTmp.String,
		dbModel.GrantsExperience.Bool,
		dbModel.Active.Bool,
		dbModel.CreatedAt.Time,
		&dbModel.DeletedAt.Time,
	)

	return
}

func (r *ChapterRepository) dbModelFromAppModel(appModel domains.Chapter) (dbModel dbModelChapter) {
	if appModel.GetID() != uuid.Nil {
		dbModel.ID.String = appModel.GetID().String()
		dbModel.ID.Valid = true
	}
	if appModel.GetLanguageID() != uuid.Nil {
		dbModel.LanguageID.String = appModel.GetLanguageID().String()
		dbModel.LanguageID.Valid = true
	}
	if appModel.GetCourseID() != uuid.Nil {
		dbModel.CourseID.String = appModel.GetCourseID().String()
		dbModel.CourseID.Valid = true
	}
	if appModel.GetRewardID() != nil {
		dbModel.RewardID.String = appModel.GetRewardID().String()
		dbModel.RewardID.Valid = true
	}
	if appModel.GetTitle() != "" {
		dbModel.Title.String = appModel.GetTitle()
		dbModel.Title.Valid = true
	}
	if appModel.GetDescription() != "" {
		dbModel.Description.String = appModel.GetDescription()
		dbModel.Description.Valid = true
	}
	if appModel.GetFuncName() != "" {
		dbModel.FuncName.String = appModel.GetFuncName()
		dbModel.FuncName.Valid = true
	}
	if appModel.GetFrontendTmp() != "" {
		dbModel.FrontendTmp.String = appModel.GetFrontendTmp()
		dbModel.FrontendTmp.Valid = true
	}
	if appModel.GetDockerTmp() != "" {
		dbModel.DockerTmp.String = appModel.GetDockerTmp()
		dbModel.DockerTmp.Valid = true
	}
	if appModel.GetCheckTmp() != "" {
		dbModel.CheckTmp.String = appModel.GetCheckTmp()
		dbModel.CheckTmp.Valid = true
	}
	if !appModel.GetCreatedAt().IsZero() {
		dbModel.CreatedAt.Time = appModel.GetCreatedAt()
		dbModel.CreatedAt.Valid = true
	}
	if !appModel.GetDeletedAt().IsZero() {
		dbModel.DeletedAt.Time = *appModel.GetDeletedAt()
		dbModel.DeletedAt.Valid = true
	}
	dbModel.GrantsExperience.Bool = appModel.GetGrantsExperience()
	dbModel.GrantsExperience.Valid = true
	dbModel.Active.Bool = appModel.GetActive()
	dbModel.Active.Valid = true

	return
}

func (d *ChapterRepository) dbModelFromAppFilter(appFilter domains.ChapterFilter) (dbModel dbModelChapter) {
	if appFilter.ID != uuid.Nil {
		dbModel.ID.String = appFilter.ID.String()
		dbModel.ID.Valid = true
	}
	if appFilter.LanguageID != uuid.Nil {
		dbModel.LanguageID.String = appFilter.LanguageID.String()
		dbModel.LanguageID.Valid = true
	}
	if appFilter.CourseID != uuid.Nil {
		dbModel.CourseID.String = appFilter.CourseID.String()
		dbModel.CourseID.Valid = true
	}
	if appFilter.RewardID != uuid.Nil {
		dbModel.RewardID.String = appFilter.RewardID.String()
		dbModel.RewardID.Valid = true
	}
	if appFilter.Title != "" {
		dbModel.Title.String = appFilter.Title
		dbModel.Title.Valid = true
	}
	if appFilter.GrantsExperience != nil {
		dbModel.GrantsExperience.Bool = *appFilter.GrantsExperience
		dbModel.GrantsExperience.Valid = true
	}
	if appFilter.Active != nil {
		dbModel.Active.Bool = *appFilter.Active
		dbModel.Active.Valid = true
	}

	return
}

func NewChapterRepository(db *sqlx.DB) domains.IChapterRepository {
	return &ChapterRepository{db: db}
}

func (r *ChapterRepository) Filter(ctx context.Context, filter domains.ChapterFilter, limit, page int64) (chapters []domains.Chapter, dataCount int64, err error) {
	dbFilter := r.dbModelFromAppFilter(filter)
	dbResult := []dbModelChapter{}

	query := `
	SELECT
		*
	FROM
		t_chapters
	WHERE
		($1::uuid IS NULL OR id = $1::uuid) AND
		($2::uuid IS NULL OR language_id = $2::uuid) AND
		($3::uuid IS NULL OR course_id = $3::uuid) AND
		($4::uuid IS NULL OR reward_id = $4::uuid) AND
		($5::text IS NULL OR title LIKE '%' || $5::text || '%') AND
		($6::boolean IS NULL OR grants_experience = $6::boolean) AND
		($7::boolean IS NULL OR active = $7::boolean) AND
		deleted_at IS NULL
	LIMIT $8 OFFSET $9;
	`

	if err = r.db.SelectContext(ctx, &dbResult, query, dbFilter.ID, dbFilter.LanguageID, dbFilter.CourseID, dbFilter.RewardID, dbFilter.Title, dbFilter.GrantsExperience, dbFilter.Active, limit, (page-1)*limit); err != nil {
		return
	}
	for _, dbModel := range dbResult {
		chapters = append(chapters, r.dbModelToAppModel(dbModel))
	}
	return
}
