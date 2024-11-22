package public

import (
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
}

func NewPublicHandler(
	service *services.Services,
	sessionStore *session.Store,
	dtoManager dto.IDTOManager,
) *PublicHandler {
	return &PublicHandler{
		services:     service,
		sessionStore: sessionStore,
		dtoManager:   dtoManager,
	}
}

func (h *PublicHandler) Init(router fiber.Router) {
	root := router.Group("/public")

	root.Get("/", func(c *fiber.Ctx) error {
		return response.Response(200, "Welcome to <github.com/C-dexTeam/codex> API (Public Zone)", nil)
	})

	h.initUserRoutes(root)
}
