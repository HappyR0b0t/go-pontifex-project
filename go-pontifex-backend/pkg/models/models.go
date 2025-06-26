package models

type User struct {
	ID         uint
	TelegramID uint64
	Username   string
	FirstName  string
	LastName   string
	Status     string
	Role       string
	CreatedAt  string
}
