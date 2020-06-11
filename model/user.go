package model

type User struct {
	UserID   int    `json:"user_id"`
	GroupID  int    `json:"group_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
