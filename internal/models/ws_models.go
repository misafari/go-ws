package models

import "github.com/gorilla/websocket"

type WsJSONResponse struct {
	Action      string   `json:"action"`
	Message     string   `json:"message"`
	MessageType string   `json:"message_type"`
	AliveUsers  []string `json:"alive_users"`
}

type WebSocketConnection struct {
	*websocket.Conn
}

type WsPayLoad struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"-"`
}
