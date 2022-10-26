package model

type ClientInfo struct {
	RoomID   string `json:"room_id"`
	ClientID string `json:"client_id"`
	X        string `json:"x"`
	Y        string `json:"y"`
	Name     string `json:"name"`
}
