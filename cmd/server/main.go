package main

import (
	"log"
	"os"

	"app/internal/auth"
	"app/internal/common"
	"app/internal/event"
	"app/internal/member"
	"app/internal/sermon"
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
		// User Auth
		&auth.User{},

		// Verse of the day
		&verse.Verse{},
		&verse.Comment{},
		&verse.VerseLike{},
		&event.Event{},

		// Data Jemaat
		&member.Site{},
		&member.Member{},
		&member.FamilyRelationship{},

		// Khotbah
		&sermon.Sermon{},
	); err != nil {
		log.Fatal(err)
	}

	// Init Firebase
	auth.InitFirebase()

	// Seeder
	// if err := member.SeedMembers(repository.DB); err != nil {
	// 	log.Fatal("Seeder error: ", err)
	// }
	if err := sermon.Seed(repository.DB); err != nil {
		log.Fatal("failed seeding sermons:", err)
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
	firebaseClient := auth.NewFirebaseAuth()
	svc := auth.NewAuthService(repo, googleCfg, firebaseClient)
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
	e.POST("/auth/mobile", h.MobileLogin)

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

	// Event routes
	eventRepo := event.NewRepository(repository.DB)
	eventSvc := event.NewService(eventRepo)
	eventH := event.NewHandler(eventSvc)
	eventH.RegisterRoutes(e, api)

	// Member (Weekly Birthday)
	memberRepo := member.NewRepository(repository.DB)
	memberSvc := member.NewService(memberRepo)
	memberH := member.NewHandler(memberSvc)

	// Sermon routes
	sermonRepo := sermon.NewRepository(repository.DB)
	sermonSvc := sermon.NewService(sermonRepo)
	sermonH := sermon.NewHandler(sermonSvc)

	// Public routes
	member.RegisterRoutes(e.Group(""), memberH)
	sermon.RegisterRoutes(e.Group(""), sermonH)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
