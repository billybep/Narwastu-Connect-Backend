package main

import (
	"fmt"
	"log"
	"os"

	"app/internal/admin"
	"app/internal/auth"
	"app/internal/event"
	adminevent "app/internal/event/admin"
	"app/internal/finance"
	adminfinance "app/internal/finance/admin"
	"app/internal/member"
	"app/internal/organization"
	adminorganization "app/internal/organization/admin"
	"app/internal/schedule"
	adminschedule "app/internal/schedule/admin"
	"app/internal/sermon"
	"app/internal/verse"
	"app/internal/wartajemaat"
	"app/pkg/repository"

	myMiddleware "app/internal/middleware" // 🔥 pakai middleware custom kamu
	supabase "app/internal/storage"

	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
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

		// WartaJemaat
		&wartajemaat.WartaJemaat{},

		/**
		#### ADMIN ------------------
		*/
		&admin.Admin{},
		/* -------------------------- */
	); err != nil {
		log.Fatal(err)
	}

	// Supabase
	supabaseClient := supabase.NewClient(
		os.Getenv("SUPABASE_URL")+"/storage/v1",
		os.Getenv("SUPABASE_KEY"),
		os.Getenv("SUPABASE_BUCKET"),
	)

	fmt.Println("🔥 Supabase initialized")

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

	// admin.Seed(repository.DB)
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

	// ========================
	// 📰 Warta Jemaat Routes
	// ========================
	wartaService := wartajemaat.NewService(repository.DB, supabaseClient)
	wartaHandler := wartajemaat.NewHandler(wartaService)
	wartajemaat.RegisterRoutes(e.Group(""), wartaHandler)
	// wartajemaat.RegisterRoutes(api, wartaHandler) private

	// Public routes
	member.RegisterRoutes(e.Group(""), memberH)
	sermon.RegisterRoutes(e.Group(""), sermonH)

	/**
	ADMIN ROUTES ---------------------------------------------------------------
	*/
	adminRepo := admin.NewRepository(repository.DB)
	adminSvc := admin.NewService(adminRepo)
	adminH := admin.NewHandler(adminSvc)

	// Untuk admin kita pisahkan prefix /admin
	// dan group ini TIDAK menggunakan JWT global, karena login pakai 2FA dulu
	adminGroup := e.Group("/admin")
	adminGroup.POST("/login", adminH.Login)

	admin.RegisterRoutes(adminGroup, adminH)

	// =========================
	// 🎉 Event Admin Routes
	// =========================
	eventAdminRepo := adminevent.NewRepository(repository.DB)
	eventAdminService := adminevent.NewService(eventAdminRepo)
	eventAdminHandler := adminevent.NewHandler(eventAdminService)
	adminevent.RegisterRoutes(adminGroup, eventAdminHandler)

	// =========================
	// 📊 Finance Admin Routes
	// =========================
	financeAdminService := adminfinance.NewService(financeRepo)
	financeAdminHandler := adminfinance.NewHandler(financeAdminService)
	adminfinance.RegisterRoutes(adminGroup, financeAdminHandler)

	// =========================
	// 📅 Schedule Admin Routes
	// =========================
	scheduleAdminRepo := adminschedule.NewRepository(repository.DB)
	scheduleAdminService := adminschedule.NewService(scheduleAdminRepo)
	scheduleAdminHandler := adminschedule.NewHandler(scheduleAdminService)
	adminschedule.RegisterRoutes(adminGroup, scheduleAdminHandler)

	// =========================
	// 🏢 Organization Admin Routes
	// =========================
	organizationAdminRepo := adminorganization.NewRepository(repository.DB)
	organizationAdminService := adminorganization.NewService(organizationAdminRepo)
	organizationAdminHandler := adminorganization.NewHandler(organizationAdminService, supabaseClient)
	adminorganization.RegisterRoutes(adminGroup, organizationAdminHandler)

	/* ------------------------------------------------------------------------- */

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
