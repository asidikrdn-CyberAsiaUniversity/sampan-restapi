package models

import (
	"gorm.io/gorm"
)

type MstTrashCategory struct {
	gorm.Model
	Category string `gorm:"type:varchar(255);unique"`
	Price    int
}
