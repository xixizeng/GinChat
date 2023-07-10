package router

import (
	"ginchat/docs"
	"ginchat/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary 测试index
// @Description 用来测试swagger
// @Tags
// @Accept json
// @Produce json
// @Success 200 {string} welcome
// @Fialure 500 {string} notwelcom
// @Router /index [get]

func Router() *gin.Engine {

	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/index", service.GetIndex)
	router.GET("/user/getUserList", service.GetUserList)

	return router
}
