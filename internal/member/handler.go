package member

import (
	supabase "app/internal/storage"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service        Service
	supabaseClient *supabase.Client
}

// func NewHandler(s Service) *Handler {
// 	return &Handler{service: s}
// }

func NewHandler(s Service, supabaseClient *supabase.Client) *Handler {
	return &Handler{
		service:        s,
		supabaseClient: supabaseClient,
	}
}

func (h *Handler) GetWeeklyBirthdays(c echo.Context) error {
	members, err := h.service.GetWeeklyBirthdays()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// mapping ke format frontend
	response := make([]map[string]interface{}, 0)
	for _, m := range members {
		response = append(response, map[string]interface{}{
			"name":       m.FullName,
			"birthDate":  m.DateOfBirth.Format("2006-01-02"),
			"familyName": m.FamilyName,
			"site":       m.Site.Name,
			"photoUrl":   m.PhotoURL,
		})
	}

	return c.JSON(http.StatusOK, response)
}

// Create Member
func (h *Handler) CreateMember(c echo.Context) error {
	var req struct {
		FullName    string `json:"fullName" validate:"required"`
		DateOfBirth string `json:"dateOfBirth" validate:"required"` // format: 2006-01-02
		FamilyName  string `json:"familyName"`
		SiteID      uint   `json:"siteId" validate:"required"`
		Gender      string `json:"gender"`
		Address     string `json:"address"`
		Phone       string `json:"phone"`
		PhotoURL    string `json:"photoUrl"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// parse tanggal
	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid date format"})
	}

	member := &Member{
		FullName:    req.FullName,
		DateOfBirth: dob,
		FamilyName:  req.FamilyName,
		SiteID:      req.SiteID,
		Gender:      req.Gender,
		Address:     req.Address,
		Phone:       req.Phone,
		PhotoURL:    req.PhotoURL,
	}

	if err := h.service.CreateMember(member); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "member created successfully",
		"member":  member,
	})
}

// Get List Members
func (h *Handler) GetMembers(c echo.Context) error {
	members, err := h.service.GetAllMembers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// mapping respons ke format frontend (optional, bisa juga langsung kirim struct)
	response := make([]map[string]interface{}, 0)
	for _, m := range members {
		response = append(response, map[string]interface{}{
			"id":          m.ID,
			"fullName":    m.FullName,
			"email":       m.Email,
			"phone":       m.Phone,
			"site":        m.Site.Name,
			"photoUrl":    m.PhotoURL,
			"isActive":    m.IsActive,
			"baptismDate": m.BaptismDate,
		})
	}

	return c.JSON(http.StatusOK, response)
}

// Get Detail Member
func (h *Handler) GetMemberByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid member ID"})
	}

	member, err := h.service.GetMemberByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "member not found"})
	}

	// mapping ke format respons
	response := map[string]interface{}{
		"id":            member.ID,
		"nik":           member.NIK,
		"fullName":      member.FullName,
		"familyName":    member.FamilyName,
		"placeOfBirth":  member.PlaceOfBirth,
		"dateOfBirth":   member.DateOfBirth.Format("2006-01-02"),
		"gender":        member.Gender,
		"address":       member.Address,
		"phone":         member.Phone,
		"email":         member.Email,
		"baptismDate":   member.BaptismDate,
		"site":          member.Site.Name,
		"maritalStatus": member.MaritalStatus,
		"photoUrl":      member.PhotoURL,
		"isActive":      member.IsActive,
	}

	return c.JSON(http.StatusOK, response)
}

// Updata Member (PUT)
func (h *Handler) UpdateMember(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid member ID"})
	}

	var req struct {
		FullName      string  `json:"fullName"`
		FamilyName    string  `json:"familyName"`
		PlaceOfBirth  string  `json:"placeOfBirth"`
		DateOfBirth   string  `json:"dateOfBirth"` // format YYYY-MM-DD
		Gender        string  `json:"gender"`
		Address       string  `json:"address"`
		Phone         string  `json:"phone"`
		Email         string  `json:"email"`
		BaptismDate   *string `json:"baptismDate"`
		SiteID        uint    `json:"siteId"`
		MaritalStatus string  `json:"maritalStatus"`
		PhotoURL      string  `json:"photoUrl"`
		IsActive      bool    `json:"isActive"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// parse tanggal lahir
	var dob time.Time
	if req.DateOfBirth != "" {
		dob, err = time.Parse("2006-01-02", req.DateOfBirth)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid dateOfBirth format"})
		}
	}

	// parse baptismDate
	var baptismDate *time.Time
	if req.BaptismDate != nil && *req.BaptismDate != "" {
		bd, err := time.Parse("2006-01-02", *req.BaptismDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid baptismDate format"})
		}
		baptismDate = &bd
	}

	updateData := &Member{
		FullName:      req.FullName,
		FamilyName:    req.FamilyName,
		PlaceOfBirth:  req.PlaceOfBirth,
		DateOfBirth:   dob,
		Gender:        req.Gender,
		Address:       req.Address,
		Phone:         req.Phone,
		Email:         req.Email,
		BaptismDate:   baptismDate,
		SiteID:        req.SiteID,
		MaritalStatus: req.MaritalStatus,
		PhotoURL:      req.PhotoURL,
		IsActive:      req.IsActive,
	}

	updatedMember, err := h.service.UpdateMember(uint(id), updateData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "member updated successfully",
		"member":  updatedMember,
	})
}

func (h *Handler) DeleteMember(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid member ID"})
	}

	if err := h.service.DeleteMember(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "member not found or already deleted"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "member deleted successfully (soft delete)",
	})
}

// Upload Profile Picture
func (h *Handler) UploadAvatar(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid member ID"})
	}

	// ✅ Cek apakah member masih aktif
	member, err := h.service.GetMemberByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "member not found"})
	}

	if member.DeletedAt.Valid {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "member has been soft deleted"})
	}

	file, fileHeader, err := c.Request().FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "failed to read file"})
	}
	defer file.Close()

	// ✅ Upload ke Supabase
	url, err := h.supabaseClient.UploadAvatar(file, fileHeader, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// ✅ Update database
	if err := h.service.UpdateMemberAvatar(uint(id), url); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "avatar uploaded successfully",
		"photo_url": url,
	})
}
