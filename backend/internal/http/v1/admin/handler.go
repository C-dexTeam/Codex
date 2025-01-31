package admin

import (
	"fmt"

	"github.com/C-dexTeam/codex/internal/config"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/internal/http/sessionStore"
	"github.com/C-dexTeam/codex/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type AdminHandler struct {
	services   *services.Services
	sess_store *session.Store
	dtoManager dto.IDTOManager
	config     *config.Config
}

func NewAdminHandler(
	service *services.Services,
	sessStore *session.Store,
	dtoManager dto.IDTOManager,
	config *config.Config,
) *AdminHandler {
	return &AdminHandler{
		services:   service,
		sess_store: sessStore,
		dtoManager: dtoManager,
		config:     config,
	}
}

func (h *AdminHandler) Init(router fiber.Router) {
	root := router.Group("/admin")
	root.Use(h.authMiddleware)
	root.Use(h.adminRoleMiddleware)

	root.Get("/", func(c *fiber.Ctx) error {
		data := sessionStore.GetSessionData(c)
		return response.Response(200, fmt.Sprintf("Dear %s %s Welcome to Codex API (Private Admin Zone)", data.Name, data.Surname), nil)
	})

	// Initialize Routes
	h.initUserRoutes(root)
	h.initRewardsRoutes(root)
	h.initTestsRoutes(root)
	h.initProgrammingLanguageRoutes(root)
	h.initCoursesRoutes(root)
	h.initChaptersRoutes(root)
	h.initAttributesRoutes(root)
}

func (h *AdminHandler) authMiddleware(c *fiber.Ctx) error {
	session, err := h.sess_store.Get(c)
	if err != nil {
		return err
	}
	user := session.Get("user")
	if user == nil {
		return serviceErrors.NewServiceErrorWithMessage(401, "unauthorized")
	}
	session_data, ok := user.(sessionStore.SessionData)
	if !ok {
		return serviceErrors.NewServiceErrorWithMessage(500, "session data error")
	}
	if session_data.Role == "Banned" {
		return serviceErrors.NewServiceErrorWithMessage(403, "Banned")
	}
	c.Locals("user", session_data)

	return c.Next()
}

func (h *AdminHandler) adminRoleMiddleware(c *fiber.Ctx) error {
	session, err := h.sess_store.Get(c)
	if err != nil {
		return err
	}
	user := session.Get("user")
	if user == nil {
		return serviceErrors.NewServiceErrorWithMessage(401, "unauthorized")
	}
	session_data, ok := user.(sessionStore.SessionData)
	if !ok {
		return serviceErrors.NewServiceErrorWithMessage(500, "session data error")
	}

	if session_data.Role == "Banned" {
		return serviceErrors.NewServiceErrorWithMessage(403, "Banned")
	}
	if session_data.Role != "admin" {
		return serviceErrors.NewServiceErrorWithMessage(403, "unauthorized!")
	}
	c.Locals("user", session_data)

	return c.Next()
}
