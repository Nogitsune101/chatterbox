package models

import (
	"bytes"
	"log"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket Connection and the Room.
type Client struct {
	User *User

	Room *Room

	// The websocket Connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	SendBuffer chan []byte
}

// ReadPump pumps messages from the websocket Connection to the Room.
//
// The application runs readPump in a per-Connection goroutine. The application
// ensures that there is at most one reader on a Connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump() {
	defer func() {
		c.Room.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		mMessage := NewMessageFromBytes(c.User, message)

		if mMessage.Type == MSGWHISPER {

			if strings.ToLower(mMessage.To) == strings.ToLower(mMessage.From) {
				eMessage := NewMessage(MSGNOTICE, c.User, "You appear to be talking to yourself again.")
				c.SendBuffer <- eMessage.ToBytes()
			} else {
				foundUser := false
				for client := range c.Room.Clients {
					if strings.ToLower(mMessage.To) == strings.ToLower(client.User.Name) {
						client.SendBuffer <- mMessage.ToBytes()
						foundUser = true
					}
				}

				if !foundUser {
					eMessage := NewMessage(MSGNOTICE, c.User, "You talk to your imagnanary friend ["+mMessage.To+"]")
					c.SendBuffer <- eMessage.ToBytes()
				} else {
					c.SendBuffer <- mMessage.ToBytes()
				}
			}
		} else {
			c.Room.Broadcast <- mMessage.ToBytes()
		}
	}
}

// WritePump pumps messages from the Room to the websocket Connection.
//
// A goroutine running writePump is started for each Connection. The
// application ensures that there is at most one writer to a Connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.SendBuffer:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The Room closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.SendBuffer)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.SendBuffer)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
