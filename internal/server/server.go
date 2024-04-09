// Package server supply Http server.
package server

import (
	"github.com/gin-gonic/gin"
)

// InitRouter init router.
func InitRouter() *gin.Engine {
	engine := gin.Default()
	return engine
}
