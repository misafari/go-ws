package handlers

import (
	"fmt"
	"log"
	"sort"
	"ws/internal/models"
)

var wsIncomeMessageChannel = make(chan models.WsPayLoad, 5)
var clients = make(map[models.WebSocketConnection]string)

func ListenToWebSocket(conn *models.WebSocketConnection)  {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload models.WsPayLoad

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
		}else {
			payload.Conn = *conn
			wsIncomeMessageChannel <- payload
		}
	}
}

func WsIncomeMessageChannelListener()  {
	var response models.WsJSONResponse

	for {
		e := <- wsIncomeMessageChannel

		switch e.Action {
		case "username":
			clients[e.Conn] = e.Username
			response.Action = "list_users"
			response.AliveUsers = getAllAliveUser()
		case "left":
			response.Action = "list_users"
			delete(clients, e.Conn)
			response.AliveUsers = getAllAliveUser()
		case "broadcast":
			response.Action = "broadcast"
			response.Message = fmt.Sprintf("<strong>%s</strong>: %s", e.Username, e.Message)
		}

		broadcastToAllUser(response)
	}
}

func getAllAliveUser() []string {
	var users []string
	for _, u := range clients {
		if u != "" {
			users = append(users, u)
		}
	}
	sort.Strings(users)
	return users
}

func broadcastToAllUser(response models.WsJSONResponse) {
	for client := range clients {
		err := client.WriteJSON(&response)
		if err != nil {
			log.Println("Websocket Error", err)
			_ = client.Close()
			delete(clients, client)
		}
	}
}