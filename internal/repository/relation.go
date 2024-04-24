package repository

import (
	"errors"
	"log/slog"
	"xrChat_backend/config"
	"xrChat_backend/internal/model"
)

func AddFriendReq(relation *model.Relation) (err error) {
	if relation.OwnerId == relation.TargetId {
		err = errors.New("自己无法添加自己为好友")
		return
	}
	if FindUserById(relation.TargetId) {
		err = errors.New("您要添加的用户不存在")
		return
	}
	// TODO if target online,send tcp,if outline add to redis.
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
