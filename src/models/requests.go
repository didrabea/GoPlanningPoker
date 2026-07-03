package models

type CreateRoomRequest struct {
	Name     string `json:"name"`
	UserName string `json:"userName"`
}

type JoinRoomRequest struct {
	UserName string `json:"userName"`
}
type VoteRequest struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
	Vote   string `json:"vote"`
}

type TopicRequest struct {
	ID          string  `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type ActiveTopicRequest struct {
	TopicID string `json:"topicId"`
}
