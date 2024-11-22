package dto

import (
	"github.com/C-dexTeam/codex/internal/domains"

	"github.com/google/uuid"
)

// AdminDTOManager handles the conversion of domain users to DTOs
type AdminDTOManager struct{}

// NewAdminDTOManager creates a new instance of AdminDTOManager
func NewAdminDTOManager() AdminDTOManager {
	return AdminDTOManager{}
}

// UserAuthDTO represents the DTO for user authentication details
// This is for get all
type UserAuthDTO struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (m *AdminDTOManager) ToUserAuthDTO(user domains.User) UserAuthDTO {
	return UserAuthDTO{
		ID:       user.GetID(),
		Username: user.GetUsername(),
		Email:    user.GetEmail(),
		Password: "*********",
	}
}

func (m *AdminDTOManager) ToUserAuthDTOs(users []domains.User) []UserAuthDTO {
	var userAuthDTOS []UserAuthDTO
	for _, user := range users {
		userAuthDTOS = append(userAuthDTOS, m.ToUserAuthDTO(user))
	}
	return userAuthDTOS
}
