package private

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initTestsRoutes(root fiber.Router) {
	testRoutes := root.Group("/tests")
	testRoutes.Get("/", h.GetTests)

	testAdmin := root.Group("/admin/tests")
	testAdmin.Use(h.adminRoleMiddleware)
	testAdmin.Patch("/", h.UpdateTest)
	testAdmin.Delete("/:id", h.DeleteTest)
	testAdmin.Post("/", h.AddTest)
}

// @Tags Test
// @Summary Get All tests
// @Description Retrieves all tests based on the provided query parameters.
// @Accept json
// @Produce json
// @Param id query string false "Test ID"
// @Param chapterID query string false "Chapter ID"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/tests/ [get]
func (h *PrivateHandler) GetTests(c *fiber.Ctx) error {
	id := c.Query("id")
	chapterID := c.Query("chapterID")
	page := c.Query("page")
	limit := c.Query("limit")

	tests, err := h.services.TestService().GetTests(c.Context(), id, chapterID, page, limit)
	if err != nil {
		return err
	}
	testDTOs := h.dtoManager.TestManager().ToTestDTOs(tests)

	return response.Response(200, "Status OK", testDTOs)
}

// @Tags Test
// @Summary Add Test
// @Description Adds Test Into DB.
// @Accept json
// @Produce json
// @Param newTest body dto.AddTestDTO true "New Test"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/tests/ [post]
func (h *PrivateHandler) AddTest(c *fiber.Ctx) error {
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
// @Router /private/admin/tests/ [patch]
func (h *PrivateHandler) UpdateTest(c *fiber.Ctx) error {
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
// @Router /private/admin/tests/{id} [delete]
func (h *PrivateHandler) DeleteTest(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.services.TestService().DeleteTest(c.Context(), id); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}
