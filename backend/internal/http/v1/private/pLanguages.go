package private

import (
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initProgrammingLanguageRoutes(root fiber.Router) {
	programmingLanguageRoutes := root.Group("/programmingLanguages")
	programmingLanguageRoutes.Get("/", h.GetProgrammingLanguages)
}

// @Tags Programming Language
// @Summary Get All Programming Languages
// @Description Retrieves all Programming languages based on the provided query parameters.
// @Accept json
// @Produce json
// @Param id query string false "Programming Language ID"
// @Param languageID query string false "Language ID"
// @Param name query string false "Programming Language Name"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/programmingLanguages/ [get]
func (h *PrivateHandler) GetProgrammingLanguages(c *fiber.Ctx) error {
	id := c.Query("id")
	languageID := c.Query("languageID")
	name := c.Query("name")
	page := c.Query("page")
	limit := c.Query("limit")

	programmingLanguages, err := h.services.ProgrammingLService().GetProgrammingLanguages(c.Context(), id, languageID, name, page, limit)
	if err != nil {
		return err
	}
	programmingLanguageDTOs := h.dtoManager.ProgrammingManager().ToPLanguageDTOs(programmingLanguages)

	return response.Response(200, "Status OK", programmingLanguageDTOs)
}
