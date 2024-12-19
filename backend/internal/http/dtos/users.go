package dto

import (
	"github.com/C-dexTeam/codex/internal/http/sessionStore"
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

// UserDTOManager handles the conversion of domain users to DTOs
type UserDTOManager struct{}

// NewUserDTOManager creates a new instance of UserDTOManager
func NewUserDTOManager() UserDTOManager {
	return UserDTOManager{}
}

type LoginResponseDTO struct {
	Role string `json:"role"`
}

func (m *UserDTOManager) ToLoginResponseDTO(role string) LoginResponseDTO {
	return LoginResponseDTO{
		Role: role,
	}
}

// UserRegisterDTO
type UserRegisterDTO struct {
	Username        string `json:"username" validate:"required,alphanum,min=3,max=30"`
	Email           string `json:"email" validate:"required,email"`
	Name            string `json:"name" validate:"required"`
	Surname         string `json:"surname" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"ConfirmPassword" validate:"required,min=8"`
}

type UserLoginDTO struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=30"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserAuthWallet struct {
	PublicKeyBase58 string `json:"publicKeyBase58"`
	Message         string `json:"message"`
	Signature       string `json:"signatureBase58"`
}
type UserAuthDTO struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (m *UserDTOManager) ToUserAuthDTO(user *repo.TUsersAuth) UserAuthDTO {
	return UserAuthDTO{
		ID:       user.ID,
		Username: user.Username.String,
		Email:    user.Email.String,
		Password: "*********",
	}
}

func (m *UserDTOManager) ToUserAuthDTOs(users []repo.TUsersAuth) []UserAuthDTO {
	var userAuthDTOS []UserAuthDTO
	for _, user := range users {
		userAuthDTOS = append(userAuthDTOS, m.ToUserAuthDTO(&user))
	}
	return userAuthDTOS
}

type UserProfileDTO struct {
	PublicKey           string `json:"publicKey"`
	RoleName            string `json:"role"`
	Username            string `json:"username"`
	Email               string `json:"email"`
	Name                string `json:"name"`
	Surname             string `json:"surname"`
	Level               int    `json:"level"`
	Experience          int    `json:"experience"`
	NextLevelExperience int    `json:"nextLevelExperience"`
}

func (UserDTOManager) ToUserProfile(userData sessionStore.SessionData) UserProfileDTO {
	return UserProfileDTO{
		PublicKey:           userData.PublicKey,
		RoleName:            userData.Role,
		Username:            userData.Username,
		Email:               userData.Email,
		Name:                userData.Name,
		Surname:             userData.Surname,
		Level:               userData.Level,
		Experience:          userData.Experience,
		NextLevelExperience: userData.NextLevelExp,
	}
}

type UserProfileUpdateDTO struct {
	Name    string `json:"name" validate:"omitempty,max=30"`
	Surname string `json:"surname" validate:"omitempty,max=30"`
}
