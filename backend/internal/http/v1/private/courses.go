package private

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/internal/http/sessionStore"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initCoursesRoutes(root fiber.Router) {
	coursesRoutes := root.Group("/courses")

	coursesRoutes.Get("/", h.GetCourses)
	coursesRoutes.Get("/popular", h.GetPopulerCourses)
	coursesRoutes.Get("/:id", h.GetCourse)
	coursesRoutes.Post("/start", h.StartCourse)
}

// @Tags Courses
// @Summary Get All Courses
// @Description Retrieves all courses based on the provided query parameters.
// @Accept json
// @Produce json
// @Param id query string false "Course ID"
// @Param languageID query string false "Language ID"
// @Param pLanguageID query string false "Programming Language ID"
// @Param title query string false "Course Title"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/courses/ [get]
func (h *PrivateHandler) GetCourses(c *fiber.Ctx) error {
	id := c.Query("id")
	languageID := c.Query("languageID")
	pLanguageID := c.Query("pLanguageID")
	title := c.Query("title")
	page := c.Query("page")
	limit := c.Query("limit")

	courses, err := h.services.CourseService().GetCourses(c.Context(), id, languageID, pLanguageID, title, page, limit)
	if err != nil {
		return err
	}
	courseDTOs := h.dtoManager.CourseManager().ToCourseDTOCount(courses)

	return response.Response(200, "Status OK", courseDTOs)
}

// @Tags Courses
// @Summary Get All Popular Courses
// @Description Retrieves all popular courses.
// @Accept json
// @Produce json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/courses/popular [get]
func (h *PrivateHandler) GetPopulerCourses(c *fiber.Ctx) error {
	page := c.Query("page")
	limit := c.Query("limit")

	courses, err := h.services.CourseService().GetPopularCourses(c.Context(), page, limit)
	if err != nil {
		return err
	}
	courseDTOs := h.dtoManager.CourseManager().ToCourseDTOs(courses)

	return response.Response(200, "Status OK", courseDTOs)
}

// @Tags Courses
// @Summary Get Course By ID
// @Description Retrieves one course.
// @Accept json
// @Produce json
// @Param id path string false "Course ID"
// @Param page query string false "Chapter Page"
// @Param limit query string false "Chapter Attribute Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/courses/{id} [get]
func (h *PrivateHandler) GetCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	page := c.Query("page")
	limit := c.Query("limit")

	course, err := h.services.CourseService().GetCourse(c.Context(), id, page, limit)
	if err != nil {
		return err
	}
	courseDTO := h.dtoManager.CourseManager().ToCourseDTO(course)

	return response.Response(200, "Status OK", courseDTO)
}

// @Tags Courses
// @Summary Starts Course
// @Description Starts the spesific course.
// @Accept json
// @Param startCourse body dto.StartCourseDTO true "Start Course"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/courses/start [post]
func (h *PrivateHandler) StartCourse(c *fiber.Ctx) error {
	userSession := sessionStore.GetSessionData(c)
	var startCourse dto.StartCourseDTO
	if err := c.BodyParser(&startCourse); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(startCourse); err != nil {
		return err
	}
	id, err := h.services.CourseService().StartCourse(c.Context(), startCourse.ID, userSession.UserID)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", id)
}
