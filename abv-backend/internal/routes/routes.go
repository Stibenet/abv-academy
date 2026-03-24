package routes

import (
	"abv-backend/internal/handlers"
	"abv-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(
	r *gin.Engine,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	jwtSecret string,
) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	users := r.Group("/users")
	users.Use(middleware.AuthMiddleware(jwtSecret))
	{
		users.GET("", middleware.RoleMiddleware("administrator", "director"), userHandler.GetUsers)
		users.GET("/:id", middleware.RoleMiddleware("administrator", "director", "parent", "child", "guest"), userHandler.GetUser)
		users.POST("", middleware.RoleMiddleware("administrator", "director"), userHandler.CreateUser)
		users.PUT("/:id", middleware.RoleMiddleware("administrator", "director"), userHandler.UpdateUser)
		users.DELETE("/:id", middleware.RoleMiddleware("administrator"), userHandler.DeleteUser)
	}
}
