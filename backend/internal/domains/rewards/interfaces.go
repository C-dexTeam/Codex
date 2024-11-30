package rewardsDomains

import "context"

type IRewardRepository interface {
	Filter(ctx context.Context, filter RewardFilter, limit, page int64) (rewards []Reward, dataCount int64, err error)
}

type IRewardService interface{}

const (
	DefaultRewardLimit = 10
)
