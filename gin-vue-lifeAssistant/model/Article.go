package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Article struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:char(36);primary_key"`
	UserId     uint      `json:"user_id" gorm:"not null"`
	CategoryId uint      `json:"category_id" gorm:"type:uint; not null"`
	Category   Category  `gorm:"foreignkey:CategoryId;"`
	Title      string    `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string    `json:"head_img" gorm:"type:varchar(100)"`
	Desc       string    `json:"desc" gorm:"type:varchar(100)"`
	Content    string    `json:"content" gorm:"type:text;not null"`
	// 解决唯一值索引 删除后不能新增同名索引的问题
	IsDel soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

// 创建文章前，利用gorm scope 赋值ID
func (article *Article) BeforeCreate(scope *gorm.DB) (err error) {
	article.ID = uuid.New()

	return
}
