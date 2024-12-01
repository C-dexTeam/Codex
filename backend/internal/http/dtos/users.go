package dto

import "github.com/C-dexTeam/codex/internal/http/sessionStore"

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

type UserProfileDTO struct {
	UserID   string `json:"userID"`
	RoleID   string `json:"roleID"`
	RoleName string `json:"roleName"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func (UserDTOManager) ToUserProfile(userData sessionStore.SessionData) UserProfileDTO {
	return UserProfileDTO{
		UserID:   userData.UserID,
		RoleID:   userData.RoleID,
		RoleName: userData.RoleName,
		Username: userData.Username,
		Email:    userData.Email,
		Name:     userData.Name,
		Surname:  userData.Surname,
	}
}

type UserProfileUpdateDTO struct {
	Name    string `json:"name" validate:"omitempty,max=30"`
	Surname string `json:"surname" validate:"omitempty,max=30"`
}
