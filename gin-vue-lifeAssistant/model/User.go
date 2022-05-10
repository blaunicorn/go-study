package model

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=3,max=12" label:"用户名"`
	RealName  string `gorm:"type:varchar(20)" json:"realname"`
	Telephone string `gorm:"varchar(11);not null;unique" json:"telephone"`
	Password  string `gorm:"size:255;not null" json:"password" validate:"required,min=6,max=20"`
	Role      int    `gorm:"type:int;Default:2" json:"role" validate:"required,gte=2" label:"角色"`
	// gte=2 大于等于2
	Unionid string `gorm:"type:varchar(50);unique" json:"unionid"`
	// 解决唯一值索引 删除后不能新增同名索引的问题
	IsDel soft_delete.DeletedAt `gorm:"softDelete:flag"`
}
