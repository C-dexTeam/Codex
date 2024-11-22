package private

import (
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/internal/http/sessionStore"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initUserRoutes(root fiber.Router) {
	user := root.Group("/user")
	user.Get("/profile", h.Profile)
}

// @Tags User
// @Summary Get User Profile
// @Description Retrieves users profile.
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/user/profile [get]
func (h *PrivateHandler) Profile(c *fiber.Ctx) error {
	userSession := sessionStore.GetSessionData(c)
	userProfileDTO := h.dtoManager.UserManager().ToUserProfile(*userSession)

	return response.Response(200, "Status OK", userProfileDTO)
}
