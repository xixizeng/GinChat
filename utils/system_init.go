package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app inited", viper.Get("app"))

}
func InitMysql() {
	var err error

	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println(" mysql inited。。。。。。。", viper.Get("mysql"))
}

func GetDB() *gorm.DB {
	return DB
}
