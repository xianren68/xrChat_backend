// Package handler handle network requests.
package handler

import (
	"xrChat_backend/internal/proto/pb"
	"xrChat_backend/internal/service"

	"xrChat_backend/pkg"

	"github.com/gin-gonic/gin"
)

// Login handle user login.
func Login(c *gin.Context) {
	loginInfo := &pb.LoginRequest{}
	err := pkg.BindProto(c, loginInfo)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	err = service.Login(loginInfo)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	pkg.HandleSuccess(c, "登录成功")

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
	}
	pkg.HandleSuccess(c, "注册成功，userId已发送至邮箱")
}
