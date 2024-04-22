package repository

import (
	"xrChat_backend/config"
	"xrChat_backend/internal/model"
	"xrChat_backend/internal/proto/pb"
)

func UpdateLine(lineInfo *pb.UpdateLine) (err error) {
	user := &model.User{}
	user.ID = uint(lineInfo.Id)
	err = config.DB.First(user).Error
	if err != nil {
		return
	}
	return config.DB.Model(user).Update("line", lineInfo.Line).Error
}

func UpdateName(nameInfo *pb.UpdateName) (err error) {
	user := &model.User{}
	user.ID = uint(nameInfo.Id)
	err = config.DB.First(user).Error
	if err != nil {
		return
	}
	return config.DB.Model(user).Update("username", nameInfo.Name).Error
}

func UpdatePhone(phoneInfo *pb.UpdatePhone) (err error) {
	user := &model.User{}
	user.ID = uint(phoneInfo.Id)
	err = config.DB.First(user).Error
	if err != nil {
		return
	}
	return config.DB.Model(user).Update("phone", phoneInfo.Phone).Error
}

func UpdateGender(genderInfo *pb.UpdateGender) (err error) {
	user := &model.User{}
	user.ID = uint(genderInfo.Id)
	err = config.DB.First(user).Error
	if err != nil {
		return
	}
	return config.DB.Model(user).Update("gender", genderInfo.Gender).Error
}
