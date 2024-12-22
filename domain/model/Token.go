package model

import (
	"time"

	"github.com/google/uuid"
)

type TokenAuth struct {
	ID        uuid.UUID `gorm:"primaryKey;type:char(36);default:(UUID())"`
	Token     string
	IsValid   bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
