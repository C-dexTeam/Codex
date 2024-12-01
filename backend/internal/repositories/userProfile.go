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
	ID         sql.NullString `db:"id"`
	UserID     sql.NullString `db:"user_id"`
	RoleID     sql.NullString `db:"role_id"`
	Name       sql.NullString `db:"name"`
	FirstLogin sql.NullBool   `db:"first_login"`
	Surname    sql.NullString `db:"surname"`
	CreatedAt  sql.NullTime   `db:"created_at"`
	DeletedAt  sql.NullTime   `db:"deleted_at"`
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
         t_users
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
            t_users
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
