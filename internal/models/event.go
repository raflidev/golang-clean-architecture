package models

type Event struct {
	ID    int64  `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}
