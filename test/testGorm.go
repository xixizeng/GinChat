// package main
//
// import (
//
//	"fmt"
//	"ginchat/models"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//
// )
//
//	type Product struct {
//		gorm.Model
//		Code  string
//		Price uint
//	}
//
//	func main() {
//		dsn := "root:123456@tcp(127.0.0.1:3307)/ginchat?charset=utf8&parseTime=True&loc=Local"
//		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//		if err != nil {
//			panic("failed to connect database")
//		}
//		db.AutoMigrate(&models.UserBasic{})
//		user := &models.UserBasic{}
//		user.Name = "申专"
//		db.Create(user)
//
//		fmt.Println(db.First(user, 1))
//		//db.First(user, "code = ?", "D42")
//		//db.Model(user).Update("Price", 200)
//		db.Model(user).Update("PassWord", "1234")
//	}
package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

//func main() {
//	r := gin.Default()
//	docs.SwaggerInfo.BasePath = "/api/v1"
//	v1 := r.Group("/api/v1")
//	{
//		eg := v1.Group("/example")
//		{
//			eg.GET("/helloworld", Helloworld)
//		}
//	}
//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
//	r.Run(":8080")
//
//}
