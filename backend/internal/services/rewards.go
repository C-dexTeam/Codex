package services

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/C-dexTeam/codex/internal/domains"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type rewardService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func newRewardService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *rewardService {
	return &rewardService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *rewardService) GetRewards(
	ctx context.Context,
	id, name, symbol, rewardType, page, limit string,
) ([]domains.Reward, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = s.utilService.D().Limits.DefaultRewardLimit
	}

	if _, err := s.utilService.ParseUUID(id); err != nil {
		return nil, err
	}

	rewards, err := s.queries.GetRewards(ctx, repo.GetRewardsParams{
		ID:         s.utilService.ParseNullUUID(id),
		Name:       s.utilService.ParseString(name),
		Symbol:     s.utilService.ParseString(symbol),
		RewardType: s.utilService.ParseString(rewardType),
		Lim:        int32(limitNum),
		Off:        (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringRewards,
			err,
		)
	}

	rewardDomains := domains.NewRewards(rewards, nil)

	return rewardDomains, nil
}

func (s *rewardService) GetReward(
	ctx context.Context,
	id, page, limit string,
) (*domains.Reward, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = s.utilService.D().Limits.DefaultRewardLimit
	}

	rewardUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return nil, err
	}

	reward, err := s.queries.GetReward(ctx, rewardUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(
				serviceErrors.StatusBadRequest,
				serviceErrors.ErrRewardNotFound,
			)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUsers, err)
	}

	rewardAttbitures, err := s.queries.GetAttributes(ctx, repo.GetAttributesParams{
		RewardID: s.utilService.ParseNullUUID(id),
		Lim:      int32(limitNum),
		Off:      (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringRewardsAttributes, err)
	}

	domainsReward := domains.NewReward(&reward, rewardAttbitures)

	return domainsReward, nil
}

func (s *rewardService) AddReward(
	ctx context.Context,
	rewardType, symbol, name, description string,
	sellerFee int,
) (uuid.UUID, error) {
	id, err := s.queries.CreateReward(ctx, repo.CreateRewardParams{
		RewardType:  rewardType,
		Symbol:      symbol,
		Name:        name,
		Description: description,
		SellerFee:   int32(sellerFee),
	})
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s *rewardService) UpdateReward(
	ctx context.Context,
	id, rewardType, symbol, name, description, imagePath, URI string,
) error {
	idUUID, err := s.utilService.ParseUUID(id)
	if err != nil {
		return err
	}

	_, err = s.queries.GetReward(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return serviceErrors.NewServiceErrorWithMessage(
				serviceErrors.StatusBadRequest,
				serviceErrors.ErrRewardNotFound,
			)
		}
		return serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringRewards,
			err,
		)
	}

	if err := s.queries.UpdateReward(ctx, repo.UpdateRewardParams{
		RewardID:    idUUID,
		RewardType:  s.utilService.ParseString(rewardType),
		Symbol:      s.utilService.ParseString(symbol),
		Name:        s.utilService.ParseString(name),
		Description: s.utilService.ParseString(description),
		ImagePath:   s.utilService.ParseString(imagePath),
		Uri:         s.utilService.ParseString(URI),
	}); err != nil {
		return err
	}

	return nil
}

func (s *rewardService) DeleteReward(
	ctx context.Context,
	id string,
) (err error) {
	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}

	_, err = s.queries.GetReward(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return serviceErrors.NewServiceErrorWithMessage(
				serviceErrors.StatusBadRequest,
				serviceErrors.ErrRewardNotFound,
			)
		}
		return serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringRewards,
			err,
		)
	}

	if err := s.queries.DeleteReward(ctx, idUUID); err != nil {
		return err
	}
	return
}

func (s *rewardService) AddRewardIntoUser(
	ctx context.Context,
	userAuthID, chapterID, courseID, rewardID string,
) error {
	userAuthUUID := uuid.MustParse(userAuthID)
	courseUUID := uuid.MustParse(courseID)
	if _, err := s.utilService.ParseUUID(chapterID); err != nil {
		return err
	}

	rewardUUID, err := s.utilService.NParseUUID(rewardID)
	if err != nil {
		return err
	}

	if ok, err := s.queries.CheckRewardByID(ctx, rewardUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringRewards, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrRewardNotFound)
	}

	// If the reward already added. Than dont add.
	ok, err := s.queries.CheckUserReward(ctx, repo.CheckUserRewardParams{
		UserAuthID: userAuthUUID,
		CourseID:   courseUUID,
		ChapterID:  s.utilService.ParseNullUUID(chapterID),
		RewardID:   rewardUUID,
	})
	if err != nil {
		return err
	}
	if ok {
		return nil
	}

	if err := s.queries.AddRewardToUser(ctx, repo.AddRewardToUserParams{
		UserAuthID: userAuthUUID,
		ChapterID:  s.utilService.ParseNullUUID(chapterID),
		CourseID:   courseUUID,
		RewardID:   rewardUUID,
	}); err != nil {
		return err
	}

	return nil
}

func (s *rewardService) GetUserRewards(
	ctx context.Context,
	userAuthID, page, limit string,
) ([]repo.UserRewardsRow, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = s.utilService.D().Limits.DefaultRewardLimit
	}

	userAuthUUID, err := s.utilService.NParseUUID(userAuthID)
	if err != nil {
		return nil, err
	}

	userRewards, err := s.queries.UserRewards(ctx, repo.UserRewardsParams{
		UserAuthID: userAuthUUID,
		Lim:        int32(limitNum),
		Off:        (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, err
	}

	return userRewards, nil
}
