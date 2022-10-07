package model

// "iam " + id + " " + connectedCLIENTS
// "total "+connectedCLIENTS+" "+id
// "dis "+info.ID+" connectedCLIENTS
// "set "+info.ID+" "+info.X+" "+info.Y

type ClientServerResponse struct {
	ResponseType     string     `json:"response_type"`
	ID               string     `json:"id"`
	ConnectedClients []string   `json:"connected_clients"`
	PointInfo        *PointInfo `json:"point_info"`
}
