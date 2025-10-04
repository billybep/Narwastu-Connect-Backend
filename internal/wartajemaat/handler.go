package wartajemaat

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetLatestWarta(c echo.Context) error {
	warta, err := h.service.GetLatest()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "no warta jemaat found",
		})
	}

	return c.JSON(http.StatusOK, warta)
}

func (h *Handler) UploadWartaJemaat(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "file is required"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to open file"})
	}
	defer src.Close()

	record, err := h.service.UploadWarta(src, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Warta jemaat uploaded successfully",
		"data":    record,
	})
}
