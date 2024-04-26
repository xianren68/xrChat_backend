package service

import (
	"xrChat_backend/internal/model"
	"xrChat_backend/internal/proto/pb"
	"xrChat_backend/internal/repository"
)

func AddFriendRes(res *pb.AddRes) (err error) {
	relation := &model.Relation{}
	relation.OwnerId = uint(res.OwnerId)
	relation.Remark = res.Remark
	relation.TargetId = uint(res.TargetId)
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
