package private

import (
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initProgrammingLanguageRoutes(root fiber.Router) {
	pLanguageRoutes := root.Group("/planguages")
	pLanguageRoutes.Get("/", h.GetProgrammingLanguages)
	pLanguageRoutes.Get("/:id", h.GetProgrammingLanguage)
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
// @Router /private/planguages/ [get]
func (h *PrivateHandler) GetProgrammingLanguages(c *fiber.Ctx) error {
	id := c.Query("id")
	languageID := c.Query("languageID")
	name := c.Query("name")
	page := c.Query("page")
	limit := c.Query("limit")

	programmingLanguages, err := h.services.ProgrammingService().GetProgrammingLanguages(c.Context(), id, languageID, name, page, limit)
	if err != nil {
		return err
	}
	programmingLanguageDTOs := h.dtoManager.ProgrammingManager().ToPLanguageDTOs(programmingLanguages)

	return response.Response(200, "Status OK", programmingLanguageDTOs)
}

// @Tags Programming Language
// @Summary Get One Programming Language By ID
// @Description Retrieves spesific Programming languages based on the provided query parameters.
// @Accept json
// @Produce json
// @Param id path string false "Programming Language ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/planguages/{id} [get]
func (h *PrivateHandler) GetProgrammingLanguage(c *fiber.Ctx) error {
	id := c.Params("id")

	programmingLanguage, err := h.services.ProgrammingService().GetProgrammingLanguage(c.Context(), id)
	if err != nil {
		return err
	}
	programmingLanguageDTO := h.dtoManager.ProgrammingManager().ToPLanguageDTO(programmingLanguage)

	return response.Response(200, "Status OK", programmingLanguageDTO)
}
