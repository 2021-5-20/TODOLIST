package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique" json:"username" binding:"required"`
	Password  string    `gorm:"size:255" json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}
