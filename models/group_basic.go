package models

import "gorm.io/gorm"

// 群组信息
type GroupBasic struct {
	gorm.Model
	Name    string //群名称
	OwnerId uint   //拥有者
	Icon    string //群头像
	Type    int    //等级
	Desc    string //
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
