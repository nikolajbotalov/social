package domain

import (
	"errors"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Nickname  string    `json:"nickname"`
	Birthday  time.Time `json:"birthday,omitempty"`
	LastVisit time.Time `json:"last_visit"`
	Interests []string  `json:"interests"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetAllResponse struct {
	Users []User `json:"users"`
	Meta  struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Total  int `json:"total"`
	} `json:"meta"`
}

var (
	ErrInvalidPagination = errors.New("invalid pagination parameters")
	ErrUserNotFound      = errors.New("user not found")
	ErrLimitExceeded     = errors.New("limit exceeded maximum allowed value")
	ErrInvalidUserID     = errors.New("invalid user ID")
	ErrEmptyUserData     = errors.New("empty user data")
)
