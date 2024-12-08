package private

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initProgrammingLanguageRoutes(root fiber.Router) {
	pLanguageRoutes := root.Group("/planguages")
	pLanguageRoutes.Get("/", h.GetProgrammingLanguages)
	pLanguageRoutes.Get("/:id", h.GetProgrammingLanguage)

	pLanguagesAdminRoutes := root.Group("/admin/planguages")
	pLanguagesAdminRoutes.Use(h.adminRoleMiddleware)
	pLanguagesAdminRoutes.Post("/", h.AddProgrammingLanguage)
	pLanguagesAdminRoutes.Patch("/", h.UpdateProgrammingLanguage)
	pLanguagesAdminRoutes.Delete("/", h.DeleteProgrammingLanguage)
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

// @Tags Programming Language
// @Summary Add Programming Language
// @Description Adds Programming Language Into DB.
// @Accept json
// @Produce json
// @Param newPLanguage body dto.AddPLanguageDTO true "New Programming Language"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/planguages/ [post]
func (h *PrivateHandler) AddProgrammingLanguage(c *fiber.Ctx) error {
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
		languageID = defaultLanguage.GetID().String()
	} else {
		languageID = newPLanguage.LanguageID
	}

	id, err := h.services.ProgrammingService().AddProgrammingLanguage(
		c.Context(),
		languageID,
		newPLanguage.Name,
		newPLanguage.Description,
		newPLanguage.DownloadCMD,
		newPLanguage.CompileCMD,
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
// @Router /private/admin/planguages/ [patch]
func (h *PrivateHandler) UpdateProgrammingLanguage(c *fiber.Ctx) error {
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
		updatePLanguage.DownloadCMD,
		updatePLanguage.CompileCMD,
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
// @Param id query string false "Programming Language ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/planguages/ [delete]
func (h *PrivateHandler) DeleteProgrammingLanguage(c *fiber.Ctx) error {
	id := c.Query("id")

	if err := h.services.ProgrammingService().DeleteProgrammingLanguage(c.Context(), id); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}
