package models

import (
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
}
