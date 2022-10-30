package model

// "iam " + id + " " + connectedCLIENTS
// "total "+connectedCLIENTS+" "+id
// "dis "+info.ID+" connectedCLIENTS
// "set "+info.ID+" "+info.X+" "+info.Y

type ServerResponse struct {
	ResponseType string `json:"response_type"`
	//	ClientInfo from where the request is coming from
	ClientInfo ClientInfo `json:"client_info"`
	RoomInfo   Room       `json:"room_info"`
}
