package web

type MQTTRequest struct {
	ClientId string
	Topic    string
	Payload  string
	MsgResp  string
}
