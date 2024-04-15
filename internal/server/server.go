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
	// user router.
	user := engine.Group("/user")
	{
		user.POST("/register", handler.Register)
		user.POST("/login", handler.Login)
		user.POST("/verifyEmail", handler.VerifyEmail)
	}
	// relation router.
	relation := engine.Group("/relation")
	{
		relation.POST("/addFriendReq", handler.AddFriendReq)
	}
	return engine
}
