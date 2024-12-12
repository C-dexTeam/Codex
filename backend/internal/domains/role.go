package domains

import (
	"context"

	serviceErrors "github.com/C-dexTeam/codex/internal/errors"

	"github.com/google/uuid"
)

// IRoleRepository is the interface that provides the methods for the role repository.
type IRoleRepository interface {
	Filter(ctx context.Context, filter RoleFilter, limit, page int64) (roles []Role, dataCount int64, err error)
}

// IRoleService is the interface that provides the methods for the role service.
type IRoleService interface {
	GetDefault(ctx context.Context) (role *Role, err error)
	GetByName(ctx context.Context, name string) (role *Role, err error)
	GetRoleByID(ctx context.Context, roleID uuid.UUID) (role *Role, err error)
}

const (
	RoleDefaultRole = "user"
	RoleWalletUser  = "wallet-user"
	RoleAdmin       = "admin"
	RolePublic      = "public"
)

// Role represents a role entity.
type Role struct {
	id   uuid.UUID
	name string
}

// RoleFilter is the struct that represents role's uniques.
type RoleFilter struct {
	ID   uuid.UUID
	Name string
}

// NewUser creates a new user.
func NewRole(name string) (*Role, error) {
	role := &Role{}
	if err := role.SetName(name); err != nil {
		return nil, err
	}

	return role, nil
}

// Unmarshal unmarshals the user for database operations.
func (d *Role) Unmarshal(id uuid.UUID, name string) {
	d.id = id
	d.name = name
}

// Setter Functions
func (d *Role) SetName(name string) error {
	if name == "" {
		return serviceErrors.NewServiceErrorWithMessage(400, "name is required")
	}
	if len(name) < 3 {
		return serviceErrors.NewServiceErrorWithMessage(400, "name must be at least 3 characters")
	} else if len(name) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(400, "name must be at most 30 characters")
	}
	d.name = name
	return nil
}

// Getter Functions
func (d *Role) GetID() uuid.UUID {
	return d.id
}

func (d *Role) GetName() string {
	return d.name
}
