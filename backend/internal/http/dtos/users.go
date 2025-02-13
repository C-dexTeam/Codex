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
	Name            string `json:"name" validate:"required,min=3,max=30"`
	Surname         string `json:"surname" validate:"required,min=3,max=60"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"ConfirmPassword" validate:"required,min=8"`
}

type UserLoginDTO struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=30"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserAuthWallet struct {
	PublicKeyBase58 string `json:"publicKeyBase58" validate:"required"`
	Message         string `json:"message" validate:"required"`
	Signature       string `json:"signatureBase58" validate:"required"`
}

type UserAuthView struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (m *UserDTOManager) ToUserAuthView(user *repo.TUsersAuth) UserAuthView {
	return UserAuthView{
		ID:       user.ID,
		Username: user.Username.String,
		Email:    user.Email.String,
		Password: "*********",
	}
}

func (m *UserDTOManager) ToUserAuthViews(users []repo.TUsersAuth) []UserAuthView {
	var userAuthDTOS []UserAuthView
	for _, user := range users {
		userAuthDTOS = append(userAuthDTOS, m.ToUserAuthView(&user))
	}
	return userAuthDTOS
}

type UserProfileView struct {
	PublicKey           string           `json:"publicKey"`
	UserID              string           `json:"userAuthID"`
	RoleName            string           `json:"role"`
	Username            string           `json:"username"`
	Email               string           `json:"email"`
	Name                string           `json:"name"`
	Surname             string           `json:"surname"`
	Level               int              `json:"level"`
	Experience          int              `json:"experience"`
	NextLevelExperience int              `json:"nextLevelExperience"`
	Statistic           *StatisticView   `json:"statistic"`
	UserRewards         []UserRewardView `json:"rewards"`
}

type StatisticView struct {
	TotalEnrolledCourses  int64 `json:"enrolledCourses"`
	CompletedCourses      int64 `json:"completedCourses"`
	TotalEnrolledChapters int64 `json:"enrolledChapters"`
	CompletedChapters     int64 `json:"compleredChapters"`
	Streak                int   `json:"streak"`
}

func (UserDTOManager) ToUserProfile(userData sessionStore.SessionData, statistic *repo.UserStatisticRow, userRewards []repo.UserRewardsRow, streak int) UserProfileView {
	rewardManager := new(RewardDTOManager)

	var stat StatisticView
	if statistic != nil {
		stat = StatisticView{
			TotalEnrolledCourses:  statistic.TotalEnrolledCourses,
			CompletedCourses:      statistic.CompletedCourses,
			TotalEnrolledChapters: statistic.TotalEnrolledChapters,
			CompletedChapters:     statistic.CompletedChapters,
			Streak:                streak,
		}
	}

	return UserProfileView{
		PublicKey:           userData.PublicKey,
		UserID:              userData.UserID,
		RoleName:            userData.Role,
		Username:            userData.Username,
		Email:               userData.Email,
		Name:                userData.Name,
		Surname:             userData.Surname,
		Level:               userData.Level,
		Experience:          userData.Experience,
		NextLevelExperience: userData.NextLevelExp,
		Statistic:           &stat,
		UserRewards:         rewardManager.ToUserRewardDTOs(userRewards),
	}
}

type UserProfileUpdateDTO struct {
	Name    string `json:"name" validate:"omitempty,max=30"`
	Surname string `json:"surname" validate:"omitempty,max=30"`
}

type MintNFTDTO struct {
	PublicKeyStr string `json:"publicKey" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Symbol       string `json:"symbol" validate:"required"`
	URI          string `json:"uri" validate:"required"`
	SellerFee    int64  `json:"sellerFee" validate:"required"`
}
