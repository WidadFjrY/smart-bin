package model

import "time"

type Group struct {
	ID        string     `gorm:"primaryKey;type:char(15);not null"`
	UserID    string     `gorm:"type:char(15);not null"`
	Name      string     `gorm:"type:varchar(100);not null"`
	Location  string     `gorm:"type:varchar(100);not null"`
	SmartBin  []SmartBin `gorm:"foreignKey:GroupID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
