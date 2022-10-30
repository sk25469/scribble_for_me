package model

type ClientInfo struct {
	RoomID   string `json:"room_id"`
	ClientID string `json:"client_id"`
	Name     string `json:"name"`
}
