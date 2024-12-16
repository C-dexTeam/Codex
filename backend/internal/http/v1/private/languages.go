package private

import (
	"fmt"

	"github.com/C-dexTeam/codex/internal/http/response"

	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initLanguageRoutes(root fiber.Router) {
	languageRoutes := root.Group("/language")
	languageRoutes.Get("/", h.GetLanguages)
}

// @Tags Language
// @Summary Get All Languages
// @Description Retrieves all languages based on the provided query parameters.
// @Accept json
// @Produce json
// @Param id query string false "Language ID"
// @Param value query string false "Value"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/language/ [get]
func (h *PrivateHandler) GetLanguages(c *fiber.Ctx) error {
	id := c.Query("id")
	value := c.Query("value")

	languages, err := h.services.LanguageService().GetLanguages(c.Context(), id, value)
	if err != nil {
		return err
	}
	fmt.Println(languages)
	languageDTOs := h.dtoManager.LanguageManager().ToLanguageDTOs(languages)

	return response.Response(200, "Status OK", languageDTOs)
}
