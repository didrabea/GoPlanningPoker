package handlers

import (
	"encoding/json"
	"net/http"
	"planning-poker/models"
	"planning-poker/state"
	"planning-poker/utils"
)

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	if utils.WithCORS(w, r) {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	var req models.VoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	state.Mu.Lock()
	defer state.Mu.Unlock()

	room, ok := state.Rooms[id]
	if !ok || room == nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	participant, ok := room.Participants[req.UserID]
	if !ok || participant == nil {
		http.Error(w, "missing participant", http.StatusBadRequest)
		return
	}

	topic, ok := room.Topics[room.ActiveTopic]
	if !ok || topic == nil {
		http.Error(w, "active topic not found", http.StatusBadRequest)
		return
	}

	if topic.Votes == nil {
		topic.Votes = make(map[string]string)
	}

	topic.Votes[participant.ID] = req.Vote

	BroadcastRoom(room)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.RoomResponse{Room: room})
}

func RevealVotesHandler(w http.ResponseWriter, r *http.Request) {
	if utils.WithCORS(w, r) {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	state.Mu.Lock()

	room, ok := state.Rooms[id]

	if !ok {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	topic, ok := room.Topics[room.ActiveTopic]
	if ok {
		topic.Revealed = true
	}

	state.Mu.Unlock()
	BroadcastRoom(room)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.RoomResponse{Room: room})

}

func ResetVotesHandler(w http.ResponseWriter, r *http.Request) {
	if utils.WithCORS(w, r) {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")

	state.Mu.Lock()

	room, ok := state.Rooms[id]

	if !ok {
		state.Mu.Unlock()
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	topic, ok := room.Topics[room.ActiveTopic]
	if ok {
		topic.Revealed = false
		topic.Votes = make(map[string]string)
	}

	state.Mu.Unlock()
	BroadcastRoom(room)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.RoomResponse{Room: room})

}
