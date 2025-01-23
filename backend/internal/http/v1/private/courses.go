package private

import (
	"path/filepath"

	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/internal/http/sessionStore"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *PrivateHandler) initCoursesRoutes(root fiber.Router) {
	coursesRoutes := root.Group("/courses")
	coursesRoutes.Get("/", h.GetCourses)
	coursesRoutes.Get("/popular", h.GetPopulerCourses)
	coursesRoutes.Get("/:id", h.GetCourse)
	coursesRoutes.Post("/start", h.StartCourse)

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
// @Summary Add Course
// @Description Adds Course Into DB.
// @Accept multipart/form-data
// @Produce json
// @Param imageFile formData file true "Course Image File"
// @Param courseInfo formData dto.AddCourseDTO true "Course Information"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/courses/ [post]
func (h *PrivateHandler) AddCourse(c *fiber.Ctx) error {
	var courseInfo dto.AddCourseDTO
	if err := c.BodyParser(&courseInfo); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(courseInfo); err != nil {
		return err
	}

	// Dosya alanını alıyoruz (sadece imageFile)
	imageFile, err := c.FormFile("imageFile")
	if err != nil {
		return err
	}

	// Dosyayı kaydetme işlemi
	extention := filepath.Ext(imageFile.Filename)
	imagePath := h.services.UploadService().MainDir() + "/" + uuid.New().String() + extention
	if err := h.services.UploadService().SaveImage(imageFile, imagePath); err != nil {
		return err
	}

	// Default Language if the languageID is ""
	var languageID string
	if courseInfo.LanguageID == "" {
		defaultLanguage, err := h.services.LanguageService().GetDefault(c.Context())
		if err != nil {
			return err
		}
		languageID = defaultLanguage.ID.String()
	} else {
		languageID = courseInfo.LanguageID
	}

	// Checks if exists
	if courseInfo.ProgrammingLanguageID != "" {
		if _, err := h.services.ProgrammingService().GetProgrammingLanguage(c.Context(), courseInfo.ProgrammingLanguageID); err != nil {
			return err
		}
	}

	id, err := h.services.CourseService().AddCourse(
		c.Context(),
		languageID,
		courseInfo.ProgrammingLanguageID,
		courseInfo.RewardID,
		courseInfo.Title,
		courseInfo.Description,
		imagePath,
		courseInfo.RewardAmount,
	)
	if err != nil {
		return err
	}

	return response.Response(200, "Status OK", id)
}

// @Tags Courses
// @Summary Update Course
// @Description Updates Course Into DB.
// @Accept multipart/form-data
// @Produce json
// @Param imageFile formData file true "Course Image File"
// @Param courseInfo formData dto.UpdateCourseDTO true "Update Course"
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

	// Dosya alanını alıyoruz (sadece imageFile)
	imageFile, err := c.FormFile("imageFile")
	if err != nil {
		return err
	}

	var newImagePath string
	if imageFile != nil {
		extention := filepath.Ext(imageFile.Filename)
		newImagePath = h.services.UploadService().MainDir() + "/" + uuid.New().String() + extention
		if err := h.services.UploadService().SaveImage(imageFile, newImagePath); err != nil {
			return err
		}
	}

	courses, err := h.services.CourseService().GetCourse(c.Context(), updateCourse.ID, "1", "1")
	if err != nil {
		return err
	}
	if courses.ImagePath != "" {
		if err := h.services.UploadService().DeleteFile(courses.ImagePath); err != nil {
			return err
		}
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

	err = h.services.CourseService().UpdateCourse(
		c.Context(),
		updateCourse.ID,
		updateCourse.LanguageID,
		updateCourse.PLanguageID,
		updateCourse.RewardID,
		updateCourse.Title,
		updateCourse.Description,
		newImagePath,
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

	if err := h.services.CourseService().DeleteCourse(c.Context(), id); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
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
