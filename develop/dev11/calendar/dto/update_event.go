package dto

type UpdateEvent struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Date   string `json:"date"`
}
