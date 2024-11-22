package services

import (
	"context"

	"github.com/C-dexTeam/codex/internal/domains"
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
		Name: domains.DefaultRole,
	}, 1, 1)
	if len(roles) != 1 {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(500, "Default role not found", err)
	}
	role = &roles[0]
	return
}

func (s *roleService) GetRoleByID(ctx context.Context, roleID uuid.UUID) (role *domains.Role, err error) {
	roles, _, err := s.roleRepository.Filter(ctx, domains.RoleFilter{
		ID: roleID,
	}, 1, 1)
	if len(roles) != 1 {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(500, "Default role not found", err)
	}
	role = &roles[0]

	return
}

func (s *roleService) GetByName(ctx context.Context, name string) (role *domains.Role, err error) {
	roles, _, err := s.roleRepository.Filter(ctx, domains.RoleFilter{
		Name: name,
	}, 1, 1)
	if len(roles) != 1 {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(404, "role not found", err)
	}
	role = &roles[0]

	return
}
