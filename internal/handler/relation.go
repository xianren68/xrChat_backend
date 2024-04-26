package handler

import (
	"strconv"
	"xrChat_backend/config"
	"xrChat_backend/internal/proto/pb"
	"xrChat_backend/internal/service"
	"xrChat_backend/pkg"

	"github.com/gin-gonic/gin"
)

// AddFriendReq add friend request.
func AddFriendReq(c *gin.Context) {
	addFriend := &pb.AddRequest{}
	err := pkg.BindProto(c, addFriend)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	// get target client
	client := config.UserPool.GetClientPool(strconv.Itoa(int(addFriend.TargetId)))
	msg := &pb.Message{
		Src: addFriend.OwnerId,
		Tar: addFriend.TargetId,
		Msg: addFriend.Msg,
	}
	sendMsg(client, msg, 5)
	pkg.HandleSuccess(c, "请求已发送")
}

// AddFriendRes response for add friend request.
func AddFriendRes(c *gin.Context) {
	res := &pb.AddRes{}
	err := pkg.BindProto(c, res)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	client := config.UserPool.GetClientPool(strconv.Itoa(int(res.TargetId)))
	message := &pb.Message{
		Src: res.OwnerId,
		Tar: res.TargetId,
	}
	if res.Msg == "false" {
		message.Msg = "用户" + res.Name + "拒绝了你的好友请求"
		sendMsg(client, message, 7)
		pkg.HandleSuccess(c, "success")
		return
	}
	err = service.AddFriendRes(res)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	message.Msg = "用户" + res.Name + "同意了你的好友请求"
	sendMsg(client, message, 7)
	pkg.HandleSuccess(c, "success")
}
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
