package models

import (
	"gorm.io/gorm"
	"time"
)

type Company struct {
	gorm.Model
	Domain     string `gorm:"varchar:191"`
	Name       string `gorm:"varchar:191"`
	State      string `gorm:"varchar:191"`
	City       string `gorm:"varchar:191"`
	CreateDate time.Time
	ExpiryDate time.Time
	Code       string `gorm:"varchar:191"`
}
