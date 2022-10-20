package model

type Room struct {
	RoomID string   `json:"room_id"`
	Group1 []string `json:"grp1"`
	Group2 []string `json:"grp2"`
}
