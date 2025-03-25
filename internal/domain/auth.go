package domain

import "time"

type Auth struct {
	ID           string `json:"id"`
	Nickname     string `json:"nickname"`
	PasswordHash string `json:"password_hash"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}
