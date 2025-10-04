package wartajemaat

import "github.com/labstack/echo/v4"

func RegisterRoutes(api *echo.Group, h *Handler) {
	api.POST("/warta/upload", h.UploadWartaJemaat)
	api.GET("/warta/latest", h.GetLatestWarta)
	// (Optional) api.GET("/warta", h.ListWartaJemaat)
}
