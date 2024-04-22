package repository

import (
	"log/slog"
	"xrChat_backend/config"
	"xrChat_backend/internal/model"
)

func CreateGroup(groupInfo *model.Group) (res *model.Group, err error) {
	err = config.DB.Create(groupInfo).Error
	if err != nil {
		slog.Error("创建群组失败", err)
		return
	}
	return
}
