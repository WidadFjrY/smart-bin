package web

type ConfigGetResponse struct {
	Id        string  `json:"id"`
	MaxHeight float64 `json:"max_height"`
	MaxWeight float64 `json:"max_weight"`
}
