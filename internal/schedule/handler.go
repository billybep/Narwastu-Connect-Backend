package schedule

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetAll(c echo.Context) error {
	schedules, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, schedules)
}

func (h *Handler) GetLatest(c echo.Context) error {
	schedule, err := h.service.GetLatest()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "no schedule found"})
	}
	return c.JSON(http.StatusOK, schedule)
}

func (h *Handler) GetByDate(c echo.Context) error {
	dateStr := c.Param("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid date format (YYYY-MM-DD)"})
	}

	schedule, err := h.service.GetByDate(date)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "schedule not found"})
	}

	return c.JSON(http.StatusOK, schedule)
}

func (h *Handler) Create(c echo.Context) error {
	var req ServiceSchedule
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if err := h.service.Create(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":  "schedule created successfully",
		"schedule": req,
	})
}
