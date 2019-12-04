package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"chatterbox/server/models"
)

// Message buffers
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Globals used to track users and rooms (Not used yet)
var users = make(map[string]*models.User)
var rooms = make(map[string]*models.Room)

// WebSocketModule registers the chatterbox websocket
func WebSocketModule(router *mux.Router) {

	// Normally, I would use a JWT Token via middleware for session handling and an authentication endpoint, but this is a simple app so I'm just using the pathing to set the user
	router.HandleFunc("/room/{roomName}/{userName}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		// Params
		var user *models.User
		var room *models.Room

		// If this is a new Room, then register it
		roomName := vars["roomName"]
		if currentRoom, ok := rooms[roomName]; ok {
			room = currentRoom
		} else {
			log.Println("Created new room [" + roomName + "]")
			rooms[roomName] = models.NewRoom(roomName)
			room = rooms[roomName]

			go room.Run()
		}

		// If this is a new User, then register it
		userName := vars["userName"]
		if currentUser, ok := users[userName]; ok {
			user = currentUser
		} else {
			log.Println("Registered new user [" + userName + "]")
			users[userName] = models.NewUser(userName)
			user = users[userName]
		}

		// Handle incoming messages
		WebSocket(user, room, w, r)
	})
}

// WebSocket handles websocket requests
func WebSocket(user *models.User, room *models.Room, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &models.Client{User: user, Room: room, Conn: conn, SendBuffer: make(chan []byte, 1024)}
	client.Room.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
