package model

//	room_type can be "public" or "private"
type ClientResponse struct {
	ReponseType string `json:"response_type"`

	// public or private
	RoomID     string      `json:"room_id"`
	ClientInfo *ClientInfo `json:"client_info"`
	RoomType   string      `json:"room_type"`
}
