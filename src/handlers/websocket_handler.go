package handlers

import (
	"net/http"
	"planning-poker/models"
	"planning-poker/utils"
)

var HubInstance = &models.Hub{
	Rooms:      make(map[string]map[*models.Client]bool),
	Broadcast:  make(chan models.RoomEvent, 100),
	Register:   make(chan *models.Client),
	Unregister: make(chan *models.Client),
}

func BroadcastRoom(room *models.Room) {
	if HubInstance == nil {
		return
	}

	HubInstance.Broadcast <- models.RoomEvent{
		RoomID: room.ID,
		Room:   room,
	}
}

func StartHub() {
	go HubInstance.Run()
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := utils.Upgrader.Upgrade(w, r, nil)

	roomID := r.URL.Query().Get("roomId")

	client := &models.Client{
		Conn:   conn,
		RoomID: roomID,
	}

	HubInstance.Register <- client

	go func() {
		defer func() {
			HubInstance.Unregister <- client
		}()

		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				break
			}
		}
	}()
}
