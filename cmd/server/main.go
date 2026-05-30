package main

import (
	"fmt"
	"github.com/ZakSlinin/sber-practice-backend/internal/auth/handler"
	"github.com/ZakSlinin/sber-practice-backend/internal/auth/repository"
	"github.com/ZakSlinin/sber-practice-backend/internal/auth/service"
	handler2 "github.com/ZakSlinin/sber-practice-backend/internal/challenges/handler"
	repository3 "github.com/ZakSlinin/sber-practice-backend/internal/challenges/repository"
	service2 "github.com/ZakSlinin/sber-practice-backend/internal/challenges/service"
	repository2 "github.com/ZakSlinin/sber-practice-backend/internal/workspace/repository"
	"github.com/ZakSlinin/sber-practice-backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	pgUser := os.Getenv("DB_USER")
	pgPassword := os.Getenv("DB_PASSWORD")
	pgHost := os.Getenv("DB_HOST")
	pgPort := os.Getenv("DB_PORT")
	pgDatabase := os.Getenv("DB_NAME")

	fmt.Printf("user=%s host=%s port=%s db=%s\n", pgUser, pgHost, pgPort, pgDatabase)
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		pgUser, pgPassword, pgHost, pgPort, pgDatabase,
	)

	m, err := migrate.New("file://migrations", dbURL)

	if err != nil {
		log.Fatalf("failed to create migrations: %s", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrations: %s", err)
	}
	log.Printf("migrated!")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	authRepo := repository.NewAuthRepository(db)
	workspaceRepo := repository2.NewWorkspaceRepository(db)
	challengesRepo := repository3.NewPostgresChallengesRepo(db)

	authService := service.NewAuthService(authRepo, workspaceRepo)
	challengesService := service2.NewChallengesService(challengesRepo)

	authHandler := handler.NewAuthHandler(authService)
	challengesHandler := handler2.NewChallengesHandler(challengesService)

	r := gin.Default()

	// -----------CORS-----------

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	api := r.Group("/api")

	{
		api.POST("/login", authHandler.Login)
		api.POST("/register", authHandler.Register)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())

		admin := protected.Group("/admin")
		admin.Use(middleware.RequireRole("admin"))
		admin.POST("/create-challenge", challengesHandler.CreateChallenge)
	}

	r.Run(":8080")
}
