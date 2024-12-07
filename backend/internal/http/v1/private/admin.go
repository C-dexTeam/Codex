package private

import (
	"github.com/C-dexTeam/codex/internal/http/response"

	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initAdminRoutes(root fiber.Router) {
	adminRoute := root.Group("/admin")
	adminRoute.Use(h.adminRoleMiddleware)
	adminRoute.Get("/user", h.GetUsers)

}

// @Tags Admin
// @Summary Get All Users
// @Description Retrieves all logs based on the provided query parameters.
// @Accept json
// @Produce json
// @Param id query string false "User ID"
// @Param username query string false "Username"
// @Param email query string false "User's Email"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/admin/user [get]
func (h *PrivateHandler) GetUsers(c *fiber.Ctx) error {
	id := c.Query("id")
	username := c.Query("username")
	email := c.Query("email")
	page := c.Query("page")
	limit := c.Query("limit")

	users, err := h.services.AdminService().GetUsers(c.Context(), id, username, email, page, limit)
	if err != nil {
		return err
	}
	userAuthDTOs := h.dtoManager.AdminManager().ToUserAuthDTOs(users)

	return response.Response(200, "Status OK", userAuthDTOs)
}
