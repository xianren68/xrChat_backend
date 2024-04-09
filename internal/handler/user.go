// Package handler handle network requests.
package handler

import (
	"io"
	"xrChat_backend/internal/proto/pb"
	"xrChat_backend/internal/service"

	"xrChat_backend/pkg"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

// Login handle user login.
func Login(c *gin.Context) {
	loginInfo := &pb.LoginRequest{}
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		_, _ = c.Writer.Write(pkg.HandleError(err))
		return
	}
	err = proto.Unmarshal(bytes, loginInfo)
	if err != nil {
		_, _ = c.Writer.Write(pkg.HandleError(err))
		return
	}
	err = service.Login(loginInfo)
	if err != nil {
		_, _ = c.Writer.Write(pkg.HandleError(err))
		return
	}
	_, _ = c.Writer.Write(pkg.HandleSuccess("登录成功"))

}

// VerifyEmail verify email.
func VerifyEmail(c *gin.Context) {
	emailMsg := &pb.EmailVerifyRequest{}
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		_, _ = c.Writer.Write(pkg.HandleError(err))
		return
	}
	err = proto.Unmarshal(bytes, emailMsg)
	if err != nil {
		_, _ = c.Writer.Write(pkg.HandleError(err))
		return
	}
	err = service.VerifyEmail(emailMsg.Email)
	if err != nil {
		_, _ = c.Writer.Write(pkg.HandleError(err))
	}
	_, _ = c.Writer.Write(pkg.HandleSuccess("验证码已发送,请注意查收"))

}

// Register handle user register.
func Register(c *gin.Context) {
	registerInfo := &pb.RegisterRequest{}
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		_, _ = c.Writer.Write(pkg.HandleError(err))
	}
	err = proto.Unmarshal(bytes, registerInfo)
	if err != nil {
		_, _ = c.Writer.Write(pkg.HandleError(err))
	}
	err = service.Register(registerInfo)
	if err != nil {
		_, _ = c.Writer.Write(pkg.HandleError(err))
	}
}
