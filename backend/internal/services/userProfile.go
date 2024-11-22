package services

import (
	"context"
	"strconv"

	"github.com/C-dexTeam/codex/internal/domains"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"

	"github.com/google/uuid"
)

type userProfileService struct {
	userProfileRepository domains.IUserProfileRepository
	utilService           IUtilService
}

func newUserProfileService(
	userProfileRepository domains.IUserProfileRepository,
	utils IUtilService,
) domains.IUserProfileService {
	return &userProfileService{
		userProfileRepository: userProfileRepository,
		utilService:           utils,
	}
}

func (s *userProfileService) GetAllUsersProfile(ctx context.Context, id, userID, roleID, name, surname, page, limit string) ([]domains.UserProfile, error) {
	var userProfileUUID, userAuthUUID, roleUUID uuid.UUID

	// Default Values
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = 10
	}

	if id != "" {
		userProfileUUID, err = uuid.Parse(id)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(400, "Invalid user profile id", err)
		}
	}
	if userID != "" {
		userAuthUUID, err = uuid.Parse(userID)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(400, "Invalid user id", err)
		}
	}
	if roleID != "" {
		roleUUID, err = uuid.Parse(roleID)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(400, "Invalid role id", err)
		}
	}

	users, _, err := s.userProfileRepository.Filter(ctx, domains.UserProfileFilter{
		ID:      userProfileUUID,
		UserID:  userAuthUUID,
		RoleID:  roleUUID,
		Name:    name,
		Surname: surname,
	}, int64(limitNum), int64(pageNum))
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(500, "error while filtering users profile", err)
	}

	return users, nil
}
