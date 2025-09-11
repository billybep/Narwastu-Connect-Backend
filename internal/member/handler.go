package member

import (
	"net/http"

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
	// var response []map[string]interface{}
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
