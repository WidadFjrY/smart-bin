package model

import "time"

type Config struct {
	ID        string  `gorm:"primaryKey;type:char(15);not null"`
	BinID     string  `gorm:"type:char(15);not null;unique"`
	MaxHeight float64 `gorm:"type:float;default:60"`
	MaxWeight float64 `gorm:"type:float;default:5"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
