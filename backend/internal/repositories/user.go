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
	ID        sql.NullString `db:"id"`
	PublicKey sql.NullString `db:"public_key"`
	Username  sql.NullString `db:"username"`
	Email     sql.NullString `db:"email"`
	Password  sql.NullString `db:"password"`
}

// dbModelToAppModel converts dbModelUsers to domains.User for application operations (e.g. return to client)
func (r *UserRepository) dbModelToAppModel(dbModel dbModelUsers) (user domains.User) {
	user.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		dbModel.PublicKey.String,
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
	if domModel.GetPublicKey() != "" {
		dbModel.PublicKey.String = domModel.GetPublicKey()
		dbModel.PublicKey.Valid = true
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
	if filter.PublicKey != "" {
		dbFilter.PublicKey.String = filter.PublicKey
		dbFilter.PublicKey.Valid = true
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
		($2::text IS NULL OR public_key = $2::text) AND
		($3::text IS NULL OR username LIKE '%' || $3::text || '%') AND
		($4::text IS NULL OR email LIKE '%' || $4::text || '%')
	LIMIT $5 OFFSET $6
	`

	// Execute the query with the extracted fields
	if err = r.db.SelectContext(ctx, &dbResult, query, dbFilter.ID, dbFilter.PublicKey, dbFilter.Username, dbFilter.Email, limit, (page-1)*limit); err != nil {
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
			t_users (public_key, username, email, password)
		VALUES
			($1, $2, $3, $4)
		RETURNING id
	`

	var id uuid.UUID
	err := tx.QueryRowxContext(ctx, query, dbModel.PublicKey, dbModel.Username, dbModel.Email, dbModel.Password).Scan(&id)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, err
	}

	return id, nil
}

func (r *UserRepository) Update(ctx context.Context, userAuth *domains.User) (err error) {
	dbModel := r.dbModelFromAppModel(*userAuth)
	query := `
		UPDATE
			t_users
		SET
			public_key = COALESCE(:public_key, public_key),
			username = COALESCE(:username, username),
			email = COALESCE(:email, email)
		WHERE
			id = :id
	`
	_, err = r.db.NamedExecContext(ctx, query, dbModel)
	if err != nil {
		return
	}
	return
}
