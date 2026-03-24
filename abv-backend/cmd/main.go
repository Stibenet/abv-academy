package main

import (
	"fmt"

	"abv-backend/docs"
	"abv-backend/internal/config"
	"abv-backend/internal/db"
	"abv-backend/internal/handlers"
	"abv-backend/internal/repository"
	"abv-backend/internal/routes"
	"abv-backend/internal/services"

	"github.com/gin-gonic/gin"
)

// @title Go CRUD API
// @version 1.0
// @description CRUD REST API на Go с PostgreSQL, JWT, ролями и Swagger
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.Load()
	database := db.NewPostgres(cfg)

	userRepo := repository.NewUserRepository(database)
	authService := services.NewAuthService(userRepo, cfg.JWTSecret, cfg.JWTTTLHours)
	userService := services.NewUserService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.AppPort)

	r := gin.Default()
	routes.RegisterRoutes(r, authHandler, userHandler, cfg.JWTSecret)

	_ = r.Run(":" + cfg.AppPort)
}
