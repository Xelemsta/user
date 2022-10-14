package db

import (
	"time"
)

// User is the model describing a "user" in database.
type User struct {
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	Email     *string   `json:"email"`
	FirstName *string   `json:"first_name"`
	Password  *string   `json:"password"`
	ID        *string   `json:"id"`
	LastName  *string   `json:"last_name"`
	Nickname  string    `json:"nickname"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Params handles possibly optional parameters
// such as offset or limit.
type Params struct {
	Offset *int64
	Limit  *int64
}
