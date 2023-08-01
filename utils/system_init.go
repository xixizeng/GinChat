package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

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
	//自定义日志模板打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)

	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println(" mysql inited。。。。。。。", viper.Get("mysql"))
}

func GetDB() *gorm.DB {
	return DB
}

func InitRedis() {
	var err error

	Red = redis.NewClient(&redis.Options{
		Password:     viper.GetString("redis.password"),
		Addr:         viper.GetString("redis.addr"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	ping, err := Red.Ping().Result()
	if err != nil {
		panic("failed to connect redis database")
	}
	fmt.Println(" redis inited。。。。。。。", ping)
}
