package services

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	hasherService "github.com/C-dexTeam/codex/pkg/hasher"

	"github.com/google/uuid"
)

type UserService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func newUserService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *UserService {
	return &UserService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *UserService) Login(ctx context.Context, username, password string) (*repo.TUsersAuth, error) {
	user, err := s.queries.GetUserAuthByUsername(ctx, s.utilService.ParseString(username))
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrUserNotFound)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}
	ok, err := hasherService.CompareHashAndPassword(user.Password.String, password)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileComparingPassword, err)
	}
	if !ok {
		return nil, serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidAuth)
	}
	return &user, nil
}

func (s *UserService) Register(ctx context.Context, username, email, password, confirmPassword, name, surname string, defaultRoleID uuid.UUID) (err error) {
	// Checking if the username is already being used
	users, err := s.queries.GetUsersAuth(ctx, repo.GetUsersAuthParams{
		Username: s.utilService.ParseString(username),
		Lim:      1,
		Off:      0,
	})
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}
	if len(users) != 0 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrUsernameBeingUsed)
	}
	// Checking if the email is already being used
	users, err = s.queries.GetUsersAuth(ctx, repo.GetUsersAuthParams{
		Email: s.utilService.ParseString(email),
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

	hashedPassword, err := hasherService.HashPassword(password)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileHashing)
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
		Username: s.utilService.ParseString(username),
		Email:    s.utilService.ParseString(email),
		Password: s.utilService.ParseString(hashedPassword),
	})
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileCreatingUserAuth, err)
	}

	if _, err := qtx.CreateUserProfile(ctx, repo.CreateUserProfileParams{
		UserAuthID: userAuthID,
		RoleID:     defaultRoleID,
		Name:       s.utilService.ParseString(name),
		Surname:    s.utilService.ParseString(surname),
	}); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileCreatingUserProfile, err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrTransactionError, err)
	}

	return nil
}

func (s *UserService) AuthWallet(ctx context.Context, publicKey, message, signature string) (*repo.TUsersAuth, error) {
	ok, err := hasherService.VerifySignature(publicKey, message, signature)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrWalletVerificationError, err)
	}
	if !ok {
		return nil, serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidWalletConnection)
	}

	// Checking if the user already has an account
	users, err := s.queries.GetUsersAuth(ctx, repo.GetUsersAuthParams{
		PublicKey: s.utilService.ParseString(publicKey),
		Lim:       1,
		Off:       0,
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

func (s *UserService) ConnectWallet(ctx context.Context, id, publicKey, message, signature string) (err error) {
	ok, err := hasherService.VerifySignature(publicKey, message, signature)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrWalletVerificationError, err)
	}
	if !ok {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrInvalidWalletConnection)
	}

	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}

	if ok, err := s.queries.CheckUserAuthByID(ctx, idUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusBadRequest, errorDomains.ErrUserNotFound)
	}

	if err := s.queries.UpdateUserAuth(ctx, repo.UpdateUserAuthParams{
		PublicKey: s.utilService.ParseString(publicKey),
	}); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileUpdatingUserAuth, err)
	}

	return
}

func (s *UserService) GetUsers(ctx context.Context, id, username, email, page, limit string) ([]repo.TUsersAuth, error) {
	// Default Values
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = 10
	}

	if _, err := s.utilService.ParseUUID(id); err != nil {
		return nil, err
	}

	users, err := s.queries.GetUsersAuth(ctx, repo.GetUsersAuthParams{
		ID:       s.utilService.ParseNullUUID(id),
		Username: s.utilService.ParseString(username),
		Email:    s.utilService.ParseString(email),
		Lim:      int32(limitNum),
		Off:      (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			errorDomains.StatusInternalServerError,
			errorDomains.ErrErrorWhileFilteringUsers,
			err,
		)
	}

	return users, nil
}
