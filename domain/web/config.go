package web

import "time"

type ConfigGetResponse struct {
	ConfigId  string  `json:"config_id"`
	MaxHeight float64 `json:"max_height"`
	MaxWeight float64 `json:"max_weight"`
}

type ConfigUpdateRequest struct {
	MaxHeight float64 `json:"max_height" validate:"required,min=1"`
	MaxWeight float64 `json:"max_weight" validate:"required,min=1"`
}

type ConfigUpdateRespponse struct {
	BinId     string    `json:"bin_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
