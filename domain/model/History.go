package model

import "time"

type History struct {
	ID        string `gorm:"primaryKey;type:char(15);not null"`
	BinID     string `gorm:"type:char(15);not null"`
	Status    string `gorm:"type:varchar(20);"`
	Message   string `gorm:"type:varchar(255);"`
	CreatedAt time.Time
}
