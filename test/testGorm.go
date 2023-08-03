package main

import (
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dns := "root:123456@tcp(127.0.0.1:3307)/ginchat?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		panic("connect sql failed ")
	}

	db.AutoMigrate(&models.UserBasic{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.GroupBasic{})
	db.AutoMigrate(&models.Contact{})
}
