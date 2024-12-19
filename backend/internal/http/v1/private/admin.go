package private

import (
	"github.com/gofiber/fiber/v2"
)

func (h *PrivateHandler) initAdminRoutes(root fiber.Router) {
	adminRoute := root.Group("/admin")
	adminRoute.Use(h.adminRoleMiddleware)
	adminRoute.Get("/user", h.GetUsers)

}
