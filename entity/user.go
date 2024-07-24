package entity

type User struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullName"`
	Username string `json:"username"`
}
