package routes

import (
	"ginauth/src/controllers/auth"
	"ginauth/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/v1/auth/login", auth.LoginHandler)
	router.POST("/v1/auth/register", auth.RegisterHandler)
	protectedRoute := router.Group("/v1/auth")
	protectedRoute.Use(middleware.JwtAuthMiddleware())
	{
		protectedRoute.GET("/me", auth.ProfileHandler)
	}

	return router
}
