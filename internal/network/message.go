package network

import "encoding/json"

type MessageType string

const (
	MessageTypeSet MessageType = "SET"
	MessageTypeGet MessageType = "GET"
	MessageTypeAck MessageType = "ACK"
)

type Message struct {
	Type    MessageType `json:"type"`
	Key     string      `json:"key,omitempty"`
	Value   string      `json:"value,omitempty"`
	Version int64       `json:"version,omitempty"`
	NodeID  string      `json:"node_id,omitempty"`
}

func EncodeMessage(msg Message) ([]byte, error) {
	return json.Marshal(msg)
}

func DecodeMessage(data []byte) (Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	return msg, err
}
