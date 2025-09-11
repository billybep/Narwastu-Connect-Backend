package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type AuthService struct {
	repo         *UserRepository
	config       *oauth2.Config
	firebaseAuth *auth.Client
}

func NewAuthService(repo *UserRepository, cfg *oauth2.Config, fb *auth.Client) *AuthService {
	return &AuthService{repo: repo, config: cfg, firebaseAuth: fb}
}

// init firebase client sekali di main.go
func NewFirebaseAuth() *auth.Client {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS")) // serviceAccountKey.json
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("failed to init firebase app: " + err.Error())
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		panic("failed to init firebase auth: " + err.Error())
	}
	return client
}

// ---- method baru di AuthService ----
func (s *AuthService) HandleFirebaseIDToken(idToken string) (*User, error) {
	ctx := context.Background()

	token, err := s.firebaseAuth.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, errors.New("invalid firebase id token")
	}

	userRec, err := s.firebaseAuth.GetUser(ctx, token.UID)
	if err != nil {
		return nil, err
	}

	email := userRec.Email
	name := userRec.DisplayName
	avatar := userRec.PhotoURL

	// 🔑 unify: cek berdasarkan email
	var u *User
	if email != "" {
		u, _ = s.repo.FindByEmail(email)
	}

	if u == nil {
		// buat user baru provider = google (unified)
		u = &User{
			Provider:   "google",
			ProviderID: token.UID,
			Email:      &email,
			Name:       &name,
			AvatarURL:  &avatar,
		}
		if err := s.repo.Create(u); err != nil {
			return nil, err
		}
	} else {
		// update data
		u.Provider = "google"
		u.ProviderID = token.UID
		if name != "" {
			u.Name = &name
		}
		if avatar != "" {
			u.AvatarURL = &avatar
		}
		_ = s.repo.Save(u)
	}

	return u, nil
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

// HandleGoogleIDToken dipakai untuk login via Flutter (google_sign_in)
func (s *AuthService) HandleGoogleIDToken(idToken string) (*User, error) {
	resp, err := http.Get("https://oauth2.googleapis.com/tokeninfo?id_token=" + idToken)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid google id token")
	}
	defer resp.Body.Close()

	var info struct {
		Sub     string `json:"sub"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	if info.Sub == "" {
		return nil, errors.New("google user sub empty")
	}

	// 🔑 unify: cek by email dulu
	var u *User
	if info.Email != "" {
		u, _ = s.repo.FindByEmail(info.Email)
	}

	if u == nil {
		u = &User{
			Provider:   "google",
			ProviderID: info.Sub,
			Email:      &info.Email,
			Name:       &info.Name,
			AvatarURL:  &info.Picture,
		}
		if err := s.repo.Create(u); err != nil {
			return nil, err
		}
	} else {
		u.ProviderID = info.Sub
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
