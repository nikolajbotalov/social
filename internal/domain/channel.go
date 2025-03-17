package domain

type Channel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	OwnerID string `json:"owner_id"`
	Posts   []Post `json:"posts,omitempty"`
}
