package services

import (
	"context"
	"database/sql"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	repo "github.com/C-dexTeam/codex/internal/repos/out"

	"github.com/google/uuid"
)

type roleService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func newRoleService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *roleService {
	return &roleService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *roleService) GetDefault(ctx context.Context) (*repo.TRole, error) {
	return s.GetByName(ctx, domains.RoleDefaultRole)
}

func (s *roleService) GetRoleByID(ctx context.Context, roleID uuid.UUID) (*repo.TRole, error) {
	role, err := s.queries.GetRoleByID(ctx, roleID)
	if role.ID == uuid.Nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusNotFound, errorDomains.ErrRoleNotFound, err)
	}

	return &role, nil
}

func (s *roleService) GetByName(ctx context.Context, name string) (*repo.TRole, error) {
	role, err := s.queries.GetRoleByName(ctx, name)
	if role.ID == uuid.Nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusNotFound, errorDomains.ErrRoleNotFound, err)
	}

	return &role, nil
}
