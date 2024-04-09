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
)

// Login user login.
func Login(userInfo *model.User) (err error) {
	selectUser := &model.User{}
	err = config.DB.Where("user_id = ?", userInfo.UserId).First(&selectUser).Error
	if err != nil {
		return err
	}
	if selectUser.Password == userInfo.Password {
		return err
	}
	return errors.New("用户名或密码错误")
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
	err = config.DB.Create(userInfo).Error
	if err != nil {
		slog.Error("Create err:", err)
		err = errors.New("注册失败")
		return err
	}
	return nil
}
