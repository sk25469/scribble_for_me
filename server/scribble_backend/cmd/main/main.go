package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/olahol/melody"
	"github.com/sk25469/scribble_backend/pkg/model"
)

func main() {
	router := gin.Default()
	mrouter := melody.New()

	// the base URL at which the client has to connect
	// renders the index.html page
	// router.GET("/", func(c *gin.Context) {
	// 	http.ServeFile(c.Writer, c.Request, "index.html")
	// })

	// the websocket server which connects by /ws,
	// it handles all the client requests and broadcasts it
	router.GET("/ws", func(c *gin.Context) {
		mrouter.HandleRequest(c.Writer, c.Request)
	})

	// every time a new session arrives, this method is triggered
	mrouter.HandleConnect(func(s *melody.Session) {
		// ss contains all the sessions
		ss, _ := mrouter.Sessions()
		siz := len(ss)

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
			s.Write([]byte("set " + info.ID + " " + info.X + " " + info.Y))
		}

		// now the new session is assigned a new id
		id := uuid.NewString()

		// and we set the initial info as the following to the current session
		// in the main server

		// Set "stores" the key, value pair for this session in the server
		s.Set("info", &model.PointInfo{ID: id, X: "0", Y: "0"})

		// write send the message to the client
		fmt.Println("writing in current client")
		err := s.Write([]byte("iam " + id + " " + strconv.Itoa(siz)))
		if err != nil {
			log.Fatal(err)
		}
		// time.Sleep(3 * time.Second)
		fmt.Println("broadcasting everywhere")
		// ss, _ = mrouter.Sessions()
		// for _, o := range ss {
		// 	o.Write([]byte("total " + strconv.Itoa(siz)))
		// }
		mrouter.BroadcastOthers([]byte("total "+strconv.Itoa(siz)), s)
	})

	// when a session disconnects, we get the info of current session
	// from the server, and then broadcasts to other session that
	// this client has been disconnected
	mrouter.HandleDisconnect(func(s *melody.Session) {
		info := s.MustGet("info").(*model.PointInfo)
		ss, _ := mrouter.Sessions()
		siz := len(ss) - 1
		mrouter.BroadcastOthers([]byte("dis "+info.ID+" "+strconv.Itoa(siz)), s)

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

			// then sends the message to all others
			mrouter.BroadcastOthers([]byte("set "+info.ID+" "+info.X+" "+info.Y), s)
			fmt.Println(info)
		}
	})

	router.Run(":5000")
}
