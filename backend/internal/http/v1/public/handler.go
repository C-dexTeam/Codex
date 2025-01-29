package public

import (
	"github.com/C-dexTeam/codex/internal/config"
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type PublicHandler struct {
	services     *services.Services
	sessionStore *session.Store
	dtoManager   dto.IDTOManager
	config       *config.Config
}

func NewPublicHandler(
	service *services.Services,
	sessionStore *session.Store,
	dtoManager dto.IDTOManager,
	config *config.Config,
) *PublicHandler {
	return &PublicHandler{
		services:     service,
		sessionStore: sessionStore,
		dtoManager:   dtoManager,
		config:       config,
	}
}

func (h *PublicHandler) Init(router fiber.Router) {
	root := router.Group("/public")

	root.Get("/", func(c *fiber.Ctx) error {
		return response.Response(200, "Welcome to Codex API (Public Zone)", nil)
	})

	h.initUserRoutes(root)
	h.initRewardsRoutes(root)
}
