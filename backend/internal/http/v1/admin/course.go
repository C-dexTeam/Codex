package admin

import (
	"path/filepath"

	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *AdminHandler) initCoursesRoutes(root fiber.Router) {
	coursesRoutes := root.Group("/courses")

	coursesRoutes.Post("/", h.AddCourse)
	coursesRoutes.Patch("/", h.UpdateCourse)
	coursesRoutes.Delete("/:id", h.DeleteCourse)
}

// @Tags Courses
// @Summary Add Course
// @Description Adds Course Into DB.
// @Accept multipart/form-data
// @Produce json
// @Param imageFile formData file true "Course Image File"
// @Param courseInfo formData dto.AddCourseDTO true "Course Information"
// @Success 200 {object} response.BaseResponse{}
// @Router /admin/courses/ [post]
func (h *AdminHandler) AddCourse(c *fiber.Ctx) error {
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
		"imagePath",
		courseInfo.RewardAmount,
	)
	if err != nil {
		return err
	}

	// Save Course Image
	extention := filepath.Ext(imageFile.Filename)
	imagePath := h.services.UploadService().CourseDir() + "/" + id.String() + extention
	if err := h.services.UploadService().SaveImage(imageFile, imagePath); err != nil {
		return err
	}

	if err := h.services.CourseService().UpdateCourse(c.Context(), id.String(), "", "", "", "", "", imagePath, 0); err != nil {
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
// @Router /admin/courses/ [patch]
func (h *AdminHandler) UpdateCourse(c *fiber.Ctx) error {
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
		newImagePath = h.services.UploadService().CourseDir() + "/" + updateCourse.ID + extention
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
// @Router /admin/courses/{id} [delete]
func (h *AdminHandler) DeleteCourse(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.services.CourseService().DeleteCourse(c.Context(), id); err != nil {
		return err
	}
	return response.Response(200, "Status OK", nil)
}
