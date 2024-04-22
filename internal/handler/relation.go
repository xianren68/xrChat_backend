package handler

import (
	"xrChat_backend/internal/proto/pb"
	"xrChat_backend/internal/service"
	"xrChat_backend/pkg"

	"github.com/gin-gonic/gin"
)

// AddFriendReq add friend request.
func AddFriendReq(c *gin.Context) {
	addFriend := &pb.AddFriendRequest{}
	err := pkg.BindProto(c, addFriend)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	err = service.AddFriendReq(addFriend)
	if err != nil {
		pkg.HandleError(c, err)
		return
	}
	pkg.HandleSuccess(c, "请求已发送")
}


func CreateGroup(c *gin.Context) {
	
}
