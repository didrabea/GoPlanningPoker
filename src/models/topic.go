package models

type Topic struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Votes       map[string]string `json:"votes"`
	Revealed    bool              `json:"revealed"`
}
