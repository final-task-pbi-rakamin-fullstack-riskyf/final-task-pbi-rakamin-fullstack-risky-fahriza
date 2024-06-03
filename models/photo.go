package models

import (
	"time"
)

type Photo struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Caption   string
	PhotoUrl  string    `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}