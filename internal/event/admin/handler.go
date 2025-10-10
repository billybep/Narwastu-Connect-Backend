package admin

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

type eventInput struct {
	Title       string    `json:"title"`
	DateTime    time.Time `json:"dateTime"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	ImageURL    string    `json:"imageUrl"`
}

func (h *Handler) CreateEvent(c echo.Context) error {
	var input eventInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}

	if err := h.service.CreateEvent(input.Title, input.Location, input.Description, input.ImageURL, input.DateTime); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "event created successfully"})
}

func (h *Handler) GetAllEvents(c echo.Context) error {
	events, err := h.service.GetAllEvents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, events)
}

func (h *Handler) GetEventByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid event id"})
	}

	event, err := h.service.GetEventByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "event not found"})
	}

	return c.JSON(http.StatusOK, event)
}

func (h *Handler) UpdateEvent(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid event id"})
	}

	var input eventInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}

	if err := h.service.UpdateEvent(uint(id), input.Title, input.Location, input.Description, input.ImageURL, input.DateTime); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "event updated successfully"})
}

func (h *Handler) DeleteEvent(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid event id"})
	}

	if err := h.service.DeleteEvent(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "event deleted successfully"})
}
