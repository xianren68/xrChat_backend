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
	// need to authentication.
	auth := engine.Group("/auth")
	auth.Use(middleware.Jwt())
	option := auth.Group("/option")
	{
		option.POST("updateLine", handler.UpdateLine)
		option.POST("updateName", handler.UpdateName)
		option.POST("updateGender", handler.UpdateGender)
		option.POST("updatePhone", handler.UpdatePhone)
	}
	// relation router.
	relation := auth.Group("/relation")
	{
		relation.POST("/addFriendReq", handler.AddFriendReq)
		relation.POST("/addFriendRes", handler.AddFriendRes)
		relation.POST("/createGroup", handler.CreateGroup)
		relation.POST("joinGroupReq", handler.JoinGroupReq)
		relation.POST("joinGroupRes", handler.JoinGroupRes)
		relation.POST("delFriend", handler.DelFriend)
		relation.POST("kickOutGroup", handler.KickOutGroup)
		relation.POST("quitGroup", handler.QuitGroup)

	}
	return engine
}
