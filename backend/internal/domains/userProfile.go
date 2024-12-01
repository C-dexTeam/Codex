package domains

import (
	"context"
	"time"

	serviceErrors "github.com/C-dexTeam/codex/internal/errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// IUserProfileRepository is the interface that provides the methods for the user repository.
type IUserProfileRepository interface {
	Filter(ctx context.Context, filter UserProfileFilter, limit, page int64) (usersProfile []UserProfile, dataCount int64, err error)
	Update(ctx context.Context, userProfile *UserProfile) (err error)
	AddTx(ctx context.Context, tx *sqlx.Tx, userProfile *UserProfile) error
	ChangeRole(ctx context.Context, userProfile *UserProfile) error
}

// IUserProfileService is the interface that provides the methods for the user service.
type IUserProfileService interface {
	GetAllUsersProfile(ctx context.Context, id, userID, roleID, name, surname, page, limit string) ([]UserProfile, error)
	Update(ctx context.Context, userProfileID, name, surname string) (err error)
	ChangeUserRole(ctx context.Context, userProfileID, newRoleID string) (err error)
}

// User represents a user entity.
type UserProfile struct {
	id         uuid.UUID
	userID     uuid.UUID
	roleID     uuid.UUID
	name       string
	surname    string
	firstLogin bool
	createdAt  time.Time
	deletedAt  time.Time
}

// UserProfileFilter is the struct that represents user's uniques.
type UserProfileFilter struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	RoleID  uuid.UUID
	Name    string
	Surname string
}

// NewUserProfile creates a new user.
func NewUserProfile(
	userID, roleID, name, surname string,
	firstLogin bool,
) (*UserProfile, error) {
	userProfile := &UserProfile{}
	if err := userProfile.SetUserID(userID); err != nil {
		return nil, err
	}
	if err := userProfile.SetRoleID(roleID); err != nil {
		return nil, err
	}

	userProfile.SetName(name)
	userProfile.SetSurname(name)
	userProfile.SetFirstLogin(firstLogin)

	return userProfile, nil
}

// Unmarshal unmarshals the userProfile for database operations.
func (d *UserProfile) Unmarshal(
	id, userID, roleID uuid.UUID,
	name, surname string,
	firstLogin bool,
	createdAt, deletedAt time.Time,
) {
	d.id = id
	d.userID = userID
	d.roleID = roleID
	d.name = name
	d.surname = surname
	d.firstLogin = firstLogin
	d.createdAt = createdAt
	d.deletedAt = deletedAt
}

// Getter Functions
func (d *UserProfile) GetID() uuid.UUID {
	return d.id
}

func (d *UserProfile) GetUserID() uuid.UUID {
	return d.userID
}

func (d *UserProfile) GetRoleID() uuid.UUID {
	return d.roleID
}

func (d *UserProfile) GetName() string {
	return d.name
}

func (d *UserProfile) GetSurname() string {
	return d.surname
}

func (d *UserProfile) GetCreatedAt() time.Time {
	return d.createdAt
}

func (d *UserProfile) GetDeletedAt() time.Time {
	return d.deletedAt
}

func (d *UserProfile) GetFirstLogin() bool {
	return d.firstLogin
}

// Setter Functions
func (d *UserProfile) SetUserID(userID string) error {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(400, "Invalid user id", err)
	}
	d.userID = userUUID

	return nil
}

func (d *UserProfile) SetRoleID(roleID string) error {
	roleUUID, err := uuid.Parse(roleID)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(400, "Invalid role id", err)
	}
	d.roleID = roleUUID

	return nil
}

func (d *UserProfile) SetName(name string) {
	d.name = name
}

func (d *UserProfile) SetSurname(surname string) {
	d.surname = surname
}

func (d *UserProfile) SetFirstLogin(firstLogin bool) {
	d.firstLogin = firstLogin
}
