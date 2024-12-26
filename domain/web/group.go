package web

import "time"

type GroupCreateRequest struct {
	Name     string `json:"name" validate:"required,min=1"`
	Location string `json:"location" validate:"required,min=1"`
}

type GroupCreateResponse struct {
	GroupId   string    `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GroupGetResponse struct {
	GroupId string `json:"group_id"`
	Group   Group  `json:"group"`
}

type Group struct {
	UserId    string      `json:"user_id"`
	Name      string      `json:"name"`
	Location  string      `json:"location"`
	Bins      []SmartBins `json:"bins"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type SmartBins struct {
	Id   string `json:"id"`
	Data Data   `json:"data"`
}

type Data struct {
	UserID               string            `json:"user_id"`
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

type GroupGetResponses struct {
	Id    string         `json:"id"`
	Group GroupResponses `json:"group"`
}

type GroupResponses struct {
	UserId    string          `json:"user_id"`
	Name      string          `json:"name"`
	Location  string          `json:"location"`
	Bins      []SmartBinNames `json:"bins"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type SmartBinNames struct {
	BinId    string `json:"bin_id"`
	Name     string `json:"name"`
	IsLocked bool   `json:"is_locked"`
}

type GroupUpdateRequest struct {
	Name     string `json:"name" validate:"required,min=1"`
	Location string `json:"location" validate:"required,min=1"`
}

type GroupUpdateResponse struct {
	Id        string    `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
}
