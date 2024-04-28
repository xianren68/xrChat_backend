package service

import (
	"errors"
	"xrChat_backend/internal/model"
	"xrChat_backend/internal/proto/pb"
	"xrChat_backend/internal/repository"
)

func AddFriendRes(res *pb.RelationOp) (err error) {
	relation := &model.Relation{}
	relation.OwnerId = uint(res.OwnerId)
	relation.TargetId = uint(res.TargetId)
	relation.Type = 1
	err = repository.AddFriendRes(relation)
	if err != nil {
		return err
	}
	return nil
}

func GetFriends(requestInfo *pb.GetFriendsRequest) (friends []*pb.Friend, err error) {
	relations, err := repository.GetRelations(uint(requestInfo.OwnerId))
	if err != nil {
		return
	}
	friends = make([]*pb.Friend, 0)
	for _, relation := range relations {
		user, err := repository.GetUserById(relation.TargetId)
		if err != nil {
			continue
		}
		friend := &pb.Friend{
			Id:     uint64(user.ID),
			Name:   user.Username,
			Avatar: user.Avatar,
			Remark: relation.Remark,
			Line:   user.Line,
			Phone:  user.Phone,
			Email:  user.Email,
			Gender: user.Gender,
		}
		friends = append(friends, friend)
	}
	return
}

func CreateGroup(requestInfo *pb.CreateGroupRequest) (err error) {
	group := &model.Group{}
	group.Name = requestInfo.Name
	group.OwnerId = uint(requestInfo.OwnerId)
	group.Avatar = requestInfo.Avatar
	group.Desc = requestInfo.Desc
	return repository.CreateGroup(group)
}

func JoinGroup(requestInfo *pb.RelationOp) (err error) {
	relation := &model.Relation{}
	relation.OwnerId = uint(requestInfo.OwnerId)
	relation.TargetId = uint(requestInfo.TargetId)
	relation.Type = 2
	return repository.JoinGroup(relation)
}

func DelFriend(requestInfo *pb.RelationOp) (err error) {
	relation := &model.Relation{}
	relation.OwnerId = uint(requestInfo.OwnerId)
	relation.TargetId = uint(requestInfo.TargetId)
	relation.Type = 1
	return repository.DelFriend(relation)
}

func KickOutGroup(requestInfo *pb.RelationOp) (err error) {
	relation := &model.Relation{}
	relation.OwnerId = uint(requestInfo.OwnerId)
	relation.TargetId = uint(requestInfo.TargetId)
	relation.Type = 2
	return repository.KickOutGroup(relation)
}

func QuitGroup(requestInfo *pb.RelationOp) (err error) {
	relation := &model.Relation{}
	relation.OwnerId = uint(requestInfo.OwnerId)
	relation.TargetId = uint(requestInfo.TargetId)
	relation.Type = 2
	err = repository.KickOutGroup(relation)
	if err != nil {
		err = errors.New("退出群组失败")
		return
	}
	return
}
