package admin

import (
	"net/http"
	"time"

	"app/internal/auth"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

// ✅ Login Admin Dashboard
func (h *Handler) Login(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payload"})
	}

	admin, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
	}

	// Gunakan helper dari internal/auth untuk generate token
	token, err := auth.GenerateAdminToken(admin.ID, admin.Email, string(admin.Role), 24*time.Hour)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"user":  admin,
	})
}

// ✅ Create Admin (role-based)
func (h *Handler) CreateAdmin(c echo.Context) error {
	var input CreateAdminInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	// Ambil full admin model yang diset oleh middleware autentikasi
	creatorVal := c.Get("admin")
	if creatorVal == nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthenticated"})
	}

	creator, ok := creatorVal.(*Admin)
	if !ok {
		// safety: unexpected type in context
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "invalid admin in context"})
	}

	newAdmin, err := h.service.CreateAdmin(input, creator)
	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "admin created successfully",
		"admin":   newAdmin,
	})
}

// ✅ Get All Admins (hanya System Admin)
func (h *Handler) GetAll(c echo.Context) error {
	admins, err := h.service.GetAllAdmins()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, admins)
}
