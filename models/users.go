package models

import (
	"time"
)

type User struct {
	ID        uint    `gorm:"primary_key"`
	Name      string  `gorm:"type:varchar(255);not null" `
	Email     string  `gorm:"uniqueIndex;not null"`
	Tags      string  `gorm:"index; nullable"`
	Rating    float64 `gorm:"nullable"`
	Password  string  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserModel struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
