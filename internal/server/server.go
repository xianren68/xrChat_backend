// Package server supply Http server.
package server

import (
	"github.com/gin-gonic/gin"
	"xrChat_backend/internal/handler"
	"xrChat_backend/internal/middleware"
)

// InitRouter init router.
func InitRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.Cors())
	user := engine.Group("/user")
	{
		user.POST("/register", handler.Register)
		user.POST("/login", handler.Login)
		user.POST("/verifyEmail", handler.VerifyEmail)
	}
	return engine
}
