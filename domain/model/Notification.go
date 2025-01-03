package model

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID        string `gorm:"primaryKey;type:char(15);not null"`
	UserID    string `gorm:"type:char(15);not null"`
	Title     string `gorm:"type:varchar(100);"`
	Desc      string `gorm:"type:varchar(255);"`
	IsRead    bool   `gorm:"type:boolean;default:false"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
