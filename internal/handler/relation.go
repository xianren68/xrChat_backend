package handler

import (
	"strconv"
	"time"
	"xrChat_backend/config"
	"xrChat_backend/internal/proto/pb"
	"xrChat_backend/internal/repository"
	"xrChat_backend/internal/service"
	"xrChat_backend/pkg"

	"github.com/gin-gonic/gin"
)

// AddFriendReq add friend request.
func AddFriendReq(c *gin.Context) {
	addFriend := &pb.RelationOp{}
	err := pkg.BindProto(c, addFriend)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	// get target client
	client := config.UserPool.GetClientPool(strconv.Itoa(int(addFriend.TargetId)))
	msg := &pb.Message{
		Src:  addFriend.OwnerId,
		Tar:  addFriend.TargetId,
		Msg:  addFriend.Msg,
		Time: uint64(time.Now().Unix()),
	}
	sendMsg(client, msg, 5)
	pkg.HandleSuccess(c, "请求已发送")
}

// AddFriendRes response for add friend request.
func AddFriendRes(c *gin.Context) {
	res := &pb.RelationOp{}
	err := pkg.BindProto(c, res)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	client := config.UserPool.GetClientPool(strconv.Itoa(int(res.TargetId)))
	message := &pb.Message{
		Src:  res.OwnerId,
		Tar:  res.TargetId,
		Msg:  "",
		Time: uint64(time.Now().Unix()),
	}
	if res.Msg == "false" {
		sendMsg(client, message, 9)
		pkg.HandleSuccess(c, "success")
		return
	}
	err = service.AddFriendRes(res)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	sendMsg(client, message, 7)
	pkg.HandleSuccess(c, "success")
}

// CreateGroup  create group.
func CreateGroup(c *gin.Context) {
	req := &pb.CreateGroupRequest{}
	err := pkg.BindProto(c, req)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	err = service.CreateGroup(req)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	pkg.HandleSuccess(c, "创建成功")
}

// GetFriends  get all friends.
func GetFriends(c *gin.Context) {
	req := &pb.GetFriendsRequest{}
	err := pkg.BindProto(c, req)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	res := &pb.GetFriendsRes{}
	res.Code = 200
	friends, err := service.GetFriends(req)
	if err != nil {
		res.Code = 500
		res.Msg = err.Error()
		pkg.WriteProto(c, res)
		return
	}
	res.Friends = friends
	res.Msg = "获取成功"
	pkg.WriteProto(c, res)
}

// JoinGroupReq   join in group request.
func JoinGroupReq(c *gin.Context) {
	joinInfo := &pb.JoinGroupReq{}
	err := pkg.BindProto(c, joinInfo)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	// get target client
	client := config.UserPool.GetClientPool(strconv.Itoa(int(joinInfo.OwnerId)))
	msg := &pb.Message{
		Src:  joinInfo.SrcId,
		Tar:  joinInfo.GroupId,
		Msg:  joinInfo.Msg,
		Time: uint64(time.Now().Unix()),
	}
	sendMsg(client, msg, 6)
	pkg.HandleSuccess(c, "请求已发送")
}

// JoinGroupRes  response for join in group request.
func JoinGroupRes(c *gin.Context) {
	res := &pb.RelationOp{}
	err := pkg.BindProto(c, res)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	client := config.UserPool.GetClientPool(strconv.Itoa(int(res.TargetId)))
	message := &pb.Message{
		Src:  res.OwnerId, // group id
		Tar:  res.TargetId,
		Msg:  "",
		Time: uint64(time.Now().Unix()),
	}
	if res.Msg == "false" {
		sendMsg(client, message, 10)
		pkg.HandleSuccess(c, "success")
		return
	}
	err = service.JoinGroup(res)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	sendMsg(client, message, 8)
	pkg.HandleSuccess(c, "success")

}

// DelFriend  delete friend.
func DelFriend(c *gin.Context) {
	info := &pb.RelationOp{}
	err := pkg.BindProto(c, info)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	err = service.DelFriend(info)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	client := config.UserPool.GetClientPool(strconv.Itoa(int(info.TargetId)))
	message := &pb.Message{
		Src:  info.OwnerId,
		Tar:  info.TargetId,
		Msg:  "",
		Time: uint64(time.Now().Unix()),
	}
	sendMsg(client, message, 11)
	pkg.HandleSuccess(c, "success")
}

// KickOutGroup  delete user from group.
func KickOutGroup(c *gin.Context) {
	info := &pb.RelationOp{}
	err := pkg.BindProto(c, info)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	err = service.KickOutGroup(info)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	client := config.UserPool.GetClientPool(strconv.Itoa(int(info.TargetId)))
	message := &pb.Message{
		Src:  info.OwnerId,
		Tar:  info.TargetId,
		Msg:  "",
		Time: uint64(time.Now().Unix()),
	}
	sendMsg(client, message, 12)
	pkg.HandleSuccess(c, "success")
}

// QuitGroup  quit group.
func QuitGroup(c *gin.Context) {
	// this is special.
	// ownerId is group id.
	// targetId is user id.
	// so we need to get ownerId by groupId and notify it.
	info := &pb.RelationOp{}
	err := pkg.BindProto(c, info)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	err = service.QuitGroup(info)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	ownerId, err := repository.GetOwnerByGroupId(uint(info.OwnerId))
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	// notify group owner.
	client := config.UserPool.GetClientPool(strconv.Itoa(int(ownerId)))
	message := &pb.Message{
		Src:  info.TargetId,
		Tar:  uint64(ownerId),
		Msg:  "",
		Time: uint64(time.Now().Unix()),
	}
	sendMsg(client, message, 13)
	pkg.HandleSuccess(c, "success")
}
