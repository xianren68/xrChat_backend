// Package handler handle network requests.
package handler

import (
	"xrChat_backend/internal/middleware"
	"xrChat_backend/internal/proto/pb"
	"xrChat_backend/internal/service"

	"xrChat_backend/pkg"

	"github.com/gin-gonic/gin"
)

// Login handle user login.
func Login(c *gin.Context) {
	loginInfo := &pb.LoginRequest{}
	userProto := &pb.LoginResponse{}
	err := pkg.BindProto(c, loginInfo)
	if err != nil {
		userProto.Code = 500
		userProto.Message = err.Error()
		pkg.WriteProto(c, userProto)
		return
	}
	err, userInfo := service.Login(loginInfo)
	if err != nil {
		userProto.Code = 500
		userProto.Message = err.Error()
		pkg.WriteProto(c, userProto)
		return
	}
	token, err := middleware.GenerateToken(uint(userInfo.ID))
	if err != nil {
		userProto.Code = 500
		userProto.Message = err.Error()
		pkg.WriteProto(c, userProto)
		return
	}
	userProto.Id = uint64(userInfo.ID)
	userProto.Username = userInfo.Username
	userProto.Line = userInfo.Line
	userProto.Avatar = userInfo.Avatar
	userProto.Code = 200
	userProto.Message = "success"
	userProto.Email = userInfo.Email
	userProto.Token = token
	userProto.Gender = userInfo.Gender
	userProto.Phone = userInfo.Phone

	pkg.WriteProto(c, userProto)

}

// VerifyEmail verify email.
func VerifyEmail(c *gin.Context) {
	emailMsg := &pb.EmailVerifyRequest{}
	err := pkg.BindProto(c, emailMsg)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	err = service.VerifyEmail(emailMsg.Email)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	pkg.HandleSuccess(c, "验证码已发送,请注意查收")
}

// Register handle user register.
func Register(c *gin.Context) {
	registerInfo := &pb.RegisterRequest{}
	err := pkg.BindProto(c, registerInfo)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	err = service.Register(registerInfo)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	pkg.HandleSuccess(c, "注册成功，邮箱账号用于登录")
}
