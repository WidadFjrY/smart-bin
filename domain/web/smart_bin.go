package web

import (
	"time"
)

type SmartBinCreateRequest struct {
	BinId string `json:"bin_id" validator:"required,min=15"`
}

type SmartBinCreateResponse struct {
	ID      string    `json:"id"`
	AddedAt time.Time `json:"added_at"`
}

type SmartBinUpdateRequest struct {
	Name     string `json:"name" validator:"required"`
	Location string `json:"location" validator:"required"`
}

type SmartBinUpdateResponse struct {
	ID        string    `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SmartBinDeleteResponse struct {
	ID        string    `json:"id"`
	DeletedAt time.Time `json:"deleted_at"`
}

type SmartBinGetResponse struct {
	Id       string   `json:"id"`
	SmartBin SmartBin `json:"smart_bin"`
}

type SmartBin struct {
	UserID               string            `json:"user_id"`
	GroupID              string            `json:"group"`
	Name                 string            `json:"name"`
	OrganicWeight        interface{}       `json:"organic_weight"`
	NonOrganicWeight     interface{}       `json:"non_organic_weight"`
	OrganicHeight        interface{}       `json:"organic_height"`
	NonOrganicHeight     interface{}       `json:"non_organic_height"`
	TotalOrganicWaste    int               `json:"total_organic_waste"`
	TotalNonOrganicWaste int               `json:"total_non_organic_waste"`
	IsLocked             bool              `json:"is_locked"`
	Location             string            `json:"location"`
	Config               ConfigGetResponse `json:"config"`
	// History              []History `gorm:"foreignKey:BinID;references:ID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SmartBinClassification struct {
	BinId      string    `json:"bin_id"`
	Waste      string    `json:"waste"`
	ClassifyAt time.Time `json:"classify_at"`
}

type ClassifyRequest struct {
	PathImage string `json:"path_img"`
}

type ClassifyResponse struct {
	Status     string `json:"status"`
	Prediction string `json:"prediction"`
}

type UpdateValueRequest struct {
	LoadCellOrganic      float64 `json:"load_cell_organic"`
	LoadCellNonOrganic   float64 `json:"load_cell_non_organic"`
	UltraSonicOrganic    float64 `json:"ultra_sonic_organic"`
	UltraSonicNonOrganic float64 `json:"ultra_sonic_non_organic"`
}

type UpdateValueResponse struct {
	ID        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Loked     bool      `json:"loked"`
	LokedDesc string    `json:"loked_desc"`
	UpdatedAt time.Time `json:"updated_at"`
}
