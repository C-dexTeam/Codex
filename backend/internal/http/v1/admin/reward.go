package admin

import (
	"fmt"
	"path/filepath"

	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/pkg/paths"
	"github.com/gofiber/fiber/v2"
)

func (h *AdminHandler) initRewardsRoutes(root fiber.Router) {
	rewardsRoutes := root.Group("/rewards")
	rewardsRoutes.Post("/", h.AddReward)
	rewardsRoutes.Delete("/:id", h.DeleteReward)
	rewardsRoutes.Patch("/", h.UpdateReward)
}

// @Tags Reward
// @Summary Add Reward
// @Description Adds Reward Into DB.
// @Accept json
// @Produce json
// @Param imageFile formData file true "Reward Image File"
// @Param rewardInformation formData dto.AddRewardDTO true "New Reward"
// @Success 200 {object} response.BaseResponse{}
// @Router /admin/rewards/ [post]
func (h *AdminHandler) AddReward(c *fiber.Ctx) error {
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
		newReward.Symbol,
		newReward.Name,
		newReward.Description,
		newReward.SellerFee,
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
	fmt.Println(uri, "psaokdapğsok")

	if err := h.services.RewardService().UpdateReward(c.Context(), id.String(), "", "", "", imagePath, uri); err != nil {
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
// @Router /admin/rewards/ [patch]
func (h *AdminHandler) UpdateReward(c *fiber.Ctx) error {
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
// @Router /admin/rewards/{id} [delete]
func (h *AdminHandler) DeleteReward(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.services.RewardService().DeleteReward(c.Context(), id); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}
