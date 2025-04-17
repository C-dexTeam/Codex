package private

import (
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initRewardsRoutes(root fiber.Router) {
	rewardsRoutes := root.Group("/rewards")
	rewardsRoutes.Get("/", h.GetRewards)
	rewardsRoutes.Get("/:id", h.GetReward)
}

// @Tags Reward
// @Summary Get All Rewards
// @Description Retrieves all rewards based on the provided query parameters.
// @Accept json
// @Produce json
// @Param rewardID query string false "Reward ID"
// @Param name query string false "Reward Name"
// @Param symbol query string false "Reward Symbol"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/rewards/ [get]
func (h *PrivateHandler) GetRewards(c *fiber.Ctx) error {
	rewardID := c.Query("rewardID")
	name := c.Query("name")
	symbol := c.Query("symbol")
	page := c.Query("page")
	limit := c.Query("limit")

	rewards, err := h.services.RewardService().GetRewards(c.Context(), rewardID, name, symbol, page, limit)
	if err != nil {
		return err
	}
	rewardDTOs := h.dtoManager.RewardManager().ToRewardDTOs(rewards)

	return response.Response(200, "Status OK", rewardDTOs)
}

// @Tags Reward
// @Summary Get Reward By ID
// @Description Retrieves one reward.
// @Accept json
// @Produce json
// @Param id path string false "Reward ID"
// @Param page query string false "Reward Attribute Page"
// @Param limit query string false "Reward Attribute Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/rewards/{id} [get]
func (h *PrivateHandler) GetReward(c *fiber.Ctx) error {
	id := c.Params("id")
	page := c.Query("page")
	limit := c.Query("limit")

	reward, err := h.services.RewardService().GetReward(c.Context(), id, page, limit)
	if err != nil {
		return err
	}
	rewardDTO := h.dtoManager.RewardManager().ToRewardDTO(reward)

	return response.Response(200, "Status OK", rewardDTO)
}
