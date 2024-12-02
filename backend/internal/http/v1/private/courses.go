package private

import (
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initCoursesRoutes(root fiber.Router) {
	coursesRoutes := root.Group("/courses")
	coursesRoutes.Get("/", h.GetCourses)
}

// @Tags Courses
// @Summary Get All Courses
// @Description Retrieves all courses based on the provided query parameters.
// @Accept json
// @Produce json
// @Param courseID query string false "Course ID"
// @Param languageID query string false "Language ID"
// @Param pLanguageID query string false "Programming Language ID"
// @Param title query string false "Course Title"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/courses/ [get]
func (h *PrivateHandler) GetCourses(c *fiber.Ctx) error {
	courseID := c.Query("courseID")
	languageID := c.Query("languageID")
	pLanguageID := c.Query("pLanguageID")
	title := c.Query("title")
	page := c.Query("page")
	limit := c.Query("limit")

	courses, err := h.services.CourseService().GetCourses(c.Context(), courseID, languageID, pLanguageID, title, page, limit)
	if err != nil {
		return err
	}
	courseDTOs := h.dtoManager.CourseManager().ToCourseDTOs(courses)

	return response.Response(200, "Status OK", courseDTOs)
}
