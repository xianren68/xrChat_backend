package repository

import (
	"errors"
	"log/slog"
	"xrChat_backend/config"
	"xrChat_backend/internal/model"
)

func AddFriendRes(relation *model.Relation) (err error) {
	if relation.OwnerId == relation.TargetId {
		err = errors.New("自己无法添加自己为好友")
		return
	}
	if FindUserById(relation.TargetId) {
		err = errors.New("您要添加的用户不存在")
		return
	}
	// start a transaction
	tx := config.DB.Begin()
	if tx.Error != nil {
		slog.Error("AddFriendRes", "err", tx.Error)
	}
	failErr := errors.New("添加好友失败")
	// add friend
	err = tx.Create(relation).Error
	if err != nil {
		slog.Error("AddFriendRes", "err", err)
		tx.Rollback()
		err = failErr
		return
	}
	re := &model.Relation{
		OwnerId:  relation.TargetId,
		TargetId: relation.OwnerId,
		Type:     1,
	}
	err = tx.Create(re).Error
	if err != nil {
		slog.Error("AddFriendRes", "err", err)
		tx.Rollback()
		err = failErr
		return
	}
	return nil

}

func FindUserById(userId uint) bool {
	md := new(model.User)
	tx := config.DB.Where("user_id = ?", userId).First(md)
	return tx.Error == nil
}

// GetMembers get all members of group.
func GetMembers(groupId uint) (members []uint, err error) {
	members = make([]uint, 0)
	err = config.DB.Model(&model.Relation{}).Where("type = ? and owner_id = ?", 2, groupId).Pluck("target_id", &members).Error
	if err != nil {
		slog.Error("GetMembers", "err", err)
		return
	}
	return
}

func GetRelations(userId uint) (relations []*model.Relation, err error) {
	relations = make([]*model.Relation, 0)
	err = config.DB.Where("owner_id = ? AND type = ?", userId, 1).Find(&relations).Error
	if err != nil {
		slog.Error("GetRelations", "err", err)
		return
	}
	return
}

func CreateGroup(group *model.Group) (err error) {

	err = config.DB.Create(group).Error
	if err != nil {
		slog.Error("CreateGroup", "err", err)
		err = errors.New("创建群组失败")
		return
	}
	return
}
