package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

type AuthService struct {
	repo   *UserRepository
	config *oauth2.Config
}

func NewAuthService(repo *UserRepository, cfg *oauth2.Config) *AuthService {
	return &AuthService{repo: repo, config: cfg}
}

func (s *AuthService) GenerateJWT(u *User) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	claims := jwt.MapClaims{
		"user_id":  u.ID,
		"email":    u.Email,
		"provider": u.Provider,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// Exchange code -> token -> userinfo, then upsert
func (s *AuthService) HandleGoogleCallback(code string) (*User, error) {
	tok, err := s.config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tok.AccessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var info GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	if info.ID == "" {
		return nil, errors.New("google user id empty")
	}

	// find or create
	u, err := s.repo.FindByProvider(info.ID, "google")
	if err != nil {
		// create
		email := info.Email
		name := info.Name
		avatar := info.Picture
		u = &User{
			Provider:   "google",
			ProviderID: info.ID,
			Email:      &email,
			Name:       &name,
			AvatarURL:  &avatar,
		}
		if err := s.repo.Create(u); err != nil {
			return nil, err
		}
	} else {
		// update minimal fields
		if info.Email != "" {
			u.Email = &info.Email
		}
		if info.Name != "" {
			u.Name = &info.Name
		}
		if info.Picture != "" {
			u.AvatarURL = &info.Picture
		}
		_ = s.repo.Save(u)
	}
	return u, nil
}
