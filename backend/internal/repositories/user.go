package repositories

import (
	"context"
	"database/sql"

	"github.com/C-dexTeam/codex/internal/domains"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

// dbModelUsers is the struct that represents the user in the database.
type dbModelUsers struct {
	ID       sql.NullString `db:"id"`
	Username sql.NullString `db:"username"`
	Email    sql.NullString `db:"email"`
	Password sql.NullString `db:"password"`
}

// dbModelToAppModel converts dbModelUsers to domains.User for application operations (e.g. return to client)
func (r *UserRepository) dbModelToAppModel(dbModel dbModelUsers) (user domains.User) {
	user.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		dbModel.Username.String,
		dbModel.Email.String,
		dbModel.Password.String,
	)
	return
}

// dbModelFromAppModel converts domains.User to dbModelUsers for database operations (e.g. insert, update)
func (r *UserRepository) dbModelFromAppModel(domModel domains.User) (dbModel dbModelUsers) {
	if domModel.GetID() != uuid.Nil {
		dbModel.ID.String = domModel.GetID().String()
		dbModel.ID.Valid = true
	}
	if domModel.GetEmail() != "" {
		dbModel.Email.String = domModel.GetEmail()
		dbModel.Email.Valid = true
	}
	if domModel.GetUsername() != "" {
		dbModel.Username.String = domModel.GetUsername()
		dbModel.Username.Valid = true
	}
	if domModel.GetPassword() != "" {
		dbModel.Password.String = domModel.GetPassword()
		dbModel.Password.Valid = true
	}

	return
}

// dbModelFromAppFilter converts domains.UserFilter to dbModelUsers for database operations (e.g. select)
func (r *UserRepository) dbModelFromAppFilter(filter domains.UserFilter) (dbFilter dbModelUsers) {
	if filter.ID != uuid.Nil {
		dbFilter.ID.String = filter.ID.String()
		dbFilter.ID.Valid = true
	}
	if filter.Username != "" {
		dbFilter.Username.String = filter.Username
		dbFilter.Username.Valid = true
	}
	if filter.Email != "" {
		dbFilter.Email.String = filter.Email
		dbFilter.Email.Valid = true
	}

	return
}

func NewUserRepository(db *sqlx.DB) domains.IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Filter(ctx context.Context, filter domains.UserFilter, limit, page int64) (users []domains.User, dataCount int64, err error) {
	dbFilter := r.dbModelFromAppFilter(filter)
	dbResult := []dbModelUsers{}

	query := `
	SELECT
		*
	FROM t_users
	WHERE
		($1::uuid IS NULL OR id = $1::uuid) AND
		($2::text IS NULL OR username LIKE '%' || $2::text || '%') AND
		($3::text IS NULL OR email LIKE '%' || $3::text || '%')
	LIMIT $4 OFFSET $5
	`

	// Execute the query with the extracted fields
	if err = r.db.SelectContext(ctx, &dbResult, query, dbFilter.ID, dbFilter.Username, dbFilter.Email, limit, (page-1)*limit); err != nil {
		return
	}
	for _, dbModel := range dbResult {
		users = append(users, r.dbModelToAppModel(dbModel))
	}
	return
}

func (r *UserRepository) AddTx(ctx context.Context, tx *sqlx.Tx, user *domains.User) (uuid.UUID, error) {
	dbModel := r.dbModelFromAppModel(*user)
	query := `
		INSERT INTO
			t_users (username, email, password)
		VALUES
			($1, $2, $3)
		RETURNING id
	`

	var id uuid.UUID
	err := tx.QueryRowxContext(ctx, query, dbModel.Username, dbModel.Email, dbModel.Password).Scan(&id)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, err
	}

	return id, nil
}
