package private

import (
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initChaptersRoutes(root fiber.Router) {
	chapterRoutes := root.Group("/chapters")
	chapterRoutes.Get("/", h.GetChapters)
}

// @Tags Chapter
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
