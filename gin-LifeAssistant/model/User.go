package model

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null" json:"username"`
	RealName  string `gorm:"type:varchar(20)" json:"realname"`
	Telephone string `gorm:"varchar(11);not null;unique" json:"telephone"`
	Password  string `gorm:"size:255;not null" json:"password"`
	Role      int    `gorm:"type:int" json:"role"`
	// 解决唯一值索引 删除后不能新增同名索引的问题
	IsDel soft_delete.DeletedAt `gorm:"softDelete:flag"`
}
