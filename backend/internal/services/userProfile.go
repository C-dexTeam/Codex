package services

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/C-dexTeam/codex/internal/domains"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	repo "github.com/C-dexTeam/codex/internal/repos/out"

	"github.com/google/uuid"
)

type userProfileService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
	mu          sync.Mutex
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
		mu:          sync.Mutex{},
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

func (s *userProfileService) GetUser(
	ctx context.Context,
	id string,
) (*repo.TUsersProfile, error) {
	// Default Values

	idUUID, err := s.utilService.ParseUUID(id)
	if err != nil {
		return nil, err
	}

	usersProfile, err := s.queries.GetUserProfile(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrUserProfileNotFound)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUserProfile, err)
	}

	return &usersProfile, nil
}

func (s *userProfileService) Update(
	ctx context.Context,
	id, name, surname string,
) (err error) {
	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}

	// Check if its exists
	_, err = s.queries.GetUserProfile(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrUserProfileNotFound)
		}
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUserProfile, err)
	}

	if err := s.queries.UpdateUserProfile(ctx, repo.UpdateUserProfileParams{
		UserProfileID: idUUID,
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
) (*repo.ChangeUserLevelRow, error) {
	idUUID, err := s.utilService.ParseUUID(id)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusBadRequest, serviceErrors.ErrInvalidID, err)
	}

	profile, err := s.queries.GetUserProfile(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrUserProfileNotFound)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUserProfile, err)
	}

	totalExp := profile.Experience.Int32 + int32(experience)

	for totalExp >= profile.NextLevelExp.Int32 {
		totalExp -= profile.NextLevelExp.Int32

		profile.Level.Int32 = profile.Level.Int32 + 1

		profile.NextLevelExp.Int32 = profile.NextLevelExp.Int32 + int32((float64(profile.NextLevelExp.Int32) * domains.ExpConstant))
	}
	profile.Experience.Int32 = int32(totalExp)

	outs, err := s.queries.ChangeUserLevel(ctx, repo.ChangeUserLevelParams{
		UserProfileID: profile.ID,
		Level:         profile.Level,
		NextLevelExp:  profile.NextLevelExp,
		Experience:    profile.Experience,
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileAddingExperience, err)
	}

	return &outs, nil
}

func (s *userProfileService) StreakUp(ctx context.Context, id string, lastStreakDate time.Time) (int32, error) {
	userProfileUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return 0, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if isSameDay(lastStreakDate, time.Now()) {
		return 0, serviceErrors.NewServiceErrorWithMessage(
			serviceErrors.StatusBadRequest,
			serviceErrors.ErrStreakAlreadyIncreased,
		)
	}

	streak, err := s.queries.StreakUp(ctx, userProfileUUID)
	if err != nil {
		return 0, err
	}

	return streak.Int32, nil
}

func (s *userProfileService) UserStatistic(ctx context.Context, userAuthID string) (*repo.UserStatisticRow, error) {
	userAuthUUID, err := s.utilService.NParseUUID(userAuthID)
	if err != nil {
		return nil, err
	}

	statistic, err := s.queries.UserStatistic(ctx, userAuthUUID)
	if err != nil {
		return nil, err
	}

	return &statistic, nil
}

func isSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
