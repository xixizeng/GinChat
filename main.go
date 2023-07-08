package main

import (
	"ginchat/router"
	"ginchat/utils"
)

//@title Go-Web开发记录
//@version 0.0.1
//@description 尝试swagger

func main() {
	utils.InitConfig()
	utils.InitMysql()
	router.Router()
	r := router.Router()
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
