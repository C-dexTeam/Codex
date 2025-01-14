package private

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/internal/http/sessionStore"
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initUserRoutes(root fiber.Router) {
	user := root.Group("/user")
	user.Get("/profile", h.Profile)
	user.Post("/profile", h.UpdateProfile)
	user.Post("/connect", h.ConnectWallet)

	userAdminRoutes := root.Group("/admin/user")
	userAdminRoutes.Use(h.adminRoleMiddleware)
	userAdminRoutes.Get("/", h.GetUsers)
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
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(newUserProfile); err != nil {
		return err
	}

	// Update Profile
	if err := h.services.UserProfileService().Update(c.Context(), userSession.UserProfileID, newUserProfile.Name, newUserProfile.Surname); err != nil {
		return err
	}

	// Mevcut session'ı alıyoruz
	sess, err := h.sess_store.Get(c)
	if err != nil {
		return err
	}
	userSession.SetNameSurname(newUserProfile.Name, newUserProfile.Surname)
	sess.Set("user", userSession)
	if err := sess.Save(); err != nil {
		return err
	}

	return response.Response(200, "Status OK", nil)
}

// @Tags User
// @Summary Connect Wallet To User
// @Description Connects Wallet.
// @Accept json
// @Produce json
// @Param newWallet body dto.UserAuthWallet true "New User Wallet"
// @Success 200 {object} response.BaseResponse{}
// @Router /private/user/connect [post]
func (h *PrivateHandler) ConnectWallet(c *fiber.Ctx) error {
	userSession := sessionStore.GetSessionData(c)

	var newWallet dto.UserAuthWallet
	if err := c.BodyParser(&newWallet); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(newWallet); err != nil {
		return err
	}

	if err := h.services.UserService().ConnectWallet(c.Context(), userSession.UserID, newWallet.PublicKeyBase58, newWallet.Message, newWallet.Signature); err != nil {
		return err
	}

	userSession.SetPublicKey(newWallet.PublicKeyBase58)

	// Mevcut session'ı alıyoruz
	sess, err := h.sess_store.Get(c)
	if err != nil {
		return err
	}
	userSession.SetPublicKey(newWallet.PublicKeyBase58)
	sess.Set("user", userSession)
	if err := sess.Save(); err != nil {
		return err
	}

	// Get First Login Role
	walletUser, err := h.services.RoleService().GetByName(c.Context(), h.defaults.Roles.RoleWalletUser)
	if err != nil {
		return err
	}

	// If the user in user Role. Change The Users role to wallet-user.
	if userSession.Role != h.defaults.Roles.RoleAdmin {
		if err := h.services.UserProfileService().ChangeUserRole(c.Context(), userSession.UserProfileID, walletUser.ID.String()); err != nil {
			return err
		}
	}

	return response.Response(200, "Status OK", nil)
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
// @Router /private/admin/user [get]
func (h *PrivateHandler) GetUsers(c *fiber.Ctx) error {
	id := c.Query("id")
	username := c.Query("username")
	email := c.Query("email")
	page := c.Query("page")
	limit := c.Query("limit")

	users, err := h.services.UserService().GetUsers(c.Context(), id, username, email, page, limit)
	if err != nil {
		return err
	}
	userAuthDTOs := h.dtoManager.UserManager().ToUserAuthDTOs(users)

	return response.Response(200, "Status OK", userAuthDTOs)
}
