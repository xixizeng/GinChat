package models

import (
	"fmt"
	"ginchat/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     uint64
	HeartbeatTime uint64
	LoginOutTime  uint64 `gorm:"column:login_out_time" json:"login-out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func FindUserByName(name string) UserBasic {
	data := UserBasic{}
	utils.DB.Where("name = ?", name).First(&data)
	return data
}

func FindUserByNameAndPassword(name, password string) UserBasic {
	data := UserBasic{}
	utils.DB.Where("name = ? and pass_word = ?", name, password).First(&data)

	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.Md5Encode(str)
	utils.DB.Model(&data).Where("id = ?", data.ID).Update("Identity", temp)
	return data
}

func FindUserByID(id uint) UserBasic {
	data := UserBasic{}
	utils.DB.Where("id = ?", id).First(&data)
	return data
}

func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("phone = ?", phone).First(&user)
}

func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("email = ?", email).First(&user)
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(user)
}
