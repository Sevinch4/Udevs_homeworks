package users

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID
	First_name string
	Last_name  string
	Email      string
	Phone      string
	Ticket_id  string
}
