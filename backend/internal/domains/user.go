package domains

import (
	"context"

	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	hasherService "github.com/C-dexTeam/codex/pkg/hasher"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// IUserRepository is the interface that provides the methods for the user repository.
type IUserRepository interface {
	Filter(ctx context.Context, filter UserFilter, limit, page int64) (users []User, dataCount int64, err error)
	AddTx(ctx context.Context, tx *sqlx.Tx, user *User) (uuid.UUID, error)
	Update(ctx context.Context, userAuth *User) (err error)
}

// IUserService is the interface that provides the methods for the user service.
type IUserService interface {
	Login(ctx context.Context, username, password string) (user *User, err error)
	Register(ctx context.Context, username, email, password, confirmPassword string, defaultRoleID uuid.UUID) (err error)
	AuthWallet(ctx context.Context, publicKey, message, signature string, defaultRoleID uuid.UUID) (user *User, err error)
	ConnectWallet(ctx context.Context, userAuthID, publicKey, message, signature string) (err error)
}

// User represents a user entity.
type User struct {
	id        uuid.UUID
	publicKey string
	username  string
	email     string
	password  string
}

// UserFilter is the struct that represents user's uniques.
type UserFilter struct {
	ID        uuid.UUID
	PublicKey string
	Username  string
	Email     string
}

// NewUser creates a new user.
func NewUser(publicKey, username, email, password string) (*User, error) {
	user := &User{}
	if err := user.SetUsername(username); err != nil {
		return nil, err
	}
	if err := user.SetEmail(email); err != nil {
		return nil, err
	}
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}
	user.SetPublicKey(publicKey)

	return user, nil
}

// Unmarshal unmarshals the user for database operations.
func (d *User) Unmarshal(id uuid.UUID, publicKey, username, email, password string) {
	d.id = id
	d.publicKey = publicKey
	d.username = username
	d.email = email
	d.password = password
}

// Getter Functions
func (d *User) GetID() uuid.UUID {
	return d.id
}

func (d *User) GetPublicKey() string {
	return d.publicKey
}

func (d *User) GetUsername() string {
	return d.username
}

func (d *User) GetEmail() string {
	return d.email
}

func (d *User) GetPassword() string {
	return d.password
}

// Setter Functions
func (d *User) SetPublicKey(publicKey string) {
	d.publicKey = publicKey
}

func (d *User) SetUsername(username string) error {
	if len(username) < 3 && len(username) != 0 {
		return serviceErrors.NewServiceErrorWithMessage(400, "username must be at least 3 characters")
	} else if len(username) > 30 {
		return serviceErrors.NewServiceErrorWithMessage(400, "username must be at most 30 characters")
	}
	d.username = username
	return nil
}

func (d *User) SetEmail(email string) error {
	if len(email) > 40 {
		return serviceErrors.NewServiceErrorWithMessage(400, "email must be at most 40 characters")
	}
	d.email = email

	return nil
}

func (d *User) SetPassword(password string) error {
	if password == "" {
		return serviceErrors.NewServiceErrorWithMessage(400, "password is required")
	}
	if len(password) < 8 {
		return serviceErrors.NewServiceErrorWithMessage(400, "password must be at least 8 characters")
	}

	hashedPassword, err := hasherService.HashPassword(password)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(500, "error while hashing the password", err)
	}
	d.password = hashedPassword

	return nil
}
