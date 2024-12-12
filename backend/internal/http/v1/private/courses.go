package private

import (
	"fmt"

	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initCoursesRoutes(root fiber.Router) {
	coursesRoutes := root.Group("/courses")
	coursesRoutes.Get("/", h.GetCourses)
	coursesRoutes.Get("/:id", h.GetCourse)

	// TODO: Reward Atama için Ayrı Endpoint
	courseAdminRoutes := root.Group("/admin/courses")
	courseAdminRoutes.Use(h.adminRoleMiddleware)
	courseAdminRoutes.Post("/", h.AddCourse)
	courseAdminRoutes.Patch("/", h.UpdateCourse)
	courseAdminRoutes.Delete("/:id", h.DeleteCourse)
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
	courseDTO := h.dtoManager.CourseManager().ToCourseDTO(*course)

	return response.Response(200, "Status OK", courseDTO)
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

// @Tags Courses
// @Summary Update Course
// @Description Updates Course Into DB.
// @Accept json
// @Produce json
// @Param updateCourse body dto.UpdateCourseDTO true "Update Course"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/courses/ [patch]
func (h *PrivateHandler) UpdateCourse(c *fiber.Ctx) error {
	var updateCourse dto.UpdateCourseDTO
	if err := c.BodyParser(&updateCourse); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(updateCourse); err != nil {
		return err
	}

	// Check if is exists
	if updateCourse.LanguageID != "" {
		if _, err := h.services.LanguageService().GetLanguage(c.Context(), updateCourse.LanguageID); err != nil {
			return err
		}
	}
	if updateCourse.PLanguageID != "" {
		if _, err := h.services.ProgrammingService().GetProgrammingLanguage(c.Context(), updateCourse.PLanguageID); err != nil {
			return err
		}
	}

	err := h.services.CourseService().UpdateCourse(
		c.Context(),
		updateCourse.ID,
		updateCourse.LanguageID,
		updateCourse.PLanguageID,
		updateCourse.RewardID,
		updateCourse.Title,
		updateCourse.Description,
		updateCourse.ImagePath,
		updateCourse.RewardAmount,
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", nil)
}

// @Tags Courses
// @Summary Delete Course
// @Description Delete Courses from DB.
// @Accept json
// @Produce json
// @Param id path string true "Course ID"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/courses/{id} [delete]
func (h *PrivateHandler) DeleteCourse(c *fiber.Ctx) error {
	id := c.Params("id")

	fmt.Println(id)

	if err := h.services.CourseService().DeleteCourse(c.Context(), id); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}
