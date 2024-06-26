package handler

import (
	"fmt"
	"log/slog"
	"strconv"
	"xrChat_backend/internal/proto/pb"
	"xrChat_backend/internal/repository"
	"xrChat_backend/internal/service"

	"github.com/fwhezfwhez/tcpx"
)

// OnLine user online
func OnLine(c *tcpx.Context) {
	message := &pb.Message{}
	Info, err := c.Bind(message)
	if err != nil {
		slog.Error("OnLine Bind error:", err)
		return
	}
	msgContent := (Info.Body).(pb.Message)
	err = c.Online(strconv.Itoa(int(msgContent.Src)))
	if err != nil {
		slog.Error("OnLine online error:", err)
		return
	}
	fmt.Println("hello world")
	slog.Info("OnLine:", msgContent.Src)
}

// OffLine user offline
func OffLine(c *tcpx.Context) {
	message := &pb.Message{}
	Info, err := c.Bind(message)
	if err != nil {
		slog.Error("OffLine Bind error:", err)
		return
	}
	msgContent := (Info.Body).(*pb.Message)
	_ = c.Offline()
	slog.Info("OffLine:", msgContent.Src)
}

func FriendMsgHandler(c *tcpx.Context) {
	message := &pb.Message{}
	Info, err := c.Bind(message)
	if err != nil {
		slog.Error("friendMsgHandler Bind error:", err)
		return
	}
	msgContent := (Info.Body).(*pb.Message)
	pool := c.GetPoolRef()
	client := pool.GetClientPool(strconv.Itoa(int(msgContent.Tar)))
	sendMsg(client, msgContent, Info.MessageID)
}

func GroupMsgHandler(c *tcpx.Context) {
	message := &pb.Message{}
	Info, err := c.Bind(message)
	if err != nil {
		slog.Error("groupMsgHandler Bind error:", err)
		return
	}
	msgContent := (Info.Body).(*pb.Message)
	members, err := repository.GetMembers(uint(msgContent.Tar))
	if err != nil {
		slog.Error("groupMsgHandler GetMembers error:", err)
		// TODO handle error.
		return
	}
	pool := c.GetPoolRef()
	for _, member := range members {
		client := pool.GetClientPool(strconv.Itoa(int(member)))
		sendMsg(client, msgContent, Info.MessageID)
	}
}

func sendMsg(c *tcpx.Context, message *pb.Message, MessageID int32) {
	if c != nil {
		// if target offline,save unread message to mysql.
		err := service.SaveUnread(message, int(MessageID))
		if err != nil {
			slog.Error("sendMsg save unread error:", err)
		}
		return
	}
	err := c.Reply(MessageID, message)
	if err != nil {
		slog.Error("sendMsg reply error:", err)
	}
}
