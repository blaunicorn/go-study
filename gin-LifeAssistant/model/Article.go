package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:char(36);primary_key"`
	UserId     uint      `json:"user_id" gorm:"not null"`
	CategoryId uint      `json:"category_id" gorm:"type:uint; not null"`
	Category   *Category
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string `json:"head_img" gorm:"type:varchar(100)"`
	Desc       string `json:"desc" gorm:"type:varchar(100)"`
	Content    string `json:"content" gorm:"type:text;not null"`
}

// 创建文章前，利用gorm scope 赋值ID
func (article *Article) BeforeCreate(scope *gorm.DB) (err error) {
	article.ID = uuid.New()

	return
}
