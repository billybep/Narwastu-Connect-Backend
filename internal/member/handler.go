package member

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
