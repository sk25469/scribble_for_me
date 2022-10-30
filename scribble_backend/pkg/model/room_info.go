package model

type Room struct {
	RoomID string       `json:"room_id"`
	Group1 []ClientInfo `json:"grp1"`
	Group2 []ClientInfo `json:"grp2"`
}
