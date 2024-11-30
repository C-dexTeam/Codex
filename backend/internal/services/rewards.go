package services

import "github.com/C-dexTeam/codex/internal/domains"

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
