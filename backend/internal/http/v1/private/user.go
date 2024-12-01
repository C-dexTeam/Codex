package private

import (
	"fmt"

	"github.com/C-dexTeam/codex/internal/domains"
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/internal/http/sessionStore"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initUserRoutes(root fiber.Router) {
	user := root.Group("/user")
	user.Get("/profile", h.Profile)
	user.Post("/profile", h.UpdateProfile)
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

// @Tags User
// @Summary Update User Profile
// @Description Updates users profile.
// @Accept json
// @Produce json
// @Param newUserProfile body dto.UserProfileUpdateDTO true "New User Profile"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/user/profile [post]
func (h *PrivateHandler) UpdateProfile(c *fiber.Ctx) error {
	userSession := sessionStore.GetSessionData(c)

	var newUserProfile dto.UserProfileUpdateDTO
	if err := c.BodyParser(&newUserProfile); err != nil {
		fmt.Println(err)
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(newUserProfile); err != nil {
		return err
	}

	// Update Profile
	if err := h.services.UserProfileService().Update(c.Context(), userSession.UserProfileID, newUserProfile.Name, newUserProfile.Surname); err != nil {
		return err
	}

	// Get Default Role
	defaultRole, err := h.services.RoleService().GetByName(c.Context(), domains.DefaultRole)
	if err != nil {
		return err
	}

	// Get First Login Role
	firstLoginRole, err := h.services.RoleService().GetByName(c.Context(), domains.FirstLogin)
	if err != nil {
		return err
	}
	// If the user First-Login Role. Change The Users role to defaultRole.
	if userSession.RoleID == firstLoginRole.GetID().String() {
		if err := h.services.UserProfileService().ChangeUserRole(c.Context(), userSession.UserProfileID, defaultRole.GetID().String()); err != nil {
			return err
		}
	}

	return response.Response(200, "Status OK", nil)
}
