package model

import (
	"gorm.io/gorm"
)

type Place struct {
	gorm.Model

	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Address     string `json:"address" gorm:"not null"`
}
