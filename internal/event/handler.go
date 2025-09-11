package event

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(e *echo.Echo, api *echo.Group) {
	// Public event routes
	e.GET("/events", h.GetEvents)
	e.GET("/events/weekly", h.GetWeeklyEvents)
	e.GET("/events/:id", h.GetEvent)

	// Protected event routes (hanya bisa diakses lewat JWT middleware)
	api.POST("/events", h.CreateEvent)
	api.PUT("/events/:id", h.UpdateEvent)
	api.DELETE("/events/:id", h.DeleteEvent)
}

func (h *Handler) GetEvents(c echo.Context) error {
	events, err := h.service.GetEvents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, events)
}

func (h *Handler) GetEvent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	event, err := h.service.GetEventByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "event not found"})
	}
	return c.JSON(http.StatusOK, event)
}

func (h *Handler) GetWeeklyEvents(c echo.Context) error {
	events, err := h.service.GetWeeklyEvents(time.Now())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, events)
}

func (h *Handler) CreateEvent(c echo.Context) error {
	var event Event
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := h.service.CreateEvent(&event); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, event)
}

func (h *Handler) UpdateEvent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var event Event
	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	event.ID = uint(id)
	if err := h.service.UpdateEvent(&event); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, event)
}

func (h *Handler) DeleteEvent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteEvent(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
