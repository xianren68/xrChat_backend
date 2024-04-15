package repository

import (
	"errors"
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
