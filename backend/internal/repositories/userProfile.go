package repositories

import (
	"context"
	"database/sql"

	"github.com/C-dexTeam/codex/internal/domains"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserProfileRepository struct {
	db *sqlx.DB
}

// dbModelUsersProfile is the struct that represents the user in the database.
type dbModelUsersProfile struct {
	ID                  sql.NullString `db:"id"`
	UserID              sql.NullString `db:"user_id"`
	RoleID              sql.NullString `db:"role_id"`
	Name                sql.NullString `db:"name"`
	Surname             sql.NullString `db:"surname"`
	FirstLogin          sql.NullBool   `db:"first_login"`
	Level               sql.NullInt64  `db:"level"`
	Experience          sql.NullInt64  `db:"experience"`
	NextLevelExperience sql.NullInt64  `db:"next_level_exp"`
	CreatedAt           sql.NullTime   `db:"created_at"`
	DeletedAt           sql.NullTime   `db:"deleted_at"`
}

// dbModelToAppModel converts dbModelUsersProfile to domains.UserProfile for application operations (e.g. return to client)
func (r *UserProfileRepository) dbModelToAppModel(dbModel dbModelUsersProfile) (userProfile domains.UserProfile) {
	userProfile.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		uuid.MustParse(dbModel.UserID.String),
		uuid.MustParse(dbModel.RoleID.String),
		dbModel.Name.String,
		dbModel.Surname.String,
		dbModel.FirstLogin.Bool,
		int(dbModel.Level.Int64),
		int(dbModel.Experience.Int64),
		int(dbModel.NextLevelExperience.Int64),
		dbModel.CreatedAt.Time,
		dbModel.DeletedAt.Time,
	)
	return
}

// dbModelFromAppModel converts domains.UserProfile to dbModelUsersProfile for database operations (e.g. insert, update)
func (r *UserProfileRepository) dbModelFromAppModel(domModel domains.UserProfile) (dbModel dbModelUsersProfile) {
	if domModel.GetID() != uuid.Nil {
		dbModel.ID.String = domModel.GetID().String()
		dbModel.ID.Valid = true
	}
	if domModel.GetUserID() != uuid.Nil {
		dbModel.UserID.String = domModel.GetUserID().String()
		dbModel.UserID.Valid = true
	}
	if domModel.GetRoleID() != uuid.Nil {
		dbModel.RoleID.String = domModel.GetRoleID().String()
		dbModel.RoleID.Valid = true
	}
	if domModel.GetName() != "" {
		dbModel.Name.String = domModel.GetName()
		dbModel.Name.Valid = true
	}
	if domModel.GetSurname() != "" {
		dbModel.Surname.String = domModel.GetSurname()
		dbModel.Surname.Valid = true
	}
	if domModel.GetLevel() != 0 {
		dbModel.Level.Int64 = int64(domModel.GetLevel())
		dbModel.Level.Valid = true
	}
	if domModel.GetExperience() != 0 {
		dbModel.Experience.Int64 = int64(domModel.GetExperience())
		dbModel.Experience.Valid = true
	}
	if domModel.GetNextLevelExperience() != 0 {
		dbModel.NextLevelExperience.Int64 = int64(domModel.GetNextLevelExperience())
		dbModel.NextLevelExperience.Valid = true
	}
	if !domModel.GetCreatedAt().IsZero() {
		dbModel.CreatedAt.Time = domModel.GetCreatedAt()
		dbModel.CreatedAt.Valid = true
	}
	if !domModel.GetDeletedAt().IsZero() {
		dbModel.DeletedAt.Time = domModel.GetDeletedAt()
		dbModel.DeletedAt.Valid = true
	}

	return
}

// dbModelFromAppModel converts domains.UserProfile to dbModelUsersProfile for database operations (e.g. insert, update)
func (r *UserProfileRepository) dbModelFromAppFilter(filter domains.UserProfileFilter) (dbFilter dbModelUsersProfile) {
	if filter.ID != uuid.Nil {
		dbFilter.ID.String = filter.ID.String()
		dbFilter.ID.Valid = true
	}
	if filter.UserID != uuid.Nil {
		dbFilter.UserID.String = filter.UserID.String()
		dbFilter.UserID.Valid = true
	}
	if filter.RoleID != uuid.Nil {
		dbFilter.RoleID.String = filter.RoleID.String()
		dbFilter.RoleID.Valid = true
	}
	if filter.Name != "" {
		dbFilter.Name.String = filter.Name
		dbFilter.Name.Valid = true
	}
	if filter.Surname != "" {
		dbFilter.Surname.String = filter.Surname
		dbFilter.Surname.Valid = true
	}
	if filter.Level != 0 {
		dbFilter.Level.Int64 = int64(filter.Level)
		dbFilter.Level.Valid = true
	}
	if filter.Experience != 0 {
		dbFilter.Experience.Int64 = int64(filter.Experience)
		dbFilter.Experience.Valid = true
	}
	if filter.NextLevelExperience != 0 {
		dbFilter.NextLevelExperience.Int64 = int64(filter.NextLevelExperience)
		dbFilter.NextLevelExperience.Valid = true
	}

	return
}

func NewUserProfileRepository(db *sqlx.DB) domains.IUserProfileRepository {
	return &UserProfileRepository{db: db}
}

func (r *UserProfileRepository) Filter(ctx context.Context, filter domains.UserProfileFilter, limit, page int64) (usersProfile []domains.UserProfile, dataCount int64, err error) {
	dbFilter := r.dbModelFromAppFilter(filter)
	dbResult := []dbModelUsersProfile{}

	query := `
	SELECT
		*
	FROM t_user_profiles
	WHERE
		($1::uuid IS NULL OR id = $1::uuid) AND
		($2::uuid IS NULL OR user_id = $2::uuid) AND
		($3::uuid IS NULL OR role_id = $3::uuid) AND
		($4::text IS NULL OR name LIKE '%' || $4::text || '%') AND
		($5::text IS NULL OR surname LIKE '%' || $5::text || '%')
	LIMIT $6 OFFSET $7
	`

	// Execute the query with the extracted fields
	if err = r.db.SelectContext(ctx, &dbResult, query, dbFilter.ID, dbFilter.UserID, dbFilter.RoleID, dbFilter.Name, dbFilter.Surname, limit, (page-1)*limit); err != nil {
		return
	}
	for _, dbModel := range dbResult {
		usersProfile = append(usersProfile, r.dbModelToAppModel(dbModel))
	}
	return
}

func (r *UserProfileRepository) AddTx(ctx context.Context, tx *sqlx.Tx, userProfile *domains.UserProfile) error {
	dbModel := r.dbModelFromAppModel(*userProfile)
	query := `
		INSERT INTO
			t_user_profiles (user_id, role_id, name, surname)
		VALUES
			(:user_id, :role_id, :name, :surname)
	`

	_, err := tx.NamedExecContext(ctx, query, dbModel)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (r *UserProfileRepository) Update(ctx context.Context, userProfile *domains.UserProfile) (err error) {
	dbModel := r.dbModelFromAppModel(*userProfile)
	query := `
		UPDATE
         t_user_profiles
		SET
			name = COALESCE(:name, name),
			surname =  COALESCE(:surname, surname),
			first_login = :first_login
		WHERE
			id = :id

	`
	_, err = r.db.NamedExecContext(ctx, query, dbModel)
	if err != nil {
		return
	}
	return
}

func (r *UserProfileRepository) ChangeRole(ctx context.Context, userProfile *domains.UserProfile) error {
	dbModel := r.dbModelFromAppModel(*userProfile)
	query := `
        UPDATE
            t_user_profiles
        SET
            role_id = :role_id
        WHERE
            id = :id
    `

	_, err := r.db.NamedExecContext(ctx, query, dbModel)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserProfileRepository) AddExp(ctx context.Context, userProfile *domains.UserProfile) error {
	dbModel := r.dbModelFromAppModel(*userProfile)
	query := `
        UPDATE
            t_user_profiles
        SET
			level =  COALESCE(:level, level),
			experience =  COALESCE(:experience, experience),
			next_level_Exp =  COALESCE(:next_level_Exp, next_level_Exp),
        WHERE
            id = :id
    `

	_, err := r.db.NamedExecContext(ctx, query, dbModel)
	if err != nil {
		return err
	}

	return nil
}
