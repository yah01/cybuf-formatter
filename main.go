package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yah01/cybuf-formatter/handler"
)

func main() {
	router := gin.Default()

	router.POST("/format", handler.Format)
	router.Run(":4193")
}
