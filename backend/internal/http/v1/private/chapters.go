package private

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initChaptersRoutes(root fiber.Router) {
	chapterRoutes := root.Group("/chapters")
	chapterRoutes.Get("/", h.GetChapters)
	chapterRoutes.Get("/:id", h.GetChapter)
	chapterRoutes.Post("/run", h.RunChapter)

	chapterAdminRoutes := root.Group("/admin/chapters")
	chapterAdminRoutes.Use(h.adminRoleMiddleware)
	chapterAdminRoutes.Post("/", h.AddChapter)
	chapterAdminRoutes.Patch("/", h.UpdateChapter)
	chapterAdminRoutes.Delete("/:id", h.DeleteChapter)
}

// @Tags Chapters
// @Summary Get All Chapters
// @Description Retrieves all chapters based on the provided query parameters.
// @Accept json
// @Produce json
// @Param id query string false "Chapter ID"
// @Param languageID query string false "Language ID"
// @Param courseID query string false "Course ID"
// @Param rewardID query string false "Reward ID"
// @Param title query string false "Chapter Title"
// @Param grantsExp query string false "Grants Experience"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/chapters/ [get]
func (h *PrivateHandler) GetChapters(c *fiber.Ctx) error {
	id := c.Query("id")
	languageID := c.Query("languageID")
	courseID := c.Query("courseID")
	rewardID := c.Query("rewardID")
	title := c.Query("title")
	page := c.Query("page")
	limit := c.Query("limit")

	// if you put "" in bool area. Its means all. Like not only true or false.
	chapters, err := h.services.ChapterService().GetChapters(c.Context(), id, languageID, courseID, rewardID, title, "", "", page, limit)
	if err != nil {
		return err
	}
	chapterDTOs := h.dtoManager.ChapterManager().ToChapterDTOs(chapters)

	return response.Response(200, "Status OK", chapterDTOs)
}

// @Tags Chapters
// @Summary Get One Chapter
// @Description Retrieves spesific chapters based on the provided query parameters.
// @Accept json
// @Produce json
// @Param id path string false "Chapter ID"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/chapters/{id} [get]
func (h *PrivateHandler) GetChapter(c *fiber.Ctx) error {
	id := c.Params("id")
	page := c.Query("page")
	limit := c.Query("limit")

	chapter, err := h.services.ChapterService().GetChapter(c.Context(), id, page, limit)
	if err != nil {
		return err
	}
	chapterDTO := h.dtoManager.ChapterManager().ToChapterDTO(chapter)

	return response.Response(200, "Status OK", chapterDTO)
}

// @Tags Chapters
// @Summary Add Chapter
// @Description Adds Chapter Into DB.
// @Accept json
// @Produce json
// @Param newChapter body dto.AddChapterDTO true "New Chapter"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/chapters/ [post]
func (h *PrivateHandler) AddChapter(c *fiber.Ctx) error {
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
		if _, _, err := h.services.RewardService().GetReward(c.Context(), newChapter.RewardID, "1", "1"); err != nil {
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
		newChapter.CheckTmp,
		newChapter.GrantsExperience,
		newChapter.Active,
		newChapter.RewardAmount,
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
// @Router /private/admin/chapters/ [patch]
func (h *PrivateHandler) UpdateChapter(c *fiber.Ctx) error {
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
		updateChapter.CheckTmp,
		updateChapter.GrantsExperience,
		updateChapter.Active,
		updateChapter.RewardAmount,
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
// @Router /private/admin/chapters/{id} [delete]
func (h *PrivateHandler) DeleteChapter(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.services.ChapterService().DeleteChapter(c.Context(), id); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}

// @Tags Chapters
// @Summary Run Chapter
// @Description Runs Chapter Code.
// @Accept json
// @Produce json
// @Param runChapter body dto.RunChapter true "Runs Chapter's Code"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/chapters/run [post]
func (h *PrivateHandler) RunChapter(c *fiber.Ctx) error {
	sessionID := c.Cookies("session_id")
	var runChapter dto.RunChapter
	if err := c.BodyParser(&runChapter); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(runChapter); err != nil {
		return err
	}

	chapter, tests, pLanguage, err := h.services.QuestService().GetQuest(c.Context(), runChapter.ChapterID, runChapter.CourseID)
	if err != nil {
		return err
	}
	quest := h.dtoManager.QuestManager().ToQuestDTO(chapter, tests, pLanguage, runChapter.UserCode)

	// Request the Run endpoint from Codex-Compiler
	err = h.services.ChapterService().Run(c.Context(), sessionID, *quest)
	if err.Error() != "" {
		return err
	}

	return response.Response(200, "Status OK", err)
}
