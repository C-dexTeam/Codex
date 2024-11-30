package private

import (
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initRewardsRoutes(root fiber.Router) {
	rewardsRoutes := root.Group("/rewards")
	rewardsRoutes.Get("/", h.GetRewards)
}

// @Tags Reward
// @Summary Get All Rewards
// @Description Retrieves all rewards based on the provided query parameters.
// @Accept json
// @Produce json
// @Param rewardID query string false "Reward ID"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/rewards/ [get]
func (h *PrivateHandler) GetRewards(c *fiber.Ctx) error {
	rewardID := c.Query("rewardID")
	page := c.Query("page")
	limit := c.Query("limit")

	rewards, err := h.services.RewardService().GetRewards(c.Context(), rewardID, page, limit)
	if err != nil {
		return err
	}
	rewardDTOs := h.dtoManager.RewardManager().ToRewardDTOs(rewards)

	return response.Response(200, "Status OK", rewardDTOs)
}
