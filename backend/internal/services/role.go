package services

import (
	"context"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"

	"github.com/google/uuid"
)

type roleService struct {
	roleRepository domains.IRoleRepository
}

func newRoleService(
	roleRepository domains.IRoleRepository,
) domains.IRoleService {
	return &roleService{
		roleRepository: roleRepository,
	}
}

func (s *roleService) GetDefault(ctx context.Context) (role *domains.Role, err error) {
	roles, _, err := s.roleRepository.Filter(ctx, domains.RoleFilter{
		Name: domains.RoleDefaultRole,
	}, 1, 1)
	if len(roles) != 1 {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusNotFound, errorDomains.ErrRoleNotFound, err)
	}
	role = &roles[0]
	return
}

func (s *roleService) GetRoleByID(ctx context.Context, roleID uuid.UUID) (role *domains.Role, err error) {
	roles, _, err := s.roleRepository.Filter(ctx, domains.RoleFilter{
		ID: roleID,
	}, 1, 1)
	if len(roles) != 1 {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusNotFound, errorDomains.ErrRoleNotFound, err)
	}
	role = &roles[0]

	return
}

func (s *roleService) GetByName(ctx context.Context, name string) (role *domains.Role, err error) {
	roles, _, err := s.roleRepository.Filter(ctx, domains.RoleFilter{
		Name: name,
	}, 1, 1)
	if len(roles) != 1 {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusNotFound, errorDomains.ErrRoleNotFound, err)
	}
	role = &roles[0]

	return
}
