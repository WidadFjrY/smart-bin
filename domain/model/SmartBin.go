package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type SmartBin struct {
	ID                   string         `gorm:"primaryKey;type:char(15);not null"`
	UserID               string         `gorm:"type:char(15);not null"`
	GroupID              *string        `gorm:"type:char(15);default:null;"`
	Name                 string         `gorm:"type:varchar(100);not null"`
	LoadCellValue        datatypes.JSON `gorm:"type:json"`
	UltraSonicValue      datatypes.JSON `gorm:"type:json"`
	TotalOrganicWaste    int            `gorm:"type:int;default:0"`
	TotalNonOrganicWaste int            `gorm:"type:int;default:0"`
	IsLocked             bool           `gorm:"type:bool;default:false"`
	Location             string         `gorm:"type:varchar(100);not null"`
	Config               Config         `gorm:"foreignKey:BinID;references:ID"`
	History              []History      `gorm:"foreignKey:BinID;references:ID"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}
