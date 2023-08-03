package router

import (
	"ginchat/docs"
	"ginchat/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	router := gin.Default()
	//swagger
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//静态资源
	router.Static("/asset", "asset/")
	router.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	router.LoadHTMLGlob("views/**/*")
	//首页
	router.GET("/", service.GetIndex)
	router.GET("/index", service.GetIndex)

	router.GET("/toRegister", service.ToRegister)
	//用户模块
	router.GET("/user/getUserList", service.GetUserList)
	router.POST("/user/createUser", service.CreateUser)
	router.DELETE("/user/deleteUser", service.DeleteUser)
	router.POST("/user/updateUser", service.UpdateUser)
	router.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPassword)
	//发送消息
	router.GET("/user/sendMsg", service.SenMsg)
	router.GET("/user/sendUserMsg", service.SenUserMsg)
	return router
}
