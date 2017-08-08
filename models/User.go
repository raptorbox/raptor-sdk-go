package models

//User a user account
type User struct {
	UUID     string   `json:"uuid"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Enabled  bool     `json:"enabled"`
	Roles    []string `json:"roles"`
}
