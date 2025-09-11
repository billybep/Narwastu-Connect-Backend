package auth

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler(s *AuthService) *AuthHandler { return &AuthHandler{service: s} }

// GET /auth/google/login
func (h *AuthHandler) GoogleLogin(c echo.Context) error {
	state := RandState(16)
	c.SetCookie(&http.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(10 * time.Minute),
	})
	url := h.service.config.AuthCodeURL(state)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// GET /auth/google/callback?state=...&code=...
func (h *AuthHandler) GoogleCallback(c echo.Context) error {
	state := c.QueryParam("state")
	code := c.QueryParam("code")
	ck, _ := c.Cookie("oauthstate")
	if ck == nil || ck.Value != state {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid oauth state"})
	}

	u, err := h.service.HandleGoogleCallback(code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	jwt, err := h.service.GenerateJWT(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "token error"})
	}
	return c.JSON(http.StatusOK, map[string]any{
		"token": jwt,
		"user":  u,
	})
}

// POST /auth/mobile
func (h *AuthHandler) MobileLogin(c echo.Context) error {
	var body struct {
		Provider string `json:"provider"` // "google" | "firebase"
		IDToken  string `json:"id_token"`
	}
	if err := c.Bind(&body); err != nil || body.IDToken == "" || body.Provider == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "provider and id_token required"})
	}

	var u *User
	var err error

	switch body.Provider {
	case "google":
		u, err = h.service.HandleGoogleIDToken(body.IDToken)
	case "firebase":
		u, err = h.service.HandleFirebaseIDToken(body.IDToken)
	default:
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "unsupported provider"})
	}

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	jwt, err := h.service.GenerateJWT(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "token error"})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"token": jwt,
		"user":  u,
	})
}

// POST /auth/signout (stateless: client hapus token)
func (h *AuthHandler) SignOut(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
