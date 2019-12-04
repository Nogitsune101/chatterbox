package models

import "log"

// Room is a chatroom
type Room struct {
	Name string

	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

// NewRoom creates a new chatroom
func NewRoom(name string) *Room {
	return &Room{
		Name:       name,
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

// Run starts the new room
func (room *Room) Run() {
	for {
		select {

		// Registers a new user to a room, then broadcasts that user has joined the room
		case client := <-room.Register:
			room.Clients[client] = true

			// Send a response letting us know we joined the room
			clientMessage := NewMessage(MSGNOTICE, client.User, "You are blinded by light as you step into the ["+room.Name+"]")
			client.SendBuffer <- clientMessage.ToBytes()

			// Now send a broadcast letting people know we have joined the room
			broadcastMessage := NewMessage(MSGNOTICE, client.User, "["+client.User.Name+"] steps through a portal and into the room")
			// TODO - This is duplicated code, probably should be consolidated
			for broadcastClient := range room.Clients {
				if broadcastClient.User.Name != client.User.Name {
					select {
					case broadcastClient.SendBuffer <- broadcastMessage.ToBytes():
					default:
						close(broadcastClient.SendBuffer)
						delete(room.Clients, broadcastClient)
					}
				}
			}

		case client := <-room.Unregister:
			if _, ok := room.Clients[client]; ok {
				// Now send a broadcast letting people know we have joined the room
				broadcastMessage := NewMessage(MSGNOTICE, client.User, "["+client.User.Name+"] blips out of existence without another word")
				// TODO - This is duplicated code, probably should be consolidated
				for broadcastClient := range room.Clients {
					if broadcastClient.User.Name != client.User.Name {
						select {
						case broadcastClient.SendBuffer <- broadcastMessage.ToBytes():
						default:
							close(broadcastClient.SendBuffer)
							delete(room.Clients, broadcastClient)
						}
					}
				}
				log.Println("[" + client.User.Name + "] has been unregistered")
				delete(room.Clients, client)
				close(client.SendBuffer)
			}
		case message := <-room.Broadcast:
			for client := range room.Clients {
				select {
				case client.SendBuffer <- message:
				default:
					close(client.SendBuffer)
					delete(room.Clients, client)
				}
			}
		}
	}
}
