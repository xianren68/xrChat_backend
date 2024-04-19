package handler

import (
	"errors"
	"xrChat_backend/internal/proto/pb"
	"xrChat_backend/internal/repository"
	"xrChat_backend/pkg"

	"github.com/gin-gonic/gin"
)

func UpdateLine(c *gin.Context) {
	lineInfo := &pb.UpdateLine{}
	// parse query
	err := pkg.BindProto(c, lineInfo)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	err = repository.UpdateLine(lineInfo)
	if err != nil {
		pkg.HandleError(c, errors.New("修改失败"))
		return
	}
	pkg.HandleSuccess(c, "修改成功")
}

func UpdateName(c *gin.Context) {
	nameInfo := &pb.UpdateName{}
	// parse query
	err := pkg.BindProto(c, nameInfo)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	err = repository.UpdateName(nameInfo)
	if err != nil {
		pkg.HandleError(c, errors.New("修改失败"))
		return
	}
	pkg.HandleSuccess(c, "修改成功")
}

func UpdatePhone(c *gin.Context) {
	phoneInfo := &pb.UpdatePhone{}
	// parse query
	err := pkg.BindProto(c, phoneInfo)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	err = repository.UpdatePhone(phoneInfo)
	if err != nil {
		pkg.HandleError(c, errors.New("修改失败"))
		return
	}
	pkg.HandleSuccess(c, "修改成功")
}

func UpdateGender(c *gin.Context) {
	genderInfo := &pb.UpdateGender{}
	// parse query
	err := pkg.BindProto(c, genderInfo)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	err = repository.UpdateGender(genderInfo)
	if err != nil {
		pkg.HandleError(c, errors.New("修改失败"))
		return
	}
	pkg.HandleSuccess(c, "修改成功")
}
