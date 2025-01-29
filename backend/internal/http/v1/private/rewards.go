package private

import (
	"path/filepath"

	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/pkg/paths"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initRewardsRoutes(root fiber.Router) {
	rewardsRoutes := root.Group("/rewards")
	rewardsRoutes.Get("/", h.GetRewards)
	rewardsRoutes.Get("/:id", h.GetReward)

	rewardAdminRoutes := root.Group("/admin/rewards")
	rewardAdminRoutes.Use(h.adminRoleMiddleware)
	rewardAdminRoutes.Post("/", h.AddReward)
	rewardAdminRoutes.Delete("/:id", h.DeleteReward)
	rewardAdminRoutes.Patch("/", h.UpdateReward)
}

// @Tags Reward
// @Summary Get All Rewards
// @Description Retrieves all rewards based on the provided query parameters.
// @Accept json
// @Produce json
// @Param rewardID query string false "Reward ID"
// @Param name query string false "Reward Name"
// @Param symbol query string false "Reward Symbol"
// @Param rewardType query string false "Reward Type"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/rewards/ [get]
func (h *PrivateHandler) GetRewards(c *fiber.Ctx) error {
	rewardID := c.Query("rewardID")
	name := c.Query("name")
	symbol := c.Query("symbol")
	rewardType := c.Query("rewardType")
	page := c.Query("page")
	limit := c.Query("limit")

	rewards, err := h.services.RewardService().GetRewards(c.Context(), rewardID, name, symbol, rewardType, page, limit)
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

	reward, attribute, err := h.services.RewardService().GetReward(c.Context(), id, page, limit)
	if err != nil {
		return err
	}
	rewardDTO := h.dtoManager.RewardManager().ToRewardDTO(reward, attribute)

	return response.Response(200, "Status OK", rewardDTO)
}

// @Tags Reward
// @Summary Add Reward
// @Description Adds Reward Into DB.
// @Accept json
// @Produce json
// @Param imageFile formData file true "Reward Image File"
// @Param rewardInformation formData dto.AddRewardDTO true "New Reward"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/rewards/ [post]
func (h *PrivateHandler) AddReward(c *fiber.Ctx) error {
	var newReward dto.AddRewardDTO
	if err := c.BodyParser(&newReward); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(newReward); err != nil {
		return err
	}

	// Dosya alanını alıyoruz (sadece imageFile)
	imageFile, err := c.FormFile("imageFile")
	if err != nil {
		return err
	}

	id, err := h.services.RewardService().AddReward(
		c.Context(),
		newReward.RewardType,
		newReward.Symbol,
		newReward.Name,
		newReward.Description,
	)
	if err != nil {
		return err
	}

	// Save Image
	extention := filepath.Ext(imageFile.Filename)
	imagePath := h.services.UploadService().Web3Dir() + "/" + id.String() + extention
	if err := h.services.UploadService().SaveImage(imageFile, imagePath); err != nil {
		return err
	}
	uri := paths.CreateURI(h.config.Application.Https, id.String(), h.config.Application.Site)

	if err := h.services.RewardService().UpdateReward(c.Context(), id.String(), "", "", "", "", imagePath, uri); err != nil {
		return err
	}

	return response.Response(200, "Status OK", id)
}

// @Tags Reward
// @Summary Update Reward
// @Description Updates Reward Into DB.
// @Accept json
// @Produce json
// @Param updateReward formData dto.UpdateRewardDTO true "Update Reward"
// @Param imageFile formData file false "Course Image File"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/rewards/ [patch]
func (h *PrivateHandler) UpdateReward(c *fiber.Ctx) error {
	var updateReward dto.UpdateRewardDTO
	if err := c.BodyParser(&updateReward); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(updateReward); err != nil {
		return err
	}

	// Dosya alanını alıyoruz (sadece imageFile)
	imageFile, err := c.FormFile("imageFile")
	if err != nil {
		return err
	}

	var newImagePath string
	if imageFile != nil {
		extention := filepath.Ext(imageFile.Filename)
		newImagePath = h.services.UploadService().Web3Dir() + "/" + updateReward.ID + extention
		if err := h.services.UploadService().SaveImage(imageFile, newImagePath); err != nil {
			return err
		}
	}

	err = h.services.RewardService().UpdateReward(
		c.Context(),
		updateReward.ID,
		updateReward.RewardType,
		updateReward.Symbol,
		updateReward.Name,
		updateReward.Description,
		newImagePath,
		"",
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", nil)
}

// @Tags Reward
// @Summary Delete Reward
// @Description Delete Rewards from DB.
// @Accept json
// @Produce json
// @Param id path string false "Reward ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/rewards/{id} [delete]
func (h *PrivateHandler) DeleteReward(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.services.RewardService().DeleteReward(c.Context(), id); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}
