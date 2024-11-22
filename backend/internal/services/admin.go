package services

import (
	"context"
	"strconv"

	"github.com/C-dexTeam/codex/internal/domains"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"

	"github.com/google/uuid"
)

type adminService struct {
	userRepository        domains.IUserRepository
	userProfileRepository domains.IUserProfileRepository
	transactionRepository domains.ITransactionRepository
	utilService           IUtilService
}

func newAdminService(
	userRepository domains.IUserRepository,
	userProfileRepository domains.IUserProfileRepository,
	transactionRepository domains.ITransactionRepository,
	utils IUtilService,
) domains.IAdminService {
	return &adminService{
		userRepository:        userRepository,
		userProfileRepository: userProfileRepository,
		transactionRepository: transactionRepository,
		utilService:           utils,
	}
}

func (s *adminService) GetAllUsers(ctx context.Context, id, username, email, page, limit string) ([]domains.User, error) {
	// Default Values
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = 10
	}

	var userUUID uuid.UUID
	if id != "" {
		userUUID, err = uuid.Parse(id)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(400, "Invalid user id", err)
		}
	}

	users, _, err := s.userRepository.Filter(ctx, domains.UserFilter{
		ID:       userUUID,
		Username: username,
		Email:    email,
	}, int64(limitNum), int64(pageNum))
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}

	return users, nil
}
