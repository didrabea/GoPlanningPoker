package handlers

import (
	"encoding/json"
	"net/http"
	"planning-poker/models"
	"planning-poker/state"
	"planning-poker/utils"
)

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	if utils.WithCORS(w, r) {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//parse json & error handling
	var req models.CreateRoomRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	if req.UserName == "" {
		http.Error(w, "User name is required", http.StatusBadRequest)
		return
	}

	creator := &models.Participant{
		ID:   utils.GenerateUserID(),
		Name: req.UserName,
	}

	room := &models.Room{
		ID:   utils.GenerateRoomID(),
		Name: req.Name,
		Participants: map[string]*models.Participant{
			creator.ID: creator,
		},
		Topics: make(map[string]*models.Topic),
	}

	defaultTopicID := utils.GenerateTopicID()

	room.Topics[defaultTopicID] = &models.Topic{
		ID:       defaultTopicID,
		Votes:    make(map[string]string),
		Revealed: false,
		Title:    "Default topic",
	}

	room.ActiveTopic = defaultTopicID

	//store in memory - Lock shared memory so only ONE request can modify it at a time.
	state.Mu.Lock()
	state.Rooms[room.ID] = room
	state.Mu.Unlock()

	response := models.RoomResponse{
		Room:        room,
		Participant: creator,
	}
	BroadcastRoom(room)
	//return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
	if utils.WithCORS(w, r) {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	var req models.JoinRoomRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	participant := &models.Participant{
		ID:   utils.GenerateUserID(),
		Name: req.UserName,
	}

	state.Mu.Lock()
	room, ok := state.Rooms[id]

	if !ok {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}
	response := models.RoomResponse{
		Room:        room,
		Participant: participant,
	}
	room.Participants[participant.ID] = participant
	state.Mu.Unlock()
	BroadcastRoom(room)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

}
