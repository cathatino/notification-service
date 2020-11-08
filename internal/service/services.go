package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetRouter get the gin.Engine (before we start to run the gin.Engine)
func GetRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// RunRouter run gin.Engine by passing the gin.Engine & its port to run with
func RunRouter(
	router *gin.Engine,
	port Port,
) error {
	addr := fmt.Sprintf(":%d", port)
	return router.Run(addr)
}
