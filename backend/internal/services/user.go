package services

import (
	"context"
	"database/sql"

	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	hasherService "github.com/C-dexTeam/codex/pkg/hasher"

	"github.com/google/uuid"
)

type userService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func newUserService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *userService {
	return &userService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *userService) Login(ctx context.Context, username, password string) (*repo.TUsersAuth, error) {
	users, err := s.queries.GetUsersAuth(ctx, repo.GetUsersAuthParams{
		Username: sql.NullString{String: username},
		Lim:      1,
		Off:      1,
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}
	if len(users) == 0 {
		return nil, serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrUserNotFound)
	}
	user := &users[0]
	ok, err := hasherService.CompareHashAndPassword(user.Password.String, password)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileComparingPassword, err)
	}
	if !ok {
		return nil, serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidAuth)
	}
	return user, nil
}

func (s *userService) Register(ctx context.Context, username, email, password, confirmPassword, name, surname string, defaultRoleID uuid.UUID) (err error) {
	// Checking if the username is already being used
	users, err := s.queries.GetUsersAuth(ctx, repo.GetUsersAuthParams{
		Username: sql.NullString{String: username},
		Lim:      1,
		Off:      1,
	})
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}
	if len(users) != 0 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrUsernameBeingUsed)
	}
	// Checking if the email is already being used
	users, err = s.queries.GetUsersAuth(ctx, repo.GetUsersAuthParams{
		Email: sql.NullString{String: email},
		Lim:   1,
		Off:   1,
	})
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}
	if len(users) != 0 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrEmailBeingUsed)
	}

	// Checking if the password and confirmPassword are the same
	if password != confirmPassword {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrPasswordsDoNotMatch)
	}

	// Begin a new transaction
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	qtx := s.queries.WithTx(tx)

	// Save the new user to the database within the transaction
	userAuthID, err := qtx.CreateUserAuth(ctx, repo.CreateUserAuthParams{
		Username: sql.NullString{String: username},
		Email:    sql.NullString{String: email},
		Password: sql.NullString{String: password},
	})
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileCreatingUserAuth, err)
	}

	if err := qtx.CreateUserProfile(ctx, repo.CreateUserProfileParams{
		UserAuthID: userAuthID,
		RoleID:     defaultRoleID,
		Name:       sql.NullString{String: name},
		Surname:    sql.NullString{String: surname},
	}); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileCreatingUserProfile, err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrTransactionError, err)
	}

	return nil
}

func (s *userService) AuthWallet(ctx context.Context, publicKey, message, signature string) (*repo.TUsersAuth, error) {
	ok, err := hasherService.VerifySignature(publicKey, message, signature)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrWalletVerificationError, err)
	}
	if !ok {
		return nil, serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidWalletConnection)
	}

	// Checking if the user already has an account
	users, err := s.queries.GetUsersAuth(ctx, repo.GetUsersAuthParams{
		PublicKey: sql.NullString{String: publicKey},
		Lim:       1,
		Off:       1,
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}
	if len(users) == 0 {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrUserNotFound, err)
	}
	user := &users[0]

	return user, err
}

func (s *userService) ConnectWallet(ctx context.Context, userAuthID, publicKey, message, signature string) (err error) {
	ok, err := hasherService.VerifySignature(publicKey, message, signature)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrWalletVerificationError, err)
	}
	if !ok {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidWalletConnection)
	}

	userAuths, err := s.queries.GetUsersAuth(ctx, repo.GetUsersAuthParams{
		ID:  sql.NullString{String: userAuthID},
		Lim: 1,
		Off: 1,
	})
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}
	if len(userAuths) == 0 {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrUserNotFound, err)
	}
	user := &userAuths[0]
	user.PublicKey = sql.NullString{String: publicKey}

	if err := s.queries.UpdateUserAuth(ctx, repo.UpdateUserAuthParams{
		PublicKey: sql.NullString{String: publicKey},
	}); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileUpdatingUserAuth, err)
	}

	return
}
