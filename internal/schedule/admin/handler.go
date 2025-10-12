package admin

import (
	"app/internal/schedule"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc}
}

// ✅ Create Schedule (Admin only)
func (h *Handler) CreateSchedule(c echo.Context) error {
	var payload schedule.ServiceSchedule
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}

	if err := h.svc.CreateSchedule(&payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "schedule created successfully"})
}

// ✅ Get All Schedules (All roles)
func (h *Handler) GetSchedules(c echo.Context) error {
	schedules, err := h.svc.GetSchedules()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, schedules)
}

// ✅ Get Schedule by ID (All roles)
func (h *Handler) GetScheduleByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid schedule id"})
	}

	schedule, err := h.svc.GetScheduleByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "schedule not found"})
	}

	return c.JSON(http.StatusOK, schedule)
}

// ✅ === UPDATE ===
func (h *Handler) UpdateSchedule(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid schedule id"})
	}

	var payload schedule.ServiceSchedule
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}

	if err := h.svc.UpdateSchedule(uint(id), &payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "schedule updated successfully"})
}

// ✅ === SOFT DELETE ===
func (h *Handler) DeleteSchedule(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid schedule id"})
	}

	if err := h.svc.DeleteSchedule(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "schedule deleted successfully"})
}
