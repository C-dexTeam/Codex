package private

import (
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initTestsRoutes(root fiber.Router) {
	testRoutes := root.Group("/tests")
	testRoutes.Get("/", h.GetTests)
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
