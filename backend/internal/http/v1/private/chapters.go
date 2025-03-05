package private

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/internal/http/sessionStore"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initChaptersRoutes(root fiber.Router) {
	chapterRoutes := root.Group("/chapters")
	chapterRoutes.Get("/", h.GetChapters)
	chapterRoutes.Get("/:id", h.GetChapter)
	chapterRoutes.Post("/run", h.RunChapter)
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
	userSession := sessionStore.GetSessionData(c)
	id := c.Params("id")
	page := c.Query("page")
	limit := c.Query("limit")

	chapter, err := h.services.ChapterService().GetChapter(c.Context(), id, page, limit)
	if err != nil {
		return err
	}

	if err := h.services.ChapterService().StartChapter(c.Context(), id, chapter.CourseID.String(), userSession.UserID); err != nil {
		return err
	}
	chapterDTO := h.dtoManager.ChapterManager().ToChapterDTO(chapter)

	return response.Response(200, "Status OK", chapterDTO)
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
	userSession := sessionStore.GetSessionData(c)
	var runChapter dto.RunChapter
	if err := c.BodyParser(&runChapter); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(runChapter); err != nil {
		return err
	}

	// Check if the ID's are correct.
	if _, err := h.services.ChapterService().GetChapter(c.Context(), runChapter.ChapterID, "1", "1"); err != nil {
		return err
	}
	course, err := h.services.CourseService().GetCourse(c.Context(), runChapter.CourseID, "1", "1")
	if err != nil {
		return err
	}

	// Check if the course started. If the user course exist than its started.
	if _, err := h.services.CourseService().UserCourse(c.Context(), userSession.UserID, runChapter.CourseID); err != nil {
		return err
	}

	chapter, tests, pLanguage, err := h.services.QuestService().GetQuest(c.Context(), runChapter.ChapterID, runChapter.CourseID)
	if err != nil {
		return err
	}
	quest := h.dtoManager.QuestManager().ToQuestDTO(chapter, tests, pLanguage, runChapter.UserCode)

	// Request the Run endpoint from Codex-Compiler
	codeResponse, err := h.services.ChapterService().Run(c.Context(), sessionID, *quest)
	if err != nil {
		return err
	}

	// CodeResponse will come from compiler-api
	if codeResponse.Correct {
		if chapter.RewardID.Valid {
			if err := h.services.RewardService().AddRewardIntoUser(c.Context(), userSession.UserID, runChapter.ChapterID, runChapter.CourseID, chapter.RewardID.UUID.String()); err != nil {
				return err
			}
		}
		if err := h.services.ChapterService().UpdateIsFinished(c.Context(), userSession.UserID, runChapter.ChapterID, runChapter.CourseID); err != nil {
			return err
		}
		progress, err := h.services.CourseService().UpdateUserCourseProgress(c.Context(), userSession.UserID, runChapter.CourseID)
		if err != nil {
			return err
		}
		// If the course finished. Than set course reward into user.
		if progress == 100 {
			if course.RewardID != nil {
				// TODO: Check this. U change the migration. Before than it works
				if err := h.services.RewardService().AddRewardIntoUser(c.Context(), userSession.UserID, "", runChapter.CourseID, course.RewardID.String()); err != nil {
					return err
				}
			}
		}
	}

	return response.Response(200, "Status OK", codeResponse)
}
