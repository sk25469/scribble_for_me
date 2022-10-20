package model

type Room struct {
	RoomID           string            `json:"room_id"`
	Group1           []string          `json:"client_grp1"`
	Group2           []string          `json:"client_grp2"`
	InteractionModel *InteractionModel `json:"interaction_model"`
}
