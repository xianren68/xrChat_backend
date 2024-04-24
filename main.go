package main

import (
	"fmt"
	"xrChat_backend/config"
	"xrChat_backend/internal/server"
)

func main() {
	// get Config and set global variable.
	config.Global()
	// start tcpServer.
	go func() {
		srv := server.TcpServer()
		err := srv.ListenAndServe("tcp", fmt.Sprintf(":%d", config.TcpPort))
		if err != nil {
			panic("tcp server start failed" + err.Error())
		}
		return
	}()
	router := server.InitRouter()
	// start httpServer.
	err := router.Run(fmt.Sprintf(":%d", config.HttpPort))
	if err != nil {
		panic(err)
	}
}
