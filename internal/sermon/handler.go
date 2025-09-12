package sermon

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

// GET /articles/weekly
func (h *Handler) GetWeeklyArticles(c echo.Context) error {
	articles, err := h.service.GetWeeklyArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Hanya kirim ringkasan
	var response []map[string]interface{}
	for _, s := range articles {
		response = append(response, map[string]interface{}{
			"id":            s.ID,
			"title":         s.Title,
			"image":         s.ImageURL,
			"profile_image": s.ProfileImage,
			"preacher":      s.Preacher,
			"date":          s.Date.Format("2006-01-02"),
		})
	}
	return c.JSON(http.StatusOK, response)
}

// GET /articles/yearly
func (h *Handler) GetYearlyArticles(c echo.Context) error {
	currentYear := time.Now().Year()
	articles, err := h.service.GetYearlyArticles(currentYear)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Hanya kirim ringkasan
	var response []map[string]interface{}
	for _, s := range articles {
		response = append(response, map[string]interface{}{
			"id":            s.ID,
			"title":         s.Title,
			"image":         s.ImageURL,
			"profile_image": s.ProfileImage,
			"preacher":      s.Preacher,
			"date":          s.Date.Format("2006-01-02"),
		})
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetMonthlyArticles(c echo.Context) error {
	sermons, err := h.service.GetMonthlyArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Hanya kirim ringkasan
	var response []map[string]interface{}
	for _, s := range sermons {
		response = append(response, map[string]interface{}{
			"id":            s.ID,
			"title":         s.Title,
			"image":         s.ImageURL,
			"profile_image": s.ProfileImage,
			"preacher":      s.Preacher,
			"date":          s.Date.Format("2006-01-02"),
		})
	}
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetDetailArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	sermon, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "article not found"})
	}

	response := map[string]interface{}{
		"id":            sermon.ID,
		"title":         sermon.Title,
		"image":         sermon.ImageURL,
		"preacher":      sermon.Preacher,
		"date":          sermon.Date.Format("2006-01-02"),
		"main_verse":    sermon.MainVerse,
		"profile_image": sermon.ProfileImage,
		"content":       sermon.Content,
	}

	return c.JSON(http.StatusOK, response)
}
