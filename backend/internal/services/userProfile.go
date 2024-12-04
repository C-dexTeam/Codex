package services

import (
	"context"
	"strconv"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
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

func (s *userProfileService) GetUsers(
	ctx context.Context,
	id, userID, roleID, name, surname, page, limit string,
) ([]domains.UserProfile, error) {
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

func (s *userProfileService) Update(
	ctx context.Context,
	userProfileID, name, surname string,
) (err error) {
	userProfileUUID, err := uuid.Parse(userProfileID)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	userProfile, _, err := s.userProfileRepository.Filter(ctx, domains.UserProfileFilter{
		ID: userProfileUUID,
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUserPorfile, err)
	}
	if len(userProfile) == 0 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusNotFound, errorDomains.ErrUserProfileNotFound)
	}
	newProfile := userProfile[0]

	newProfile.SetName(name)
	newProfile.SetSurname(surname)
	newProfile.SetFirstLogin(true)

	if err := s.userProfileRepository.Update(ctx, &newProfile); err != nil {
		return err
	}

	return nil
}

func (s *userProfileService) ChangeUserRole(
	ctx context.Context,
	userProfileID, newRoleID string,
) (err error) {
	userProfileUUID, err := uuid.Parse(userProfileID)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	userProfile, _, err := s.userProfileRepository.Filter(ctx, domains.UserProfileFilter{
		ID: userProfileUUID,
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUserPorfile, err)
	}
	if len(userProfile) == 0 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusNotFound, errorDomains.ErrUserProfileNotFound)
	}

	newProfile := userProfile[0]
	newProfile.SetRoleID(newRoleID)
	if err := s.userProfileRepository.ChangeRole(ctx, &newProfile); err != nil {
		return err
	}

	return nil
}

func (s *userProfileService) AddUserExp(ctx context.Context,
	userProfileID string, experience int,
) (err error) {
	userProfileUUID, err := uuid.Parse(userProfileID)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	userProfile, _, err := s.userProfileRepository.Filter(ctx, domains.UserProfileFilter{
		ID: userProfileUUID,
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUserPorfile, err)
	}
	if len(userProfile) == 0 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusNotFound, errorDomains.ErrUserProfileNotFound)
	}
	profile := userProfile[0]

	totalExp := profile.GetExperience() + experience

	for totalExp >= profile.GetNextLevelExperience() {
		totalExp -= profile.GetNextLevelExperience()

		profile.SetLevel(profile.GetLevel() + 1)

		profile.SetNextLevelExperience()
	}
	profile.SetExperience(totalExp)

	if err := s.userProfileRepository.AddExp(ctx, &profile); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileAddingExperience, err)
	}

	return nil
}
