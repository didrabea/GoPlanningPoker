package models

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn   *websocket.Conn
	RoomID string
}

type Hub struct {
	Rooms      map[string]map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan RoomEvent
}

type WSMessage struct {
	Type string `json:"type"`
	Room *Room  `json:"room"`
}

func (h *Hub) Run() {
	for {
		select {

		case client := <-h.Register:
			if h.Rooms[client.RoomID] == nil {
				h.Rooms[client.RoomID] = make(map[*Client]bool)
			}
			h.Rooms[client.RoomID][client] = true

		case client := <-h.Unregister:
			if clients := h.Rooms[client.RoomID]; clients != nil {
				delete(clients, client)
				_ = client.Conn.Close()
			}

		case event := <-h.Broadcast:
			clients := h.Rooms[event.RoomID]

			if clients == nil {
				continue
			}

			msg, _ := json.Marshal(WSMessage{
				Type: "room_updated",
				Room: event.Room,
			})

			for client := range clients {
				err := client.Conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					delete(clients, client)
					_ = client.Conn.Close()
				}
			}
		}
	}
}
