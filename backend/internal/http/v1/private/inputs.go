package private

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initInputRoutes(root fiber.Router) {
	inputRoutes := root.Group("/input")
	inputRoutes.Get("/", h.GetInputs)

	// TODO: Reward Atama için Ayrı Endpoint
	inputAdminRoutes := root.Group("/admin/inputs")
	inputAdminRoutes.Use(h.adminRoleMiddleware)
	inputAdminRoutes.Post("/", h.AddInput)
}

// @Tags Inputs
// @Summary Get All inputs
// @Description Retrieves all inputs based on the provided query parameters.
// @Accept json
// @Produce json
// @Param id query string false "Input ID"
// @Param testID query string false "Input's Test ID"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/input/ [get]
func (h *PrivateHandler) GetInputs(c *fiber.Ctx) error {
	id := c.Query("id")
	testID := c.Query("testID")
	page := c.Query("page")
	limit := c.Query("limit")

	inputs, err := h.services.TestService().GetInputs(c.Context(), id, testID, page, limit)
	if err != nil {
		return err
	}
	inputDTOs := h.dtoManager.TestManager().ToInputDTOs(inputs)

	return response.Response(200, "Status OK", inputDTOs)
}

// @Tags Inputs
// @Summary Add Input
// @Description Adds Input Into DB.
// @Accept json
// @Produce json
// @Param newInput body dto.AddGeneralDTO true "New Input"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/input/ [post]
func (h *PrivateHandler) AddInput(c *fiber.Ctx) error {
	var newInput dto.AddGeneralDTO
	if err := c.BodyParser(&newInput); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(newInput); err != nil {
		return err
	}

	id, err := h.services.TestService().AddInput(
		c.Context(),
		newInput.TestID,
		newInput.Value,
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", id)
}
