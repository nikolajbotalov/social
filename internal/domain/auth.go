package domain

type Auth struct {
	ID           string `json:"id"`
	Nickname     string `json:"nickname"`
	PasswordHash string `json:"password_hash"`
}
