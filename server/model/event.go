package model

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model

	PlaceID     uint   `json:"place_id" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}
