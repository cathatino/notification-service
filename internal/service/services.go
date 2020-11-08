package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetRouter() {
	router := gin.Default()
	fmt.Println(router)
}

func init() {
	GetRouter()
}