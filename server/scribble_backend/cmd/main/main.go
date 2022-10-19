package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/olahol/melody"
	"github.com/sk25469/scribble_backend/pkg/model"
	"github.com/sk25469/scribble_backend/pkg/utils"
)

func main() {
	router := gin.Default()
	mrouter := melody.New()

	// the base URL at which the client has to connect
	// renders the index.html page
	router.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./static/index.html")
	})

	// the websocket server which connects by /ws,
	// it handles all the client requests and broadcasts it

	var connectedClients []string
	var response *model.ClientServerResponse

	router.GET("/ws", func(c *gin.Context) {
		mrouter.HandleRequest(c.Writer, c.Request)
	})

	// every time a new session arrives, this method is triggered
	mrouter.HandleConnect(func(s *melody.Session) {
		// ss contains all the sessions
		ss, _ := mrouter.Sessions()
		log.Printf("%v\n", connectedClients)

		// whenever a new session joins, we want to show the current
		// progress of all the sessions which happened earlier into the
		// current session, it basically is like
		// lets say we have a chat screen, and when a new user comes
		// they can see all the messages which happened till it joined

		for _, o := range ss {

			// we are taking the "info" from all the sessions
			info := o.MustGet("info").(*model.PointInfo)

			// and we are rendering the gophers based on all the x and y
			// coordinates of other connected sessions
			response.ResponseType = "set"
			response.ID = info.ID
			response.PointInfo = &model.PointInfo{ID: info.ID, X: info.X, Y: info.Y}

			jsonResponse, err := json.Marshal(&response)
			if err != nil {
				log.Print("can't marshall reponse")
			}
			s.Write([]byte(jsonResponse))
		}

		// now the new session is assigned a new id
		id := uuid.NewString()

		// and we set the initial info as the following to the current session
		// in the main server

		// Set "stores" the key, value pair for this session in the server
		pointInfo := model.PointInfo{ID: id, X: "0", Y: "0"}
		s.Set("info", &pointInfo)

		// write send the message to the client to set its id, and the size of the
		// fmt.Printf("after sending client %v\n", connectedClients)
		connectedClients = append(connectedClients, id)
		// current connected sessions
		response = &model.ClientServerResponse{ResponseType: "iam", ID: id, ConnectedClients: connectedClients, PointInfo: &pointInfo}
		jsonResponse, err := json.Marshal(&response)
		if err != nil {
			log.Print("can't marshall reponse")
		}

		err = s.Write([]byte(jsonResponse))

		if err != nil {
			log.Fatal(err)
		}
		response.ResponseType = "total"
		jsonResponse, err = json.Marshal(&response)
		if err != nil {
			log.Print("can't marshall reponse")
		}
		// broadcasts others the new total no. of sessions
		// with the id of the new joined client
		err = mrouter.BroadcastOthers([]byte(jsonResponse), s)
		if err != nil {
			log.Fatal(err)
		}
	})

	// when a session disconnects, we get the info of current session
	// from the server, and then broadcasts to other session that
	// this client has been disconnected
	mrouter.HandleDisconnect(func(s *melody.Session) {
		info := s.MustGet("info").(*model.PointInfo)
		var err error
		connectedClients, err = utils.Remove(connectedClients, info.ID)
		if err != nil {
			log.Fatal(err)
		}
		response.ConnectedClients = connectedClients
		response.ID = info.ID
		response.ResponseType = "dis"
		jsonResponse, err := json.Marshal(&response)
		if err != nil {
			log.Print("can't marshall reponse")
		}
		fmt.Printf("size before broadcasting %v\n", len(connectedClients))
		mrouter.BroadcastOthers([]byte(jsonResponse), s)

	})

	// every time the client sends some new message,
	// we parse it and broadcasts it to every other client
	// connected with our server
	mrouter.HandleMessage(func(s *melody.Session, msg []byte) {

		// the client sends the message as {e.pageX, e.pageY}
		p := strings.Split(string(msg), " ")
		if len(p) == 2 {
			// we get the info of the current session from the server
			info := s.MustGet("info").(*model.PointInfo)

			// we assign the x and y coordinates to it,
			// every time there is some new activity on the client
			info.X = p[0]
			info.Y = p[1]
			response.ResponseType = "set"
			response.ID = info.ID
			response.PointInfo = &model.PointInfo{ID: info.ID, X: p[0], Y: p[1]}

			jsonResponse, err := json.Marshal(&response)
			if err != nil {
				log.Print("can't marshall reponse")
			}

			// then sends the message to all others
			mrouter.BroadcastOthers([]byte(jsonResponse), s)
			fmt.Println(info)
		}
	})

	router.Run(":5000")
}
