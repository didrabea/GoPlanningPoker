package main

import (
	"fmt"
	"log"
	"net/http"
	"planning-poker/handlers"
)

func main() {
	go handlers.StartHub()
	http.HandleFunc("/ws", handlers.WsHandler)
	http.HandleFunc("/api/rooms/create", handlers.CreateRoomHandler)
	http.HandleFunc("/api/rooms/join", handlers.JoinRoomHandler)
	http.HandleFunc("/api/rooms/vote", handlers.VoteHandler)
	http.HandleFunc("/api/rooms/reveal", handlers.RevealVotesHandler)
	http.HandleFunc("/api/rooms/reset", handlers.ResetVotesHandler)
	http.HandleFunc("/api/rooms/topics/adjust", handlers.AdjustTopicHandler)
	http.HandleFunc("/api/rooms/topics/set-active", handlers.ChangeActiveTopicHandler)

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
