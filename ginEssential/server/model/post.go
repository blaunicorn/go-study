package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID         uuid.UUID `gorm:"type:char(36);primary_key"`
	UserId     uint      `json:"user_id" gorm:"not null"`
	CategoryId uint      `json:"category_id" gorm:"not null"`
	Category   *Category
	Title      string         `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string         `json:"head_img"`
	Content    string         `json:"content" gorm:"type:text;not null"`
	CreatedAt  time.Time      `json:"created_at" `
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// 创建文章前，利用gorm scope 赋值ID
func (post *Post) BeforeCreate(scope *gorm.DB) (err error) {
	post.ID = uuid.New()

	return
}
