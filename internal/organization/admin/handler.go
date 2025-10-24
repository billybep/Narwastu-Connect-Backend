package admin

import (
	"app/internal/organization"
	supabase "app/internal/storage"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service        Service
	supabaseClient *supabase.Client
}

func NewHandler(svc Service, supa *supabase.Client) *Handler {
	return &Handler{service: svc, supabaseClient: supa}
}

// CREATE
func (h *Handler) CreateOrganization(c echo.Context) error {
	var org organization.Organization
	if err := c.Bind(&org); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	if err := h.service.Create(&org); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, org)
}

// READ ALL
func (h *Handler) GetOrganizations(c echo.Context) error {
	orgs, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, orgs)
}

// READ BY ID
func (h *Handler) GetOrganizationByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	org, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
	}
	return c.JSON(http.StatusOK, org)
}

// UPDATE
func (h *Handler) UpdateOrganization(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var org organization.Organization
	if err := c.Bind(&org); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	org.ID = uint(id)

	if err := h.service.Update(&org); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, org)
}

// === SOFT DELETE ===
func (h *Handler) DeleteOrganization(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid organization ID"})
	}

	if err := h.service.DeleteOrganization(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "organization deleted (soft)"})
}

// Upload Profile Picture
func (h *Handler) UploadProfilePic(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid ID"})
	}

	// ✅ Cek apakah data masih aktif
	org, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "organization not found"})
	}
	if org.DeletedAt.Valid {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "organization has been soft deleted"})
	}

	// ✅ Hapus foto lama dari Supabase (jika ada)
	if org.ProfilePic != "" {
		if err := h.supabaseClient.DeleteFile(org.ProfilePic); err != nil {
			fmt.Printf("[WARN] gagal hapus foto lama: %v\n", err)
		}
	}

	// ✅ Ambil file dari form
	file, fileHeader, err := c.Request().FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "failed to read file"})
	}
	defer file.Close()

	// ✅ Upload ke Supabase
	url, err := h.supabaseClient.UploadOrganizationPhoto(file, fileHeader, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// ✅ Simpan URL baru ke DB
	if err := h.service.UpdateProfilePic(uint(id), url); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "profile picture uploaded successfully",
		"profile_url": url,
	})
}
