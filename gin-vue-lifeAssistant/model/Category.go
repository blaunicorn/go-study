package model

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Category struct {
	gorm.Model
	// ID   uint   `json:"id"`
	Name string `gorm:"varchar(50);not null;unique" json:"name"`
	// 解决唯一值索引 删除后不能新增同名索引的问题
	IsDel soft_delete.DeletedAt `gorm:"softDelete:flag"`
}
