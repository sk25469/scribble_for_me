package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/olahol/melody"
	"github.com/sk25469/scribble_backend/pkg/config"
	"github.com/sk25469/scribble_backend/pkg/model"
	"github.com/sk25469/scribble_backend/pkg/utils"
)

var connectedClients []string
var mrouter = config.GetWebSocketRouter()
var response *model.ServerResponse
var rooms map[string]*model.Room

// 1st key is roomID, 2nd key is clientID and value is session
var totalClientsInSession map[string](map[string]*melody.Session)

// TYPES OF REQUEST SENT BY SERVER
//
//  1. "new" : A new client joins the network, but is not currently in any room
//
//  2. "iam" : The newly joined client sends a name and the kind of room it wants to join
//
//  3. "total" : Other clients are informed if any new client joins that room
//
//  4. "set" : When a client is drawing, it will send its x,y co-ordinate to others in the room
//
//  5. "dis" : Informs others in the room that "id" has disconnected
func OnConnect(s *melody.Session) {

	// now the new session is assigned a new id
	id := uuid.NewString()

	// and we set the initial info as the following to the current session
	// in the main server

	// Set "stores" the key, value pair for this session in the server
	// we will be adding the name after this request is sent and a response is recieved
	clientInfo := model.ClientInfo{ClientID: id, X: "0", Y: "0"}
	s.Set("info", &clientInfo)

	// write send the message to the client to set its id, and the size of the
	// fmt.Printf("after sending client %v\n", connectedClients)
	connectedClients = append(connectedClients, id)
	// current connected sessions
	response = &model.ServerResponse{ResponseType: "iam", ID: id, ConnectedClients: connectedClients, ClientInfo: &clientInfo}
	jsonResponse, err := json.Marshal(&response)
	if err != nil {
		log.Print("can't marshall reponse")
	}

	err = s.Write([]byte(jsonResponse))

	if err != nil {
		log.Fatal(err)
	}

}

func OnDisconnect(s *melody.Session) {
	info := s.MustGet("info").(*model.ClientInfo)
	var err error
	connectedClients, err = utils.Remove(connectedClients, info.ClientID)
	if err != nil {
		log.Fatal(err)
	}
	response.ConnectedClients = connectedClients
	response.ID = info.ClientID
	response.ResponseType = "dis"
	jsonResponse, err := json.Marshal(&response)
	if err != nil {
		log.Print("can't marshall reponse")
	}
	fmt.Printf("size before broadcasting %v\n", len(connectedClients))
	mrouter.BroadcastOthers([]byte(jsonResponse), s)

}

// TYPES OF REQUEST SENT BY CLIENT
//
//  1. "connect-new" : When a client has entered its name and he wants to join a new room
//
//  2. "connect" : Client wants to connect to an existing room with ID
//
//  3. "move" : A client is drawing on the screen
func OnMessage(s *melody.Session, msg []byte) {
	var clientResponse *model.ClientResponse
	err := json.Unmarshal(msg, &clientResponse)
	if err != nil {
		log.Fatal(err)
	}
	info := s.MustGet("info").(*model.ClientInfo)
	clientID := info.ClientID
	clientName := clientResponse.ClientInfo.Name
	if clientResponse.ReponseType == "connect-new" {
		var newRoomID string

		// if he wants to create a private room, a new key for room is created
		// otherwise it is assigned as "public"
		if clientResponse.RoomType == "private" {
			newRoomID = utils.GetKey()
		} else {
			newRoomID = "public"
		}
		// update the info of the current session with its roomID and user name
		s.Set("info", &model.ClientInfo{RoomID: newRoomID, ClientID: clientID, Name: clientName, X: "0", Y: "0"})

		grp1, grp2 := utils.InsertClientInRoom(rooms[newRoomID].Group1, rooms[newRoomID].Group2, clientID)
		rooms[newRoomID] = &model.Room{RoomID: newRoomID, Group1: grp1, Group2: grp2}

		// mapping the clientID with the sessions
		totalClientsInSession[newRoomID][clientID] = s
	} else {
		// check if the roomID exists
		roomID := clientResponse.RoomID
		if _, ok := rooms[roomID]; !ok {
			log.Fatal("given room doesn't exists")
		}
		s.Set("info", &model.ClientInfo{RoomID: roomID, ClientID: clientID, Name: clientName, X: "0", Y: "0"})

		grp1, grp2 := utils.InsertClientInRoom(rooms[roomID].Group1, rooms[roomID].Group2, clientID)
		rooms[roomID] = &model.Room{RoomID: roomID, Group1: grp1, Group2: grp2}
		totalClientsInSession[roomID][clientID] = s

		//	TODO: Send other clients connected in the room info that a new client has joined

	}
	// if len(p) == 2 {
	// 	// we get the info of the current session from the server
	// 	info := s.MustGet("info").(*model.ClientInfo)

	// 	// we assign the x and y coordinates to it,
	// 	// every time there is some new activity on the client
	// 	info.X = p[0]
	// 	info.Y = p[1]
	// 	response.ResponseType = "set"
	// 	response.ID = info.ClientID
	// 	response.ClientInfo = &model.ClientInfo{ClientID: info.ClientID, X: p[0], Y: p[1]}

	// 	jsonResponse, err := json.Marshal(&response)
	// 	if err != nil {
	// 		log.Print("can't marshall reponse")
	// 	}

	// 	// then sends the message to all others
	// 	mrouter.BroadcastOthers([]byte(jsonResponse), s)
	// 	fmt.Println(info)
	// }
}

// sends message in a group, can be in same or another
func SendMessageInGroup(grp []string, clientInfo *model.ClientInfo) {
	roomID := clientInfo.RoomID
	clientID := clientInfo.ClientID
	for i := 0; i < len(grp); i++ {
		id := grp[i]
		session, ok := totalClientsInSession[roomID][id]
		if !ok {
			log.Fatal("user is not in the session")
		}
		serverResponse := model.ServerResponse{ResponseType: "total", ID: clientID, ConnectedClients: connectedClients, ClientInfo: clientInfo}
		jsonReponse, err := json.Marshal(&serverResponse)
		if err != nil {
			log.Fatal("cant parse json response")
		}
		session.Write([]byte(jsonReponse))
	}
}
