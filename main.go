package main

import (
	"fmt"
	"xrChat_backend/config"
	"xrChat_backend/internal/server"
)

func main() {
	config.Global()
	router := server.InitRouter()
	err := router.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		panic(err)
	}
}
