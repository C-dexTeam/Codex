package v1

import (
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	"github.com/C-dexTeam/codex/internal/http/v1/private"
	"github.com/C-dexTeam/codex/internal/http/v1/public"
	"github.com/C-dexTeam/codex/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type V1Handler struct {
	services   *services.Services
	dtoManager dto.IDTOManager
}

func NewV1Handler(services *services.Services, dtoManager dto.IDTOManager) *V1Handler {
	return &V1Handler{
		services:   services,
		dtoManager: dtoManager,
	}
}

func (h *V1Handler) Init(router fiber.Router, sessionStore *session.Store) {
	root := router.Group("/v1")
	root.Get("/", func(c *fiber.Ctx) error {
		return response.Response(200, "Welcome to Codex API (Root Zone)", nil)
	})

	// Init Fiber Session Store
	public.NewPublicHandler(h.services, sessionStore, h.dtoManager).Init(root)
	private.NewPrivateHandler(h.services, sessionStore, h.dtoManager).Init(root)
}
