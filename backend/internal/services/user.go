package services

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
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

func (s *UserService) SetPublicKey(ctx context.Context, userID, publicKey string) error {
	// Checking if the user already has an account
	users, err := s.queries.GetUsersAuth(ctx, repo.GetUsersAuthParams{
		ID:  s.utilService.ParseNullUUID(userID),
		Lim: 1,
		Off: 0,
	})
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUsers, err)
	}
	if len(users) == 0 {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusBadRequest, serviceErrors.ErrUserNotFound, err)
	}

	if publicKey == "" {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrPublicKeyEmpty)
	}

	if err := s.queries.SetPublicKey(ctx, repo.SetPublicKeyParams{
		UserAuthID: uuid.MustParse(userID),
		PublicKey:  s.utilService.ParseString(publicKey),
	}); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileSettingPublicKey, err)
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

func (s *UserService) MintNFT(sessionID, publcKey, name, symbol, uri string, sellerFree int) (*response.BaseResponse, error) {
	data, err := s.mintRequest(sessionID, publcKey, name, symbol, uri, sellerFree)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusInternalServerError, serviceErrors.ErrWeb3RunError)
	}

	return data, nil
}

func (s *UserService) mintRequest(sessionID string, publicKeyStr, name, symbol, uri string, sellerFree int) (*response.BaseResponse, error) {
	// nginx domain because we are inside of docker & i'm going to do load balancer.
	url := "http://nginx/web3-api/nft/mint"

	// Create a mintDTO
	mintDTO := dto.MintNFTDTO{
		PublicKeyStr: publicKeyStr,
		Name:         name,
		Symbol:       symbol,
		URI:          uri,
		SellerFee:    int64(sellerFree),
	}

	// Serialize mintDTO to JSON
	requestBody, err := json.Marshal(mintDTO)
	if err != nil {
		return nil, response.Response(500, "Error marshalling mintDTO", err)
	}

	// Create a new POST request with the JSON body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, response.Response(500, "Error creating POST request", err)
	}

	// Set the Content-Type header to application/json
	req.Header.Add("Content-Type", "application/json")

	// Add the Codex-Compiler header
	req.Header.Add("Codex-Web3", hasherService.MD5Hash(s.utilService.D().Secret))

	// Add the session_id cookie to the request
	req.AddCookie(&http.Cookie{
		Name:  "session_id",
		Value: sessionID,
	})

	// Create an HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, response.Response(500, "Error making POST request", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, response.Response(500, "Error reading response body", nil)
	}

	fmt.Println(string(body))

	var data response.BaseResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(500, "Error decoding session data", err)
	}

	return &data, nil
}
