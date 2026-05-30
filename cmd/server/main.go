package main

import (
	"fmt"
	"github.com/ZakSlinin/sber-practice-backend/internal/auth/handler"
	"github.com/ZakSlinin/sber-practice-backend/internal/auth/repository"
	"github.com/ZakSlinin/sber-practice-backend/internal/auth/service"
	repository2 "github.com/ZakSlinin/sber-practice-backend/internal/workspace/repository"
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

	authService := service.NewAuthService(authRepo, workspaceRepo)

	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

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
	}

	r.Run(":8080")
}
