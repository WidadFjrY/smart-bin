package model

import (
	"time"
)

type User struct {
	ID       string `gorm:"primaryKey;type:char(15);not null"`
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(100);not null"`
	// Notification []Notification `gorm:"foreignKey:UserID;references:ID"`
	SmartBin  []SmartBin `gorm:"foreignKey:UserID;references:ID"`
	Group     []Group    `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
