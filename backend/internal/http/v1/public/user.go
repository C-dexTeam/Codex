package public

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/internal/http/sessionStore"

	"github.com/gofiber/fiber/v2"
)

func (h *PublicHandler) initUserRoutes(root fiber.Router) {
	root.Post("/login", h.Login)
	root.Post("/register", h.Register)
	root.Post("/wallet", h.AuthWallet)
	root.Post("/logout", h.Logout)
}

// @Tags Auth
// @Summary Login
// @Description Login
// @Accept json
// @Produce json
// @Param login body dto.UserLoginDTO true "Login"
// @Success 200 {object} response.BaseResponse{}
// @Router /public/login [post]
func (h *PublicHandler) Login(c *fiber.Ctx) error {
	var login dto.UserLoginDTO
	if err := c.BodyParser(&login); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(login); err != nil {
		return err
	}
	userAuthData, err := h.services.UserService().Login(c.Context(), login.Username, login.Password)
	if err != nil {
		return err
	}

	userProfile, err := h.services.UserProfileService().GetUserWithUserAuthID(c.Context(), userAuthData.ID.String())
	if err != nil {
		return err
	}

	userRole, err := h.services.RoleService().GetRoleByID(c.Context(), userProfile.RoleID)
	if err != nil {
		return err
	}

	statistic, err := h.services.UserProfileService().UserStatistic(c.Context(), userAuthData.ID.String())
	if err != nil {
		return err
	}

	userRewards, err := h.services.RewardService().GetUserRewards(c.Context(), userAuthData.ID.String(), "", "")
	if err != nil {
		return err
	}

	sess, err := h.sessionStore.Get(c)
	if err != nil {
		return err
	}
	sessionData := sessionStore.SessionData{}
	sessionData.ParseFromUser(userAuthData, userProfile, userRole)
	sess.Set("user", sessionData)
	if err := sess.Save(); err != nil {
		return err
	}
	profileResponse := h.dtoManager.UserManager().ToUserProfile(sessionData, statistic, userRewards, sessionData.Streak)

	return response.Response(200, "Login successful", profileResponse)
}

// @Tags Auth
// @Summary Auth Wallet
// @Description Auth Wallet
// @Accept json
// @Produce json
// @Param wallet body dto.UserAuthWallet true "Wallet"
// @Success 200 {object} response.BaseResponse{}
// @Router /public/wallet [post]
func (h *PublicHandler) AuthWallet(c *fiber.Ctx) error {
	var wallet dto.UserAuthWallet
	if err := c.BodyParser(&wallet); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(wallet); err != nil {
		return err
	}

	userAuthData, err := h.services.UserService().AuthWallet(
		c.Context(),
		wallet.PublicKeyBase58,
		wallet.Message,
		wallet.Signature,
	)
	if err != nil {
		return err
	}

	userProfile, err := h.services.UserProfileService().GetUserWithUserAuthID(c.Context(), userAuthData.ID.String())
	if err != nil {
		return err
	}

	statistic, err := h.services.UserProfileService().UserStatistic(c.Context(), userAuthData.ID.String())
	if err != nil {
		return err
	}

	userRewards, err := h.services.RewardService().GetUserRewards(c.Context(), userAuthData.ID.String(), "", "")
	if err != nil {
		return err
	}

	userRole, err := h.services.RoleService().GetRoleByID(c.Context(), userProfile.RoleID)
	if err != nil {
		return err
	}

	sess, err := h.sessionStore.Get(c)
	if err != nil {
		return err
	}
	sessionData := sessionStore.SessionData{}
	sessionData.ParseFromUser(userAuthData, userProfile, userRole)
	sess.Set("user", sessionData)
	if err := sess.Save(); err != nil {
		return err
	}
	profileResponse := h.dtoManager.UserManager().ToUserProfile(sessionData, statistic, userRewards, sessionData.Streak)

	return response.Response(200, "Login successful", profileResponse)
}

// @Tags Auth
// @Summary Register
// @Description Register
// @Accept json
// @Produce json
// @Param register body dto.UserRegisterDTO true "Register"
// @Success 200 {object} response.BaseResponse{}
// @Router /public/register [post]
func (h *PublicHandler) Register(c *fiber.Ctx) error {
	var register dto.UserRegisterDTO
	if err := c.BodyParser(&register); err != nil {
		return err
	}
	if err := h.services.UtilService().Validator().ValidateStruct(register); err != nil {
		return err
	}

	defaultRole, err := h.services.RoleService().GetByName(c.Context(), h.config.Defaults.Roles.DefaultRole)
	if err != nil {
		return err
	}

	if err := h.services.UserService().Register(
		c.Context(),
		register.Username,
		register.Email,
		register.Password,
		register.ConfirmPassword,
		register.Name,
		register.Email,
		defaultRole.ID,
	); err != nil {
		return err
	}
	return response.Response(200, "Register successful", nil)
}

// @Tags Auth
// @Summary Logout
// @Description Logout
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /public/logout [post]
func (h *PublicHandler) Logout(c *fiber.Ctx) error {
	session, err := h.sessionStore.Get(c)
	if err != nil {
		return response.Response(500, "Failed to get session", err)
	}
	if err := session.Destroy(); err != nil {
		return response.Response(500, "Failed to destroy session", err)
	}

	return response.Response(200, "Logout successful", nil)
}
