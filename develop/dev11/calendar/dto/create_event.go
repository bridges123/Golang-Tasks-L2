package dto

type CreateEvent struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Date   string `json:"date"`
}
