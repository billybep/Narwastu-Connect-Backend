package main

import (
	"log"
	"os"

	"app/internal/auth"
	"app/internal/common"
	"app/internal/verse"
	"app/pkg/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	println("[Fx Main] Called")

	// Init DB
	repository.InitDB()
	if err := repository.DB.AutoMigrate(
		&auth.User{},
		&verse.Verse{},
		&verse.Comment{},
		&verse.VerseLike{},
	); err != nil {
		log.Fatal(err)
	}

	// Init Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	// ========================
	// 📌 Public Routes
	// ========================
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	// Auth routes
	repo := auth.NewUserRepository()
	googleCfg := auth.NewGoogleOAuthConfig()
	svc := auth.NewAuthService(repo, googleCfg)
	h := auth.NewAuthHandler(svc)

	e.GET("/auth/google/login", h.GoogleLogin)
	e.GET("/auth/google/callback", h.GoogleCallback)
	e.POST("/auth/signout", h.SignOut)

	// Verse routes (public)
	verseRepo := verse.NewVerseRepository()
	verseSvc := verse.NewVerseService(verseRepo)
	verseH := verse.NewVerseHandler(verseSvc)

	e.GET("/verse/latest", verseH.GetLatestVerse)
	e.POST("/verse", verseH.CreateVerse)

	// ========================
	// 🔒 Protected Routes (/api/*)
	// ========================
	api := e.Group("/api", common.StrictJWTMiddleware)

	// Sample protected route
	api.GET("/me", func(c echo.Context) error {
		claims := c.Get("claims")
		return c.JSON(200, claims)
	})

	// Verse actions (protected)
	api.POST("/verse/:id/like", verseH.LikeVerse)
	api.POST("/verse/:id/share", verseH.ShareVerse)
	api.POST("/verse/:id/comment", verseH.AddComment)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
