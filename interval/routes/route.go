package routes

import (
	"github.com/gin-gonic/gin"
)

func Serve(addr string) {
	router := gin.Default()

	router.POST("/mysql/connect")

	router.Run(addr)
}
