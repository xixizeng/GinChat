package router

import (
	"ginchat/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/index", service.GetIndex)
	router.GET("/getuserlist", service.GetUserList)
	return router
}
