package admin

import (
	adminmiddleware "app/internal/admin"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, h *Handler) {
	orgGroup := g.Group(
		"/organization",
		adminmiddleware.AdminJWTMiddleware,
		adminmiddleware.RequireAdminRole("system_administrator", "administrator"),
	)

	orgGroup.POST("", h.CreateOrganization)
	orgGroup.GET("", h.GetOrganizations)
	orgGroup.GET("/:id", h.GetOrganizationByID)
	orgGroup.PUT("/:id", h.UpdateOrganization)
	orgGroup.DELETE("/:id", h.DeleteOrganization)
	orgGroup.POST("/:id/upload", h.UploadProfilePic)
}
