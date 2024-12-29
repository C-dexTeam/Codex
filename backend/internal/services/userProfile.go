package services

import (
	"context"
	"database/sql"
	"math"
	"strconv"

	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	repo "github.com/C-dexTeam/codex/internal/repos/out"

	"github.com/google/uuid"
)

type userProfileService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func newUserProfileService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *userProfileService {
	return &userProfileService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *userProfileService) GetUsers(
	ctx context.Context,
	id, userAuthID, roleID, name, surname, page, limit string,
) ([]repo.TUsersProfile, error) {
	// Default Values
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = s.utilService.D().Limits.DefaultUserLimit
	}

	if _, err = s.utilService.ParseUUID(id); err != nil {
		return nil, err
	}
	if _, err = s.utilService.ParseUUID(userAuthID); err != nil {
		return nil, err
	}
	if _, err = s.utilService.ParseUUID(roleID); err != nil {
		return nil, err
	}

	usersProfile, err := s.queries.GetUsersProfile(ctx, repo.GetUsersProfileParams{
		ID:         s.utilService.ParseNullUUID(id),
		UserAuthID: s.utilService.ParseNullUUID(userAuthID),
		RoleID:     s.utilService.ParseNullUUID(roleID),
		Name:       s.utilService.ParseString(name),
		Surname:    s.utilService.ParseString(surname),
		Lim:        int32(limitNum),
		Off:        (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUserProfile, err)
	}

	return usersProfile, nil
}

func (s *userProfileService) Update(
	ctx context.Context,
	id, name, surname string,
) (err error) {
	if _, err := s.utilService.ParseUUID(id); err != nil {
		return err
	}

	usersProfile, err := s.queries.GetUsersProfile(ctx, repo.GetUsersProfileParams{
		ID:  uuid.NullUUID{UUID: uuid.MustParse(id), Valid: true},
		Lim: 1,
		Off: 1,
	})
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUserProfile, err)
	}
	if len(usersProfile) == 0 {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusNotFound, serviceErrors.ErrUserProfileNotFound)
	}
	newProfile := usersProfile[0]

	if name != "" {
		newProfile.Name.String = name
	}
	if surname != "" {
		newProfile.Surname.String = surname
	}

	if err := s.queries.UpdateUserProfile(ctx, repo.UpdateUserProfileParams{
		UserProfileID: newProfile.ID,
		Name:          s.utilService.ParseString(name),
		Surname:       s.utilService.ParseString(surname),
	}); err != nil {
		return err
	}

	return nil
}

func (s *userProfileService) ChangeUserRole(
	ctx context.Context,
	id, newRoleID string,
) (err error) {
	if _, err := s.utilService.ParseUUID(id); err != nil {
		return err
	}

	usersProfile, err := s.queries.GetUsersProfile(ctx, repo.GetUsersProfileParams{
		ID:  s.utilService.ParseNullUUID(id),
		Lim: 1,
		Off: 1,
	})
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUserProfile, err)
	}
	if len(usersProfile) == 0 {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusNotFound, serviceErrors.ErrUserProfileNotFound)
	}
	newProfile := usersProfile[0]
	if newRoleID != "" {
		newProfile.RoleID = uuid.MustParse(newRoleID)
	}

	if err := s.queries.ChangeUserRole(ctx, repo.ChangeUserRoleParams{
		UserProfileID: newProfile.ID,
		RoleID:        newProfile.RoleID,
	}); err != nil {
		return err
	}

	return nil
}

func (s *userProfileService) AddUserExp(ctx context.Context,
	id string, experience int,
) (err error) {
	if _, err := s.utilService.ParseUUID(id); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusBadRequest, serviceErrors.ErrInvalidID, err)
	}

	usersProfile, err := s.queries.GetUsersProfile(ctx, repo.GetUsersProfileParams{
		ID:  s.utilService.ParseNullUUID(id),
		Lim: 1,
		Off: 1,
	})
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUserProfile, err)
	}
	if len(usersProfile) == 0 {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusNotFound, serviceErrors.ErrUserProfileNotFound)
	}
	profile := usersProfile[0]

	totalExp := profile.Experience.Int32 + int32(experience)

	for totalExp >= profile.NextLevelExp.Int32 {
		totalExp -= profile.NextLevelExp.Int32

		profile.Level.Int32 = profile.Level.Int32 + 1

		profile.NextLevelExp.Int32 = int32(float64(profile.Experience.Int32) + math.Pow(float64(profile.Level.Int32), 1.2))
	}
	profile.Experience.Int32 = int32(totalExp)

	if err := s.queries.ChangeUserLevel(ctx, repo.ChangeUserLevelParams{
		UserProfileID: profile.ID,
		Level:         profile.Level,
		NextLevelExp:  profile.NextLevelExp,
		Experience:    profile.Experience,
	}); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileAddingExperience, err)
	}

	return nil
}
