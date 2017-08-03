package models

//LoginState state of a login
type LoginState struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}
