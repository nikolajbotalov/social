package domain

import "time"

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Nickname  string    `json:"nickname"`
	Birthday  string    `json:"birthday,omitempty"`
	LastVisit time.Time `json:"last_visit"`
	Interests []string  `json:"interests"`
	Channels  []int64   `json:"channels"`
	Following []int64   `json:"following"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
