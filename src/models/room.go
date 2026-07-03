package models

type Room struct {
	ID           string                  `json:"id"`
	Name         string                  `json:"name"`
	Participants map[string]*Participant `json:"participants"`
	Topics       map[string]*Topic       `json:"topics"`
	ActiveTopic  string                  `json:"activeTopic"`
}

type RoomResponse struct {
	Room        *Room        `json:"room"`
	Participant *Participant `json:"participant,omitempty"`
}

type RoomEvent struct {
	RoomID string
	Room   *Room
}
