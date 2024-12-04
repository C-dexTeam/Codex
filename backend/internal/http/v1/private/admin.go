package private

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"

	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initAdminRoutes(root fiber.Router) {
	adminRoute := root.Group("/admin")
	adminRoute.Use(h.adminRoleMiddleware)
	adminRoute.Get("/user", h.GetUsers)

	rewardsRoute := adminRoute.Group("/rewards")
	rewardsRoute.Post("/", h.AddReward)
	rewardsRoute.Delete("/", h.DeleteReward)
	rewardsRoute.Patch("/", h.UpdateReward)

	attributeRoute := adminRoute.Group("/attributes")
	attributeRoute.Post("/", h.AddAttribute)
	attributeRoute.Delete("/", h.DeleteAttribute)
	attributeRoute.Patch("/", h.UpdateAttribute)
}

// @Tags Admin
// @Summary Get All Users
// @Description Retrieves all logs based on the provided query parameters.
// @Accept json
// @Produce json
// @Param id query string false "User ID"
// @Param username query string false "Username"
// @Param email query string false "User's Email"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/user [get]
func (h *PrivateHandler) GetUsers(c *fiber.Ctx) error {
	id := c.Query("id")
	username := c.Query("username")
	email := c.Query("email")
	page := c.Query("page")
	limit := c.Query("limit")

	users, err := h.services.AdminService().GetUsers(c.Context(), id, username, email, page, limit)
	if err != nil {
		return err
	}
	userAuthDTOs := h.dtoManager.AdminManager().ToUserAuthDTOs(users)

	return response.Response(200, "Status OK", userAuthDTOs)
}

// @Tags Reward
// @Summary Add Reward
// @Description Adds Reward Into DB.
// @Accept json
// @Produce json
// @Param newReward body dto.AddRewardDTO true "New Reward"
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

	id, err := h.services.RewardService().AddReward(
		c.Context(),
		newReward.RewardType,
		newReward.Symbol,
		newReward.Name,
		newReward.Description,
		newReward.ImagePath,
		newReward.URI,
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", id)
}

// @Tags Reward
// @Summary Update Reward
// @Description Updates Reward Into DB.
// @Accept json
// @Produce json
// @Param updateReward body dto.UpdateRewardDTO true "Update Reward"
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

	err := h.services.RewardService().UpdateReward(
		c.Context(),
		updateReward.ID,
		updateReward.RewardType,
		updateReward.Symbol,
		updateReward.Name,
		updateReward.Description,
		updateReward.ImagePath,
		updateReward.URI,
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
// @Param rewardID query string false "Reward ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/rewards/ [delete]
func (h *PrivateHandler) DeleteReward(c *fiber.Ctx) error {
	rewardID := c.Query("rewardID")

	if err := h.services.RewardService().DeleteReward(c.Context(), rewardID); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}

// @Tags Attribute
// @Summary Add Attribute
// @Description Adds Attribute Into DB.
// @Accept json
// @Produce json
// @Param newAttribute body dto.AddAttributeDTO true "New Attribute"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/attributes/ [post]
func (h *PrivateHandler) AddAttribute(c *fiber.Ctx) error {
	var newAttribute dto.AddAttributeDTO
	if err := c.BodyParser(&newAttribute); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(newAttribute); err != nil {
		return err
	}

	id, err := h.services.AttributeService().AddAttribute(
		c.Context(),
		newAttribute.RewardID,
		newAttribute.TraitType,
		newAttribute.Value,
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", id)
}

// @Tags Attribute
// @Summary Update Attribute
// @Description Updates Attribute Into DB.
// @Accept json
// @Produce json
// @Param updateAttribute body dto.UpdateAttributeDTO true "Update Attribute"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/attributes/ [patch]
func (h *PrivateHandler) UpdateAttribute(c *fiber.Ctx) error {
	var updateAttribute dto.UpdateAttributeDTO
	if err := c.BodyParser(&updateAttribute); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(updateAttribute); err != nil {
		return err
	}

	err := h.services.AttributeService().UpdateAttribute(
		c.Context(),
		updateAttribute.ID,
		updateAttribute.RewardID,
		updateAttribute.TraitType,
		updateAttribute.Value,
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", nil)
}

// @Tags Attribute
// @Summary Delete Attribute
// @Description Delete Attributes from DB.
// @Accept json
// @Produce json
// @Param attributeID query string false "Attribute ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/attributes/ [delete]
func (h *PrivateHandler) DeleteAttribute(c *fiber.Ctx) error {
	attributeID := c.Query("attributeID")

	if err := h.services.AttributeService().DeleteAttribute(c.Context(), attributeID); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}
