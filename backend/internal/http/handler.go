package http

import (
	"github.com/C-dexTeam/codex/docs"
	"github.com/C-dexTeam/codex/internal/config"
	"github.com/C-dexTeam/codex/internal/config/models"
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/sessionStore"
	v1 "github.com/C-dexTeam/codex/internal/http/v1"
	"github.com/C-dexTeam/codex/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Handler struct {
	services *services.Services
	config   *config.Config
}

func NewHandler(
	services *services.Services,
	config *config.Config,
) *Handler {
	return &Handler{
		services: services,
		config:   config,
	}
}

func (h *Handler) Init(devMode bool, RedisConfig *models.RedisConfig, middlewares ...func(*fiber.Ctx) error) *fiber.App {
	app := fiber.New()
	for i := range middlewares {
		app.Use(middlewares[i])
	}

	if devMode {
		docs.SwaggerInfo.Version = config.Version
		app.Get("/api/dev/*", swagger.New(swagger.Config{
			Title:                "Codex Backend",
			TryItOutEnabled:      true,
			PersistAuthorization: true,
		}))
	}

	app.Static("/api/uploads", "./uploads")

	root := app.Group("/api")
	sessionStore := sessionStore.NewSessionStore(RedisConfig)
	dtoManager := dto.CreateNewDTOManager()

	// init routes
	v1.NewV1Handler(h.services, dtoManager, h.config).Init(root, sessionStore)

	return app
}
