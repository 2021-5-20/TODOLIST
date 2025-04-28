package models

import (
	"time"
)

type Memo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"size:255" json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 1) 外键字段，存储 User 的主键
	UserID uint `gorm:"not null;index" json:"user_id"`

	// 2) 关联到 User 模型
	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
