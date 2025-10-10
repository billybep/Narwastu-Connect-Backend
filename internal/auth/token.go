package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var adminSecret = []byte(os.Getenv("JWT_ADMIN_SECRET"))

type AdminClaims struct {
	AdminID uint   `json:"admin_id"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	jwt.RegisteredClaims
}

// ✅ Generate Token untuk Admin Dashboard
func GenerateAdminToken(id uint, email string, role string, ttl time.Duration) (string, error) {
	claims := AdminClaims{
		AdminID: id,
		Email:   email,
		Role:    role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			Subject:   "admin_dashboard",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(adminSecret)
}

// ✅ Parse Token
func ParseAdminToken(tokenStr string) (*AdminClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		return adminSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AdminClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
