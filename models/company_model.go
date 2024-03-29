package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Domain     string `gorm:"varchar:191"`
	Name       string `gorm:"varchar:191"`
	State      string `gorm:"varchar:191"`
	City       string `gorm:"varchar:191"`
	CreateDate string `gorm:"varchar:191"`
	ExpiryDate string `gorm:"varchar:191"`
	Code       string `gorm:"varchar:191"`
}
