package server

import (
	"xrChat_backend/config"
	"xrChat_backend/internal/handler"

	"github.com/fwhezfwhez/tcpx"
)

func TcpServer() *tcpx.TcpX {
	srv := tcpx.NewTcpX(tcpx.ProtobufMarshaller{}).WithBuiltInPool(true)
	config.UserPool = srv.GetPool()
	srv.AddHandler(1, handler.OnLine)
	srv.AddHandler(2, handler.OffLine)
	srv.AddHandler(3, handler.FriendMsgHandler)
	srv.AddHandler(4, handler.GroupMsgHandler)
	return srv
}
