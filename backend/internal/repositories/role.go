package repositories

import (
	"context"
	"database/sql"

	"github.com/C-dexTeam/codex/internal/domains"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RoleRepository struct {
	db *sqlx.DB
}

// dbModelRoles is the struct that represents the role in the database.
type dbModelRoles struct {
	ID   sql.NullString `db:"id"`
	Name sql.NullString `db:"name"`
}

// dbModelToAppModel converts dbModelRoles to domains.role for application operations (e.g. return to client)
func (r *RoleRepository) dbModelToAppModel(dbModel dbModelRoles) (role domains.Role) {
	role.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		dbModel.Name.String,
	)
	return
}

// dbModelFromAppFilter converts domains.RoleFilter to dbModelRole for database operations (e.g. select)
func (r *RoleRepository) dbModelFromAppFilter(filter domains.RoleFilter) (dbFilter dbModelRoles) {
	if filter.ID != uuid.Nil {
		dbFilter.ID.String = filter.ID.String()
		dbFilter.ID.Valid = true
	}
	if filter.Name != "" {
		dbFilter.Name.String = filter.Name
		dbFilter.Name.Valid = true
	}

	return
}

func NewRoleRepository(db *sqlx.DB) domains.IRoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) Filter(ctx context.Context, filter domains.RoleFilter, limit, page int64) (roles []domains.Role, dataCount int64, err error) {
	dbFilter := r.dbModelFromAppFilter(filter)
	dbResult := []dbModelRoles{}

	query := `
	SELECT
		*
	FROM t_roles
	WHERE
		($1::uuid IS NULL OR id = $1::uuid) AND
		($2::text IS NULL OR name LIKE '%' || $2::text || '%')
	LIMIT $3 OFFSET $4
	`

	// Execute the query with the extracted fields
	if err = r.db.SelectContext(ctx, &dbResult, query, dbFilter.ID, dbFilter.Name, limit, (page-1)*limit); err != nil {
		return
	}
	for _, dbModel := range dbResult {
		roles = append(roles, r.dbModelToAppModel(dbModel))
	}
	return
}
