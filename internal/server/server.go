// Package server supply Http server.
package server

import (
	"xrChat_backend/internal/handler"
	"xrChat_backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter init router.
func InitRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.Cors())
	// user router.
	user := engine.Group("/user")
	{
		user.POST("/register", handler.Register)
		user.POST("/login", handler.Login)
		user.POST("/verifyEmail", handler.VerifyEmail)
	}
	auth := engine.Group("/auth")
	auth.Use(middleware.Jwt())
	option := auth.Group("/option")
	{
		option.POST("updateLine", handler.UpdateLine)
	}
	// relation router.
	relation := auth.Group("/relation")
	{
		relation.POST("/addFriendReq", handler.AddFriendReq)
	}
	return engine
}
