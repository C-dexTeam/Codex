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

func (s *rewardService) GetRewards(
	ctx context.Context,
	id, name, symbol, rewardType, page, limit string,
) (rewards []domains.Reward, err error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = domains.DefaultRewardLimit
	}

	var rewardUUID uuid.UUID
	if id != "" {
		rewardUUID, err = uuid.Parse(id)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}

	rewards, _, err = s.rewardRepository.Filter(ctx, domains.RewardFilter{
		ID:         rewardUUID,
		Name:       name,
		Symbol:     symbol,
		RewardType: id,
	}, int64(limitNum), int64(pageNum))
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringRewards, err)
	}

	return rewards, nil
}

func (s *rewardService) GetReward(
	ctx context.Context,
	id, page, limit string,
) (reward *domains.Reward, err error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = domains.DefaultAttributeLimit
	}

	var rewardUUID uuid.UUID
	rewardUUID, err = uuid.Parse(id)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	rewards, _, err := s.rewardRepository.Filter(ctx, domains.RewardFilter{
		ID: rewardUUID,
	}, 1, 1)
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringRewards, err)
	}
	if len(rewards) != 1 {
		return nil, serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusNotFound, errorDomains.ErrRewardNotFound)
	}
	reward = &rewards[0]

	rewardAttributes, _, err := s.attributeRepository.Filter(ctx, domains.AttributeFilter{
		RewardID: reward.GetID(),
	}, int64(limitNum), int64(pageNum))
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringRewardsAttributes, err)
	}
	reward.SetAttribute(rewardAttributes)

	return
}

func (s *rewardService) AddReward(
	ctx context.Context,
	rewardType, symbol, name, description, imagePath, URI string,
) (uuid.UUID, error) {
	newReward, err := domains.NewReward(
		"",
		rewardType,
		symbol,
		name,
		description,
		imagePath,
		URI,
		nil,
	)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := s.rewardRepository.Add(ctx, newReward)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s *rewardService) UpdateReward(
	ctx context.Context,
	id, rewardType, symbol, name, description, imagePath, URI string,
) error {
	var idUUID uuid.UUID
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	rewards, _, err := s.rewardRepository.Filter(ctx, domains.RewardFilter{
		ID: idUUID,
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringRewards, err)
	}
	if len(rewards) != 1 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusNotFound, errorDomains.ErrRewardNotFound)
	}

	updateReward, err := domains.NewReward(
		id,
		rewardType,
		symbol,
		name,
		description,
		imagePath,
		URI,
		nil,
	)
	if err != nil {
		return err
	}

	if err := s.rewardRepository.Update(ctx, updateReward); err != nil {
		return err
	}

	return nil
}

func (s *rewardService) DeleteReward(
	ctx context.Context,
	id string,
) (err error) {
	var idUUID uuid.UUID
	idUUID, err = uuid.Parse(id)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	rewards, _, err := s.rewardRepository.Filter(ctx, domains.RewardFilter{
		ID: idUUID,
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringRewards, err)
	}
	if len(rewards) != 1 {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrRewardNotFound, err)
	}

	if err = s.rewardRepository.Delete(ctx, idUUID); err != nil {
		return
	}
	return
}
