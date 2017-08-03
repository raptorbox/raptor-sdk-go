package models

//User a user account
type User struct {
	Username string
	Password string
	Email    string
	Enabled  bool
	Roles    []string
}
