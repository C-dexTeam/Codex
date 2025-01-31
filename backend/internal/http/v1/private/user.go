package private

import (
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
	user.Post("/connect", h.ConnectWallet)
	user.Post("/streak", h.StreakUp)
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

	statistic, err := h.services.UserProfileService().UserStatistic(c.Context(), userSession.UserID)
	if err != nil {
		return err
	}
	userRewards, err := h.services.RewardService().GetUserRewards(c.Context(), userSession.UserID, "", "")
	if err != nil {
		return err
	}

	userProfileDTO := h.dtoManager.UserManager().ToUserProfile(*userSession, statistic, userRewards, userSession.Streak)

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
	walletUser, err := h.services.RoleService().GetByName(c.Context(), h.config.Defaults.Roles.RoleWalletUser)
	if err != nil {
		return err
	}

	// If the user in user Role. Change The Users role to wallet-user.
	if userSession.Role != h.config.Defaults.Roles.RoleAdmin {
		if err := h.services.UserProfileService().ChangeUserRole(c.Context(), userSession.UserProfileID, walletUser.ID.String()); err != nil {
			return err
		}
	}

	return response.Response(200, "Status OK", nil)
}

// @Tags User
// @Summary Streak Up
// @Description + your streak and gain exp.
// @Accept json
// @Produce json
// @Success 200 {object} response.BaseResponse{}
// @Router /private/user/streak [post]
func (h *PrivateHandler) StreakUp(c *fiber.Ctx) error {
	userSession := sessionStore.GetSessionData(c)

	userProfile, err := h.services.UserProfileService().GetUser(c.Context(), userSession.UserProfileID)
	if err != nil {
		return err
	}

	streak, err := h.services.UserProfileService().StreakUp(c.Context(), userSession.UserProfileID, userProfile.LastStreakDate.Time)
	if err != nil {
		return err
	}

	out, err := h.services.UserProfileService().AddUserExp(c.Context(), userSession.UserProfileID, domains.StreakExp)
	if err != nil {
		return err
	}

	sess, err := h.sess_store.Get(c)
	if err != nil {
		return err
	}
	userSession.SetStreak(int(streak))
	userSession.SetLevel(out.Level.Int32, out.Experience.Int32, out.NextLevelExp.Int32)
	sess.Set("user", userSession)
	if err := sess.Save(); err != nil {
		return err
	}

	return response.Response(200, "Status OK", streak)
}
