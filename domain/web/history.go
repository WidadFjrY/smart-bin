package web

import "time"

type HistoryCreateRequest struct {
	BinID   string `json:"bin_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type HistoryGetResponse struct {
	BinId   string    `json:"id"`
	History []History `json:"history"`
}

type History struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
