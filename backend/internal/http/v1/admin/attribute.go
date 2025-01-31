package admin

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *AdminHandler) initAttributesRoutes(root fiber.Router) {
	attributeRoutes := root.Group("/attributes")

	attributeRoutes.Post("/", h.AddAttribute)
	attributeRoutes.Delete("/", h.DeleteAttribute)
	attributeRoutes.Patch("/", h.UpdateAttribute)
}

// @Tags Attributes
// @Summary Add Attribute
// @Description Adds Attribute Into DB.
// @Accept json
// @Produce json
// @Param newAttribute body dto.AddAttributeDTO true "New Attribute"
// @Success 200 {object} response.BaseResponse{}
// @Router /admin/attributes/ [post]
func (h *AdminHandler) AddAttribute(c *fiber.Ctx) error {
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
// @Router /admin/attributes/ [patch]
func (h *AdminHandler) UpdateAttribute(c *fiber.Ctx) error {
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
// @Router /admin/attributes/:id [delete]
func (h *AdminHandler) DeleteAttribute(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.services.AttributeService().DeleteAttribute(c.Context(), id); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}
