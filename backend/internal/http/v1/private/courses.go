package private

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initCoursesRoutes(root fiber.Router) {
	coursesRoutes := root.Group("/courses")
	coursesRoutes.Get("/", h.GetCourses)

	courseAdminRoutes := root.Group("/admin/courses")
	courseAdminRoutes.Use(h.adminRoleMiddleware)
	courseAdminRoutes.Post("/", h.AddCourse)

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
	courseDTOs := h.dtoManager.CourseManager().ToCourseDTOs(courses)

	return response.Response(200, "Status OK", courseDTOs)
}

// @Tags Courses
// @Summary Add Course
// @Description Adds Course Into DB.
// @Accept json
// @Produce json
// @Param newCourse body dto.AddCourseDTO true "New Course"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/courses/ [post]
func (h *PrivateHandler) AddCourse(c *fiber.Ctx) error {
	var newCourse dto.AddCourseDTO
	if err := c.BodyParser(&newCourse); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(newCourse); err != nil {
		return err
	}

	// Default Language if the languageID is ""
	var languageID string
	if newCourse.LanguageID == "" {
		defaultLanguage, err := h.services.LanguageService().GetDefault(c.Context())
		if err != nil {
			return err
		}
		languageID = defaultLanguage.GetID().String()
	} else {
		languageID = newCourse.LanguageID
	}

	if newCourse.PLanguageID != "" {
		if _, err := h.services.ProgrammingService().GetProgrammingLanguage(c.Context(), newCourse.PLanguageID); err != nil {
			return err
		}
	}

	id, err := h.services.CourseService().AddCourse(
		c.Context(),
		languageID,
		newCourse.PLanguageID,
		newCourse.RewardID,
		newCourse.Title,
		newCourse.Description,
		newCourse.ImagePath,
		newCourse.RewardAmount,
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", id)
}
