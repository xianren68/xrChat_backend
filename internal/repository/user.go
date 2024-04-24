// Package repository  Provides interaction with the database.
package repository

import (
	"context"
	"errors"
	"log/slog"
	"time"
	"xrChat_backend/config"
	"xrChat_backend/internal/model"
	"xrChat_backend/pkg"

	"gorm.io/gorm"
)

// Login user login.
func Login(userInfo *model.User) (err error, user *model.User) {
	user = &model.User{}
	err = config.DB.Where("email = ?", userInfo.Email).First(user).Error
	if err != nil {
		slog.Error("%s", err)
		err = errors.New("账号或密码错误")
		return
	}
	userInfo.Password = pkg.EncryptPassword(user.Salt, userInfo.Password)
	if user.Password != userInfo.Password {
		err = errors.New("账号或密码错误")
		return
	}

	return
}

// VerifyEmail save email address and verify code to redis.
func VerifyEmail(email string, code string) (err error) {
	ctx := context.Background()
	err = config.RedisClient.Set(ctx, email, code, 30*60*time.Second).Err()
	if err != nil {
		slog.Error("Redis Set err:", err)
		return err
	}
	err = pkg.SendEmail(config.FromEmail, email, code, config.EmCode)
	if err != nil {
		slog.Error("Send Email err:", err)
		return err
	}
	return nil
}

// VerifyEmailCode verify email.
func VerifyEmailCode(email string, code string) (err error) {
	ctx := context.Background()
	var res string
	res, err = config.RedisClient.Get(ctx, email).Result()
	if err != nil {
		slog.Error("Redis Get err:", err)
		err = errors.New("邮箱验证失败")
		return
	}
	if res != code {
		err = errors.New("邮箱验证失败")
		return
	}
	return
}

// RegisterUser add user to database.
func RegisterUser(userInfo *model.User) (err error) {
	exist := EmailIsExist(userInfo.Email)
	if exist {
		return errors.New("该邮箱已被注册过")
	}
	err = config.DB.Create(userInfo).Error
	if err != nil {
		slog.Error("Create err:", err)
		err = errors.New("注册失败")
		return err
	}
	return nil
}

// EmailIsExist judge email is exist.
func EmailIsExist(email string) bool {
	result := config.DB.Where("email = ?", email).First(&model.User{})
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

// GetUserById
func GetUserById(id uint) (user *model.User, err error) {
	user = &model.User{}
	err = config.DB.Where("id = ?", id).First(user).Error
	if err != nil {
		slog.Error("get user by id error: ", err)
		return
	}
	return
}
