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
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/index", service.GetIndex)
	router.GET("/user/getUserList", service.GetUserList)
	router.PUT("/user/createUser", service.CreateUser)
	router.DELETE("/user/deleteUser", service.DeleteUser)
	router.POST("/user/updateUser", service.UpdateUser)
	router.POST("/user/findUserByNameAndPassword", service.FindUserByNameAndPassword)
	return router
}
