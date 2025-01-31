package admin

import (
	"fmt"

	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/gofiber/fiber/v2"
)

func (h *AdminHandler) initUserRoutes(root fiber.Router) {
	user := root.Group("/user")

	user.Get("/", h.GetUsers)
}

// @Tags User
// @Summary Get All Users
// @Description Retrieves all users based on the provided query parameters.
// @Accept json
// @Produce json
// @Param id query string false "User ID"
// @Param username query string false "Username"
// @Param email query string false "User's Email"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} response.BaseResponse{}
// @Router /admin/user [get]
func (h *AdminHandler) GetUsers(c *fiber.Ctx) error {
	fmt.Println("Selam")

	id := c.Query("id")
	username := c.Query("username")
	email := c.Query("email")
	page := c.Query("page")
	limit := c.Query("limit")

	users, err := h.services.UserService().GetUsers(c.Context(), id, username, email, page, limit)
	if err != nil {
		return err
	}
	userAuthDTOs := h.dtoManager.UserManager().ToUserAuthViews(users)

	return response.Response(200, "Status OK", userAuthDTOs)
}
