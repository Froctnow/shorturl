package models

type URL struct {
	ID     string `db:"id"`
	URL    string `db:"url"`
	UserID string `db:"user_id"`
}
