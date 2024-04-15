// Package service Support business logic operations.
package service

import (
	"errors"
	"fmt"
	"xrChat_backend/internal/model"
	"xrChat_backend/internal/proto/pb"
	"xrChat_backend/internal/repository"
	"xrChat_backend/pkg"
)

func Login(loginInfo *pb.LoginRequest) error {
	user := &model.User{}
	user.Email = loginInfo.Email
	user.Password = loginInfo.Password
	return repository.Login(user)

}

func VerifyEmail(email string) (err error) {
	err = pkg.VerifyEmailAddress(email)
	if err != nil {
		err = errors.New("邮箱地址有误")
		return
	}
	code := pkg.GenValidateCode(6)
	err = repository.VerifyEmail(email, code)
	if err != nil {
		err = errors.New("生成验证码出错")
		return
	}
	return nil
}

func Register(registerInfo *pb.RegisterRequest) error {
	user := &model.User{}
	user.Username = registerInfo.Username
	// generate salt
	salt, err := pkg.GenerateSalt()
	fmt.Println(salt)
	if err != nil {
		return err
	}
	user.Salt = salt
	user.Password = pkg.EncryptPassword(salt, registerInfo.Password)
	user.Email = registerInfo.Email
	err = repository.VerifyEmailCode(registerInfo.Email, registerInfo.Code)
	if err != nil {
		return err
	}
	return repository.RegisterUser(user)
}
