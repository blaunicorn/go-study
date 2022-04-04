package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	// ID   uint   `json:"id"`
	Name string `gorm:"varchar(50);not null;unique" json:"name"`
}
