package models

import (
	"strconv"
	"strings"
)

var separator = "\x00"

// MessageType indicates the type of message being sent
type MessageType int

// Message Types
const (
	MSGERROR MessageType = iota + 1
	MSGWARNING
	MSGNOTICE
	MSGINFO
	MSGWHISPER
	MSGEMOTE
	MSGSAY
)

// Reverse Message Types (golang rolls like this)
var rMessageTypes = map[string]MessageType{
	"1": MSGERROR,
	"2": MSGWARNING,
	"3": MSGNOTICE,
	"4": MSGINFO,
	"5": MSGWHISPER,
	"6": MSGEMOTE,
	"7": MSGSAY,
}

// Message class handles message conversions
type Message struct {
	Type    MessageType
	From    string
	To      string
	Message string
}

// NewMessage is used to initialize new message objects
func NewMessage(mType MessageType, user *User, message string) *Message {
	return &Message{
		Type:    mType,
		From:    user.Name,
		Message: message,
	}
}

// NewMessageFromBytes is used to initialize new message objects from incoming web byte stream
func NewMessageFromBytes(user *User, bMessage []byte) *Message {
	message := Message{}
	message.Parse(bMessage)
	message.From = strings.ToLower(user.Name)

	return &message
}

// Parse reads a byte stream into the message
func (m *Message) Parse(bMessage []byte) {
	messageParts := strings.Split(string(bMessage), separator)
	m.Type = rMessageTypes[messageParts[0]]
	if m.Type == MSGWHISPER {
		m.To = messageParts[1]
		m.Message = messageParts[2]
	} else {
		m.Message = messageParts[1]
	}
}

// ToBytes converts a message to a new byte array to feed to the web socket pump
func (m *Message) ToBytes() []byte {
	sType := strconv.Itoa(int(m.Type))
	if m.Type == MSGWHISPER {
		return []byte(sType + separator + m.From + separator + m.To + separator + m.Message)
	} else {
		return []byte(sType + separator + m.From + separator + m.Message)
	}
}
