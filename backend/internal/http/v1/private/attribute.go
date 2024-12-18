package private

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initAttributesRoutes(root fiber.Router) {
	attributeRoutes := root.Group("/attributes")
	attributeRoutes.Get("/", h.GetAttributes)

	attributeAdminRoutes := root.Group("/admin/attributes")
	attributeAdminRoutes.Use(h.adminRoleMiddleware)
	attributeAdminRoutes.Post("/", h.AddAttribute)
	attributeAdminRoutes.Delete("/", h.DeleteAttribute)
	attributeAdminRoutes.Patch("/", h.UpdateAttribute)
}

// @Tags Attributes
// @Summary Get All Attributes
// @Description Retrieves all attribute based on the provided query parameters.
// @Accept json
// @Produce json
// @Param id query string false "Attribute ID"
// @Param rewardID query string false "Reward ID"
// @Param traitType query string false "TraitType of Attribute"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/attributes/ [get]
func (h *PrivateHandler) GetAttributes(c *fiber.Ctx) error {
	id := c.Query("id")
	rewardID := c.Query("rewardID")
	traitType := c.Query("traitType")
	page := c.Query("page")
	limit := c.Query("limit")

	attributes, err := h.services.AttributeService().GetAttributes(c.Context(), id, rewardID, traitType, page, limit)
	if err != nil {
		return err
	}
	// attributeDTOs := h.dtoManager.RewardManager().ToAttributeDTOs(attributes)

	return response.Response(200, "Status OK", attributes)
}

// @Tags Attributes
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

// @Tags Attributes
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

// @Tags Attributes
// @Summary Delete Attribute
// @Description Delete Attributes from DB.
// @Accept json
// @Produce json
// @Param id path string false "Attribute ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/attributes/:id [delete]
func (h *PrivateHandler) DeleteAttribute(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.services.AttributeService().DeleteAttribute(c.Context(), id); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}
