package services

import (
	"context"
	"fmt"

	"github.com/C-dexTeam/codex/internal/domains"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	hasherService "github.com/C-dexTeam/codex/pkg/hasher"

	"github.com/google/uuid"
)

type userService struct {
	userRepository        domains.IUserRepository
	userProfileRepository domains.IUserProfileRepository
	transactionRepository domains.ITransactionRepository
	utilService           IUtilService
}

func newUserService(
	userRepository domains.IUserRepository,
	userProfileRepository domains.IUserProfileRepository,
	transactionRepository domains.ITransactionRepository,
	utils IUtilService,
) domains.IUserService {
	return &userService{
		userRepository:        userRepository,
		userProfileRepository: userProfileRepository,
		transactionRepository: transactionRepository,
		utilService:           utils,
	}
}

func (s *userService) Login(ctx context.Context, username, password string) (user *domains.User, err error) {
	users, _, err := s.userRepository.Filter(ctx, domains.UserFilter{
		Username: username,
	}, 1, 1)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}
	if len(users) == 0 {
		return nil, serviceErrors.NewServiceErrorWithMessage(400, "username or password not match")
	}
	user = &users[0]
	ok, err := hasherService.CompareHashAndPassword(user.GetPassword(), password)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(500, "error while comparing password", err)
	}
	if !ok {
		return nil, serviceErrors.NewServiceErrorWithMessage(400, "username or password not match")
	}
	return user, nil
}

func (s *userService) Register(ctx context.Context, username, email, password, confirmPassword string, defaultRoleID uuid.UUID) (err error) {
	// Checking if the username is already being used
	users, _, err := s.userRepository.Filter(ctx, domains.UserFilter{Username: username}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}
	if len(users) != 0 {
		return serviceErrors.NewServiceErrorWithMessageAndError(400, "username already being used", nil)
	}

	// Checking if the email is already being used
	users, _, err = s.userRepository.Filter(ctx, domains.UserFilter{Email: email}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(500, "error while filtering users", err)
	}
	if len(users) != 0 {
		return serviceErrors.NewServiceErrorWithMessageAndError(400, "email already being used", nil)
	}

	// Checking if the password and confirmPassword are the same
	if password != confirmPassword {
		return serviceErrors.NewServiceErrorWithMessage(400, "passwords do not match")
	}

	// Begin a new transaction
	tx, err := s.transactionRepository.Begin()
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(500, "error while beginning transaction", err)
	}

	// Rollback the transaction in case of any error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Creating a new user model
	newUser, err := domains.NewUser(username, email, password)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(500, "error while creating user", err)
	}

	// Save the new user to the database within the transaction
	authUserID, err := s.userRepository.AddTx(ctx, tx, newUser)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(500, "error while adding the user", err)
	}

	// Creating a new user profile
	newUserProfile, err := domains.NewUserProfile(authUserID.String(), defaultRoleID.String(), "", "")
	if err != nil {
		fmt.Println(1, err)
		return serviceErrors.NewServiceErrorWithMessageAndError(500, "error while creating the user profile", err)
	}

	// Save the new user profile to the database within the transaction
	if err := s.userProfileRepository.AddTx(ctx, tx, newUserProfile); err != nil {
		fmt.Println(2, err)

		return serviceErrors.NewServiceErrorWithMessageAndError(500, "error while creating the user profile", err)
	}

	// Commit the transaction
	if err = s.transactionRepository.Commit(tx); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(500, "error while committing transaction", err)
	}

	return nil
}
