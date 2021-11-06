package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `json:"name" gorm:"type:varchar(50);not null;unique"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}
