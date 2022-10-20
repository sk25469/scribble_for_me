package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/olahol/melody"
	"github.com/sk25469/scribble_backend/pkg/config"
	"github.com/sk25469/scribble_backend/pkg/model"
	"github.com/sk25469/scribble_backend/pkg/utils"
)

var connectedClients []string
var mrouter = config.GetWebSocketRouter()
var response *model.ServerResponse

// TYPES OF REQUEST
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
	// response.ResponseType = "total"
	// jsonResponse, err = json.Marshal(&response)
	// if err != nil {
	// 	log.Print("can't marshall reponse")
	// }
	// broadcasts others the new total no. of sessions
	// with the id of the new joined client
	// err = mrouter.BroadcastOthers([]byte(jsonResponse), s)
	// if err != nil {
	// 	log.Fatal(err)
	// }
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

func OnMessage(s *melody.Session, msg []byte) {
	p := strings.Split(string(msg), " ")
	if len(p) == 2 {
		// we get the info of the current session from the server
		info := s.MustGet("info").(*model.ClientInfo)

		// we assign the x and y coordinates to it,
		// every time there is some new activity on the client
		info.X = p[0]
		info.Y = p[1]
		response.ResponseType = "set"
		response.ID = info.ClientID
		response.ClientInfo = &model.ClientInfo{ClientID: info.ClientID, X: p[0], Y: p[1]}

		jsonResponse, err := json.Marshal(&response)
		if err != nil {
			log.Print("can't marshall reponse")
		}

		// then sends the message to all others
		mrouter.BroadcastOthers([]byte(jsonResponse), s)
		fmt.Println(info)
	}
}
