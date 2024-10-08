package models

type User struct {
	ID          int     `json:"id"`
	Username    string  `json:"username"`
	Description string  `json:"description"`
	Photo       *string `json:"photo"`
}

type Session struct {
	UserID      int       `json:"user_id"`
	IsBanned    bool      `json:"is_banned"`
	Permissions []*string `json:"permissions"`
}
