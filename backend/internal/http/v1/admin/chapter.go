package admin

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *AdminHandler) initChaptersRoutes(root fiber.Router) {
	chapterRoutes := root.Group("/chapters")

	chapterRoutes.Post("/", h.AddChapter)
	chapterRoutes.Patch("/", h.UpdateChapter)
	chapterRoutes.Delete("/:id", h.DeleteChapter)
}

// @Tags Chapters
// @Summary Add Chapter
// @Description Adds Chapter Into DB.
// @Accept json
// @Produce json
// @Param newChapter body dto.AddChapterDTO true "New Chapter"
// @Success 200 {object} response.BaseResponse{}
// @Router /admin/chapters/ [post]
func (h *AdminHandler) AddChapter(c *fiber.Ctx) error {
	var newChapter dto.AddChapterDTO
	if err := c.BodyParser(&newChapter); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(newChapter); err != nil {
		return err
	}

	// Default Language if the languageID is ""
	var languageID string
	if newChapter.LanguageID == "" {
		defaultLanguage, err := h.services.LanguageService().GetDefault(c.Context())
		if err != nil {
			return err
		}
		languageID = defaultLanguage.ID.String()
	} else {
		languageID = newChapter.LanguageID
	}

	if _, err := h.services.CourseService().GetCourse(c.Context(), newChapter.CourseID, "1", "1"); err != nil {
		return err
	}

	if newChapter.RewardID != "" {
		if _, err := h.services.RewardService().GetReward(c.Context(), newChapter.RewardID, "1", "1"); err != nil {
			return err
		}
	}

	id, err := h.services.ChapterService().AddChapter(
		c.Context(),
		newChapter.CourseID,
		languageID,
		newChapter.RewardID,
		newChapter.Title,
		newChapter.Description,
		newChapter.Content,
		newChapter.FuncName,
		newChapter.FrontendTmp,
		newChapter.DockerTmp,
		newChapter.Order,
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", id)
}

// @Tags Chapters
// @Summary Update Chapter
// @Description Updates Chapter Into DB.
// @Accept json
// @Produce json
// @Param updateChapter body dto.UpdateChapterDTO true "Update Chapters"
// @Success 200 {object} response.BaseResponse{}
// @Router /admin/chapters/ [patch]
func (h *AdminHandler) UpdateChapter(c *fiber.Ctx) error {
	var updateChapter dto.UpdateChapterDTO
	if err := c.BodyParser(&updateChapter); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(updateChapter); err != nil {
		return err
	}

	err := h.services.ChapterService().UpdateChapter(
		c.Context(),
		updateChapter.ID,
		updateChapter.CourseID,
		updateChapter.LanguageID,
		updateChapter.RewardID,
		updateChapter.Title,
		updateChapter.Description,
		updateChapter.Content,
		updateChapter.FuncName,
		updateChapter.FrontendTmp,
		updateChapter.DockerTmp,
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", nil)
}

// @Tags Chapters
// @Summary Delete Chapter
// @Description Delete Chapters from DB.
// @Accept json
// @Produce json
// @Param id path string false "Chapter ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /admin/chapters/{id} [delete]
func (h *AdminHandler) DeleteChapter(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.services.ChapterService().DeleteChapter(c.Context(), id); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}
