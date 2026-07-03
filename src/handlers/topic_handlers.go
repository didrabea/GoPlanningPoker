package handlers

import (
	"encoding/json"
	"net/http"
	"planning-poker/models"
	"planning-poker/state"
	"planning-poker/utils"
)

func AdjustTopicHandler(w http.ResponseWriter, r *http.Request) {
	if utils.WithCORS(w, r) {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	var req models.TopicRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	state.Mu.Lock()
	defer state.Mu.Unlock()

	room, ok := state.Rooms[id]
	if !ok {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	topic, exists := room.Topics[req.ID]

	// create new topic
	if !exists {
		topic = &models.Topic{
			ID: utils.GenerateTopicID(),
			//todo add title and description
			Votes:    make(map[string]string),
			Revealed: false,
		}

		room.Topics[req.ID] = topic
	}

	// partial updates
	if req.Title != nil {
		topic.Title = *req.Title
	}

	if req.Description != nil {
		topic.Description = *req.Description
	}

	BroadcastRoom(room)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

func ChangeActiveTopicHandler(w http.ResponseWriter, r *http.Request) {
	if utils.WithCORS(w, r) {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	var req models.ActiveTopicRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	state.Mu.Lock()
	defer state.Mu.Unlock()

	room, ok := state.Rooms[id]
	if !ok {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	_, exists := room.Topics[req.TopicID]

	if !exists {
		http.Error(w, "Topic not found", http.StatusBadRequest)
		return
	}

	room.ActiveTopic = req.TopicID

	BroadcastRoom(room)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}
