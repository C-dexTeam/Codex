package services

import (
	"context"
	"database/sql"

	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	repo "github.com/C-dexTeam/codex/internal/repos/out"

	"github.com/google/uuid"
)

type RoleService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func newRoleService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *RoleService {
	return &RoleService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *RoleService) GetDefault(ctx context.Context) (*repo.TRole, error) {
	return s.GetByName(ctx, s.utilService.D().Roles.DefaultRole)
}

func (s *RoleService) GetRoleByID(ctx context.Context, roleID uuid.UUID) (*repo.TRole, error) {
	role, err := s.queries.GetRoleByID(ctx, roleID)
	if role.ID == uuid.Nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusNotFound, errorDomains.ErrRoleNotFound, err)
	}

	return &role, nil
}

func (s *RoleService) GetByName(ctx context.Context, name string) (*repo.TRole, error) {
	role, err := s.queries.GetRoleByName(ctx, name)
	if role.ID == uuid.Nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusNotFound, errorDomains.ErrRoleNotFound, err)
	}

	return &role, nil
}
