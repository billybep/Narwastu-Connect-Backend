package main

import (
	"fmt"
	"log"
	"os"

	"app/internal/auth"
	"app/internal/event"
	"app/internal/finance"
	"app/internal/member"
	myMiddleware "app/internal/middleware" // 🔥 pakai middleware custom kamu
	"app/internal/organization"
	"app/internal/schedule"
	"app/internal/sermon"
	supabase "app/internal/storage"
	"app/internal/verse"
	"app/pkg/repository"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	// Init DB
	repository.InitDB()
	if err := repository.DB.AutoMigrate(
		// User Auth
		&auth.User{},

		// Verse of the day
		&verse.Verse{},
		&verse.VerseLike{},
		&event.Event{},

		// Data Jemaat
		&member.Site{},
		&member.Member{},
		&member.FamilyRelationship{},

		// Khotbah
		&sermon.Sermon{},

		// Jadwal Melayani
		&schedule.ServiceSchedule{},

		// Keuangan / Finance
		&finance.Account{},
		&finance.Transaction{},
		&finance.Balance{},

		// Organization
		&organization.Organization{},
	); err != nil {
		log.Fatal(err)
	}

	// Supabase
	supabaseClient := supabase.NewClient(
		os.Getenv("SUPABASE_URL")+"/storage/v1",
		os.Getenv("SUPABASE_KEY"),
		os.Getenv("SUPABASE_BUCKET"),
	)
	log.Println("SUPABASE_URL =", os.Getenv("SUPABASE_URL"))
	log.Println("SUPABASE_KEY =", os.Getenv("SUPABASE_KEY"))
	log.Println("SUPABASE_BUCKET =", os.Getenv("SUPABASE_BUCKET"))

	fmt.Println("🔥 Supabase initialized")

	// supabaseURL := os.Getenv("SUPABASE_URL")
	// supabaseKey := os.Getenv("SUPABASE_KEY")
	// supabaseBucket := os.Getenv("SUPABASE_BUCKET")

	// supabaseClient := supabase.NewClient(supabaseURL, supabaseKey, supabaseBucket)

	// Init Firebase
	auth.InitFirebase()

	// Seeder
	// if err := member.SeedMembers(repository.DB); err != nil {
	// 	log.Fatal("Seeder error: ", err)
	// }
	// if err := sermon.Seed(repository.DB); err != nil {
	// 	log.Fatal("failed seeding sermons:", err)
	// }
	// if err := schedule.SeedSchedules(repository.DB); err != nil {
	// 	log.Fatal("failed seeding schedules:", err)
	// }
	// if err := finance.SeedFinance(repository.DB); err != nil {
	// 	log.Fatal("failed seeding finance:", err)
	// }
	// seed data
	// if err := organization.Seed(repository.DB); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("✅ Organization seeding done")

	// println("[Seeder] Done.")

	// Init Echo
	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
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
	// api := e.Group("/api", common.StrictJWTMiddleware)
	api := e.Group("/api", myMiddleware.JWTMiddleware(os.Getenv("JWT_SECRET")))

	api.DELETE("/auth/delete", h.DeleteMyAccount) // 🔥 delete account pakai JWT

	// Sample protected route
	api.GET("/me", func(c echo.Context) error {
		claims := c.Get("claims")
		return c.JSON(200, claims)
	})

	// Verse actions (protected)
	api.POST("/verse/:id/like", verseH.LikeVerse)
	api.POST("/verse/:id/share", verseH.ShareVerse)

	// Event routes
	eventRepo := event.NewRepository(repository.DB)
	eventSvc := event.NewService(eventRepo)
	eventH := event.NewHandler(eventSvc)
	eventH.RegisterRoutes(e, api)

	// Member (Weekly Birthday)

	memberRepo := member.NewRepository(repository.DB)
	memberSvc := member.NewService(memberRepo)
	// memberH := member.NewHandler(memberSvc)
	memberH := member.NewHandler(memberSvc, supabaseClient)

	// Sermon routes
	sermonRepo := sermon.NewRepository(repository.DB)
	sermonSvc := sermon.NewService(sermonRepo)
	sermonH := sermon.NewHandler(sermonSvc)

	// Services Schedule
	scheduleRepo := schedule.NewRepository(repository.DB)
	scheduleService := schedule.NewService(scheduleRepo)
	scheduleHandler := schedule.NewHandler(scheduleService)
	schedule.RegisterRoutes(e.Group(""), scheduleHandler)

	// Finance routes
	financeRepo := finance.NewRepository(repository.DB)
	financeSvc := finance.NewService(financeRepo)
	financeH := finance.NewHandler(financeSvc)
	finance.RegisterRoutes(api, financeH)

	// Organization
	organizationRepo := organization.NewRepository(repository.DB)
	organizationSvc := organization.NewService(organizationRepo)
	organizationH := organization.NewHandler(organizationSvc)
	organization.RegisterRoutes(api, organizationH)

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
