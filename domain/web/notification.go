package web

type NotificationCreateRequest struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
