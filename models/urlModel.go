package models

import (
	"gorm.io/gorm"
)

type URLS struct {
	gorm.Model
	ID   string `gorm:"unique,index"`
	Link string `gorm:"index"`
}
