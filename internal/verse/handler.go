package verse

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type VerseHandler struct {
	svc *VerseService
}

func NewVerseHandler(svc *VerseService) *VerseHandler {
	return &VerseHandler{svc: svc}
}

// Standardized response struct
type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Create new verse
func (h *VerseHandler) CreateVerse(c echo.Context) error {
	var v Verse
	if err := c.Bind(&v); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: "invalid request"})
	}
	if v.VerseText == "" || v.VerseReference == "" {
		return c.JSON(http.StatusBadRequest, Response{Error: "verse text and reference required"})
	}
	if err := h.svc.CreateVerse(&v); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Error: "failed to create verse"})
	}
	return c.JSON(http.StatusCreated, Response{Message: "verse created", Data: v})
}

// Get latest verse
func (h *VerseHandler) GetLatestVerse(c echo.Context) error {
	v, err := h.svc.GetLatestVerse()
	if err != nil {
		return c.JSON(http.StatusNotFound, Response{Error: "no verse found"})
	}
	return c.JSON(http.StatusOK, Response{Data: v})
}

// Like a verse
func (h *VerseHandler) LikeVerse(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid verse id"})
	}

	claimsRaw := c.Get("claims")
	claims, ok := claimsRaw.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid claims"})
	}
	idRaw, ok := claims["user_id"]
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "id not in claims"})
	}
	userIDFloat, ok := idRaw.(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "id is invalid type"})
	}
	userID := uint(userIDFloat)

	if err := h.svc.LikeVerse(userID, uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "verse liked"})
}

// Share a verse
func (h *VerseHandler) ShareVerse(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: "invalid verse id"})
	}
	if err := h.svc.ShareVerse(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, Response{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, Response{Message: "verse shared"})
}

// Add comment
func (h *VerseHandler) AddComment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: "invalid verse id"})
	}

	var body struct {
		User    string `json:"user"`
		Content string `json:"content"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: "invalid request"})
	}
	if body.Content == "" {
		return c.JSON(http.StatusBadRequest, Response{Error: "comment content required"})
	}

	if err := h.svc.AddComment(uint(id), body.Content, body.User); err != nil {
		return c.JSON(http.StatusNotFound, Response{Error: err.Error()})
	}
	return c.JSON(http.StatusCreated, Response{Message: "comment added"})
}
