package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Login    string `json:"login" gorm:"not null"`
	Password string `json:"password" gorm:"not null"` // hash
	Age      uint8  `json:"age" gorm:"not null"`
	Role     uint8  `json:"role" gorm:"not null"`
}
