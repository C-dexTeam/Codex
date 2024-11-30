package services

import (
	"context"
	"strconv"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type rewardService struct {
	rewardRepository    domains.IRewardRepository
	attributeRepository domains.IAttributeRepository
}

func newRewardService(
	rewardRepository domains.IRewardRepository,
	attributeRepository domains.IAttributeRepository,
) domains.IRewardService {
	return &rewardService{
		rewardRepository:    rewardRepository,
		attributeRepository: attributeRepository,
	}
}

func (s *rewardService) GetRewards(ctx context.Context, rewardID, page, limit string) (rewards []domains.Reward, err error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = domains.DefaultRewardLimit
	}

	var rewardUUID uuid.UUID = uuid.Nil
	if rewardID != "" {
		rewardUUID, err = uuid.Parse(rewardID)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}

	rewards, _, err = s.rewardRepository.Filter(ctx, domains.RewardFilter{
		ID: rewardUUID,
	}, int64(limitNum), int64(pageNum))
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringRewards, err)
	}

	return rewards, nil
}

// Eğer get by id yaparsa spesifik olarak o zaman attribute gitsin.
