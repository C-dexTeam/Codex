package services

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

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
			return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrUserNotFound)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUsers, err)
	}
	ok, err := hasherService.CompareHashAndPassword(user.Password.String, password)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileComparingPassword, err)
	}
	if !ok {
		return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrInvalidAuth)
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
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUsers, err)
	}
	if len(users) != 0 {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrUsernameBeingUsed)
	}
	// Checking if the email is already being used
	users, err = s.queries.GetUsersAuth(ctx, repo.GetUsersAuthParams{
		Email: s.utilService.ParseString(email),
		Lim:   1,
		Off:   1,
	})
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUsers, err)
	}
	if len(users) != 0 {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrEmailBeingUsed)
	}

	// Checking if the password and confirmPassword are the same
	if password != confirmPassword {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrPasswordsDoNotMatch)
	}

	hashedPassword, err := hasherService.HashPassword(password)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileHashing)
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
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileCreatingUserAuth, err)
	}

	if _, err := qtx.CreateUserProfile(ctx, repo.CreateUserProfileParams{
		UserAuthID: userAuthID,
		RoleID:     defaultRoleID,
		Name:       s.utilService.ParseString(name),
		Surname:    s.utilService.ParseString(surname),
	}); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileCreatingUserProfile, err)
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrTransactionError, err)
	}

	return nil
}

func (s *UserService) AuthWallet(ctx context.Context, publicKey, message, signature string) (*repo.TUsersAuth, error) {
	ok, err := hasherService.VerifySignature(publicKey, message, signature)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrWalletVerificationError, err)
	}
	if !ok {
		return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrInvalidWalletConnection)
	}

	// Checking if the user already has an account
	users, err := s.queries.GetUsersAuth(ctx, repo.GetUsersAuthParams{
		PublicKey: s.utilService.ParseString(publicKey),
		Lim:       1,
		Off:       0,
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUsers, err)
	}
	if len(users) == 0 {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusBadRequest, serviceErrors.ErrUserNotFound, err)
	}
	user := &users[0]

	return user, err
}

func (s *UserService) ConnectWallet(ctx context.Context, id, publicKey, message, signature string) (err error) {
	ok, err := hasherService.VerifySignature(publicKey, message, signature)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrWalletVerificationError, err)
	}
	if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrInvalidWalletConnection)
	}

	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}

	if ok, err := s.queries.CheckUserAuthByID(ctx, idUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUsers, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrUserNotFound)
	}

	if err := s.queries.UpdateUserAuth(ctx, repo.UpdateUserAuthParams{
		PublicKey: s.utilService.ParseString(publicKey),
	}); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileUpdatingUserAuth, err)
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
		limitNum = s.utilService.D().Limits.DefaultUserLimit
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
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringUsers,
			err,
		)
	}

	return users, nil
}
