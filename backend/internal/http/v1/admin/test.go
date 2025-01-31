package admin

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *AdminHandler) initTestsRoutes(root fiber.Router) {
	testRoutes := root.Group("/tests")

	testRoutes.Patch("/", h.UpdateTest)
	testRoutes.Delete("/:id", h.DeleteTest)
	testRoutes.Post("/", h.AddTest)
}

// @Tags Test
// @Summary Add Test
// @Description Adds Test Into DB.
// @Accept json
// @Produce json
// @Param newTest body dto.AddTestDTO true "New Test"
// @Success 200 {object} response.BaseResponse{}
// @Router /admin/tests/ [post]
func (h *AdminHandler) AddTest(c *fiber.Ctx) error {
	var newTest dto.AddTestDTO
	if err := c.BodyParser(&newTest); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(newTest); err != nil {
		return err
	}

	id, err := h.services.TestService().AddTest(
		c.Context(),
		newTest.ChapterID,
		newTest.InputValue,
		newTest.OutputValue,
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", id)
}

// @Tags Test
// @Summary Update Test
// @Description Updates tests Into DB.
// @Accept json
// @Produce json
// @Param updateTest body dto.UpdateTestDTO true "Update Test"
// @Success 200 {object} response.BaseResponse{}
// @Router /admin/tests/ [patch]
func (h *AdminHandler) UpdateTest(c *fiber.Ctx) error {
	var updateTest dto.UpdateTestDTO
	if err := c.BodyParser(&updateTest); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(updateTest); err != nil {
		return err
	}

	err := h.services.TestService().UpdateTest(
		c.Context(),
		updateTest.ID,
		updateTest.InputValue,
		updateTest.OutputValue,
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", nil)
}

// @Tags Test
// @Summary Delete Test
// @Description Delete tests from DB.
// @Accept json
// @Produce json
// @Param id path string false "Taest ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /admin/tests/{id} [delete]
func (h *AdminHandler) DeleteTest(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.services.TestService().DeleteTest(c.Context(), id); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}
