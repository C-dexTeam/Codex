package services

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
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
) ([]repo.TReward, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = domains.DefaultRewardLimit
	}

	rewardUUID, err := s.utilService.ParseUUID(id)
	if err != nil {
		return nil, err
	}

	fmt.Println(s.utilService.ParseNullUUID(rewardUUID.String()))
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
			errorDomains.StatusInternalServerError,
			errorDomains.ErrErrorWhileFilteringRewards,
			err,
		)
	}

	return rewards, nil
}

func (s *rewardService) GetReward(
	ctx context.Context,
	id, page, limit string,
) (*repo.TReward, []repo.TAttribute, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = domains.DefaultAttributeLimit
	}

	rewardUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return nil, nil, err
	}

	reward, err := s.queries.GetReward(ctx, rewardUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, nil, serviceErrors.NewServiceErrorWithMessage(
				errorDomains.StatusBadRequest,
				errorDomains.ErrRewardNotFound,
			)
		}
		return nil, nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringUsers, err)
	}

	rewardAttbitures, err := s.queries.GetAttributes(ctx, repo.GetAttributesParams{
		RewardID: s.utilService.ParseNullUUID(id),
		Lim:      int32(limitNum),
		Off:      (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringRewardsAttributes, err)
	}

	return &reward, rewardAttbitures, nil
}

func (s *rewardService) AddReward(
	ctx context.Context,
	rewardType, symbol, name, description, imagePath, URI string,
) (uuid.UUID, error) {
	// TODO: Gelen verilerin kontrolü lazım. Len > 30 gibi normalde domainde yapıyorsun
	id, err := s.queries.CreateReward(ctx, repo.CreateRewardParams{
		RewardType:  rewardType,
		Symbol:      symbol,
		Name:        name,
		Description: description,
		ImagePath:   imagePath,
		Uri:         URI,
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
				errorDomains.StatusBadRequest,
				errorDomains.ErrRewardNotFound,
			)
		}
		return serviceErrors.NewServiceErrorWithMessageAndError(
			errorDomains.StatusInternalServerError,
			errorDomains.ErrErrorWhileFilteringRewards,
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
	idUUID, err := s.utilService.ParseUUID(id)
	if err != nil {
		return err
	}

	_, err = s.queries.GetReward(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return serviceErrors.NewServiceErrorWithMessage(
				errorDomains.StatusBadRequest,
				errorDomains.ErrRewardNotFound,
			)
		}
		return serviceErrors.NewServiceErrorWithMessageAndError(
			errorDomains.StatusInternalServerError,
			errorDomains.ErrErrorWhileFilteringRewards,
			err,
		)
	}

	if err := s.queries.DeleteReward(ctx, idUUID); err != nil {
		return err
	}
	return
}
