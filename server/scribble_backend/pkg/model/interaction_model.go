package model

// "iam " + id + " " + connectedCLIENTS
// "total "+connectedCLIENTS+" "+id
// "dis "+info.ID+" connectedCLIENTS
// "set "+info.ID+" "+info.X+" "+info.Y

type InteractionModel struct {
	ResponseType     string      `json:"response_type"`
	ID               string      `json:"id"`
	ConnectedClients []string    `json:"connected_clients"`
	ClientInfo       *ClientInfo `json:"client_info"`
}
