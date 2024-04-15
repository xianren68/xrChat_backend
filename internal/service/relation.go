package service

import (
	"xrChat_backend/internal/model"
	"xrChat_backend/internal/proto/pb"
	"xrChat_backend/internal/repository"
)

func AddFriendReq(requestInfo *pb.AddFriendRequest) (err error) {
	relation := &model.Relation{}
	relation.OwnerId = uint(requestInfo.OwnerId)
	relation.Remark = requestInfo.Remark
	relation.TargetId = uint(requestInfo.TargetId)
	err = repository.AddFriendReq(relation)
	if err != nil {
		return err
	}
	return nil
}
