package services

import (
	"context"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
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
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}
	if len(users) == 0 {
		return nil, serviceErrors.NewServiceErrorWithMessage(400, "username or password not match")
	}
	user = &users[0]
	ok, err := hasherService.CompareHashAndPassword(user.GetPassword(), password)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, "error while comparing password", err)
	}
	if !ok {
		return nil, serviceErrors.NewServiceErrorWithMessage(400, "username or password not match")
	}
	return user, nil
}

func (s *userService) Register(ctx context.Context, username, email, password, confirmPassword, name, surname string, defaultRoleID uuid.UUID) (err error) {
	// Checking if the username is already being used
	users, _, err := s.userRepository.Filter(ctx, domains.UserFilter{Username: username}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}
	if len(users) != 0 {
		return serviceErrors.NewServiceErrorWithMessageAndError(400, "username already being used", nil)
	}

	// Checking if the email is already being used
	users, _, err = s.userRepository.Filter(ctx, domains.UserFilter{Email: email}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
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
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, "error while beginning transaction", err)
	}

	// Rollback the transaction in case of any error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Creating a new user model
	newUser, err := domains.NewUser("", username, email, password)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, "error while creating user", err)
	}

	// Save the new user to the database within the transaction
	authUserID, err := s.userRepository.AddTx(ctx, tx, newUser)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, "error while adding the user", err)
	}

	// Creating a new user profile
	newUserProfile, err := domains.NewUserProfile(authUserID.String(), defaultRoleID.String(), name, surname, true, 1, 0, 100)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, "error while creating the user profile", err)
	}

	// Save the new user profile to the database within the transaction
	if err := s.userProfileRepository.AddTx(ctx, tx, newUserProfile); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, "error while creating the user profile", err)
	}

	// Commit the transaction
	if err = s.transactionRepository.Commit(tx); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, "error while committing transaction", err)
	}

	return nil
}

func (s *userService) AuthWallet(ctx context.Context, publicKey, message, signature string) (user *domains.User, err error) {
	ok, err := hasherService.VerifySignature(publicKey, message, signature)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(400, "error while verifing signature", err)
	}
	if !ok {
		return nil, serviceErrors.NewServiceErrorWithMessage(400, "unable to verify the signature with the provided public key")
	}

	// Checking if the user already has an account
	users, _, err := s.userRepository.Filter(ctx, domains.UserFilter{PublicKey: publicKey}, 1, 1)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}
	if len(users) == 0 {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrUserProfileNotFound, err)
	}
	user = &users[0]

	return
}

func (s *userService) ConnectWallet(ctx context.Context, userAuthID, publicKey, message, signature string) (err error) {
	// ok, err := hasherService.VerifySignature(publicKey, message, signature)
	// if err != nil {
	// 	return serviceErrors.NewServiceErrorWithMessageAndError(400, "error while verifing signature", err)
	// }
	// if !ok {
	// 	return serviceErrors.NewServiceErrorWithMessage(400, "unable to verify the signature with the provided public key")
	// }

	userAuths, _, err := s.userRepository.Filter(ctx, domains.UserFilter{
		ID: uuid.MustParse(userAuthID),
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}
	if len(userAuths) == 0 {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusNotFound, errorDomains.ErrUserNotFound, err)
	}
	user := &userAuths[0]
	user.SetPublicKey(publicKey)

	if err := s.userRepository.Update(ctx, user); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileUpdatingUserAuth, err)
	}

	return
}
