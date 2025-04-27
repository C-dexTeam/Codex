package private

import (
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initAttributesRoutes(root fiber.Router) {
	attributeRoutes := root.Group("/attributes")
	attributeRoutes.Get("/", h.GetAttributes)
	attributeRoutes.Get("/:id", h.GetAttribute)
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
	attributeDTOs := h.dtoManager.RewardManager().ToAttributeDTOs(attributes)

	return response.Response(200, "Status OK", attributeDTOs)
}

// @Tags Attributes
// @Summary Get One Attribute
// @Description Retrieves spesific Attribute based on the provided query parameters.
// @Accept json
// @Produce json
// @Param id path string false "Attribute ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/attributes/{id} [get]
func (h *PrivateHandler) GetAttribute(c *fiber.Ctx) error {
	id := c.Params("id")

	attribute, err := h.services.AttributeService().GetAttributeByID(c.Context(), id)
	if err != nil {
		return err
	}
	chapterDTO := h.dtoManager.RewardManager().ToAttributeDTO(attribute)

	return response.Response(200, "Status OK", chapterDTO)
}
