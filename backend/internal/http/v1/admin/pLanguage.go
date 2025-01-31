package admin

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *AdminHandler) initProgrammingLanguageRoutes(root fiber.Router) {
	pLanguageRoutes := root.Group("/planguages")

	pLanguageRoutes.Post("/", h.AddProgrammingLanguage)
	pLanguageRoutes.Patch("/", h.UpdateProgrammingLanguage)
	pLanguageRoutes.Delete("/:id", h.DeleteProgrammingLanguage)
}

// @Tags Programming Language
// @Summary Add Programming Language
// @Description Adds Programming Language Into DB.
// @Accept json
// @Produce json
// @Param newPLanguage body dto.AddPLanguageDTO true "New Programming Language"
// @Success 200 {object} response.BaseResponse{}
// @Router /admin/planguages/ [post]
func (h *AdminHandler) AddProgrammingLanguage(c *fiber.Ctx) error {
	var newPLanguage dto.AddPLanguageDTO
	if err := c.BodyParser(&newPLanguage); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(newPLanguage); err != nil {
		return err
	}

	// Default Language if the languageID is ""
	var languageID string
	if newPLanguage.LanguageID == "" {
		defaultLanguage, err := h.services.LanguageService().GetDefault(c.Context())
		if err != nil {
			return err
		}
		languageID = defaultLanguage.ID.String()
	} else {
		languageID = newPLanguage.LanguageID
	}

	id, err := h.services.ProgrammingService().AddProgrammingLanguage(
		c.Context(),
		languageID,
		newPLanguage.Name,
		newPLanguage.Description,
		newPLanguage.ImagePath,
		newPLanguage.FileExtention,
		newPLanguage.MonacoEditor,
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", id)
}

// @Tags Programming Language
// @Summary Update Programming Language
// @Description Updates Programming Language Into DB.
// @Accept json
// @Produce json
// @Param updatePLanguage body dto.UpdatePLanguageDTO true "Update Programming Language"
// @Success 200 {object} response.BaseResponse{}
// @Router /admin/planguages/ [patch]
func (h *AdminHandler) UpdateProgrammingLanguage(c *fiber.Ctx) error {
	var updatePLanguage dto.UpdatePLanguageDTO
	if err := c.BodyParser(&updatePLanguage); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(updatePLanguage); err != nil {
		return err
	}

	err := h.services.ProgrammingService().UpdateProgrammingLanguage(
		c.Context(),
		updatePLanguage.ID,
		updatePLanguage.LanguageID,
		updatePLanguage.Name,
		updatePLanguage.Description,
		updatePLanguage.ImagePath,
		updatePLanguage.FileExtention,
		updatePLanguage.MonacoEditor,
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", nil)
}

// @Tags Programming Language
// @Summary Delete Programming Language
// @Description Delete Programming Languages from DB.
// @Accept json
// @Produce json
// @Param id path string false "Programming Language ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /admin/planguages/{id} [delete]
func (h *AdminHandler) DeleteProgrammingLanguage(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.services.ProgrammingService().DeleteProgrammingLanguage(c.Context(), id); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}
