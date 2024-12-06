package public

import (
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PublicHandler) initRewardsRoutes(root fiber.Router) {
	rewardsRoutes := root.Group("/rewards")
	rewardsRoutes.Get("/metadata/:id", h.GetReward)
}

// @Tags Metadata
// @Summary Get Reward By ID
// @Description Retrieves one reward.
// @Accept json
// @Produce json
// @Param id path string false "Reward ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /public/rewards/metadata/{id} [get]
func (h *PublicHandler) GetReward(c *fiber.Ctx) error {
	id := c.Params("id")

	reward, err := h.services.RewardService().GetReward(c.Context(), id, "1", "50")
	if err != nil {
		return err
	}
	rewardDTO := h.dtoManager.RewardManager().ToRewardDTO(reward)

	return response.Response(200, "Status OK", rewardDTO)
}
