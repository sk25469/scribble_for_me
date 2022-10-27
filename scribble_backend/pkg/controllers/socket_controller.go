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
var privateRoomsMap map[string]*model.Room
var publicRoomsBasedPriorityQueue model.PriorityQueue
var publicRoomsMap map[string]*model.Room

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
	newClientInfo := model.ClientInfo{ClientID: clientID, Name: clientName}
	if clientResponse.ReponseType == "connect-new" {
		var newRoomID string

		//	if he wants to create a private room, a new key for room is created

		if clientResponse.RoomType == "private" {
			newRoomID = utils.GetKey()
			grp1, grp2 := utils.InsertClientInRoom(privateRoomsMap[newRoomID].Group1, privateRoomsMap[newRoomID].Group2, clientID)
			privateRoomsMap[newRoomID] = &model.Room{RoomID: newRoomID, Group1: grp1, Group2: grp2}
		} else {
			// If there are no public rooms, create a new room and assign client to it
			if publicRoomsBasedPriorityQueue.Len() == 0 {
				newRoomID = utils.GetKey()
				AddAndUpdatePublicRooms([]string{}, []string{}, clientID, newRoomID)

			} else {
				topRoom := publicRoomsBasedPriorityQueue.Pop().(*model.Room)
				newRoomID = topRoom.RoomID
				totalClients := len(topRoom.Group1) + len(topRoom.Group2)
				//	max 10 clients can be in a room
				//	if crosses 10, new room is formed
				if totalClients == 10 {
					newRoomID := utils.GetKey()
					AddAndUpdatePublicRooms([]string{}, []string{}, clientID, newRoomID)

				} else {
					// update the groups of the current room which has the lowest no. of clients
					newRoom := AddAndUpdatePublicRooms(topRoom.Group1, topRoom.Group2, clientID, topRoom.RoomID)
					BroadcastMessageInRoom(newRoom, &newClientInfo)
				}
			}

		}
		// update the info of the current session with its roomID and user name
		newClientInfo.RoomID = newRoomID
		s.Set("info", &newClientInfo)

		// mapping the clientID with the sessions
		totalClientsInSession[newRoomID][clientID] = s
	} else {
		// check if the roomID exists
		roomID := clientResponse.RoomID
		if _, ok := privateRoomsMap[roomID]; !ok {
			log.Fatal("given room doesn't exists")
		}
		//	check if the room already has 10 members
		totalClients := len(privateRoomsMap[roomID].Group1) + len(privateRoomsMap[roomID].Group2)
		if totalClients == 10 {
			log.Fatal("max no. of clients reached")
		}
		newClientInfo.RoomID = roomID
		s.Set("info", &newClientInfo)
		grp1, grp2 := utils.InsertClientInRoom(privateRoomsMap[roomID].Group1, privateRoomsMap[roomID].Group2, clientID)
		newRoom := model.Room{RoomID: roomID, Group1: grp1, Group2: grp2}
		privateRoomsMap[roomID] = &newRoom
		totalClientsInSession[roomID][clientID] = s
		BroadcastMessageInRoom(&newRoom, &newClientInfo)

	}

	// TODO: Create logic for updating points while drawing on screen
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

// updates the groups with equal distribution, inserts the updated room in the priority queue
// and update the public room in the map
func AddAndUpdatePublicRooms(group1, group2 []string, clientID, newRoomID string) *model.Room {
	grp1, grp2 := utils.InsertClientInRoom(group1, group2, clientID)
	newRoom := model.Room{RoomID: newRoomID, Group1: grp1, Group2: grp2}
	publicRoomsBasedPriorityQueue.Push(newRoom)
	// update the mapping for public room
	publicRoomsMap[newRoomID] = &newRoom
	return &newRoom
}

// broadcast message in a room
func BroadcastMessageInRoom(room *model.Room, clientInfo *model.ClientInfo) {
	// broadcast in group1
	clientID := clientInfo.ClientID
	for _, client := range room.Group1 {
		session := totalClientsInSession[room.RoomID][client]
		serverResponse := model.ServerResponse{ResponseType: "total", ID: clientID, ConnectedClients: connectedClients, ClientInfo: clientInfo}
		jsonReponse, err := json.Marshal(&serverResponse)
		if err != nil {
			log.Fatal("cant parse json response")
		}
		session.Write([]byte(jsonReponse))
	}

	for _, client := range room.Group2 {
		session := totalClientsInSession[room.RoomID][client]
		serverResponse := model.ServerResponse{ResponseType: "total", ID: clientID, ConnectedClients: connectedClients, ClientInfo: clientInfo}
		jsonReponse, err := json.Marshal(&serverResponse)
		if err != nil {
			log.Fatal("cant parse json response")
		}
		session.Write([]byte(jsonReponse))
	}
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
