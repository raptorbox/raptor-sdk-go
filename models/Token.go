package models

//Token a token identifier
type Token struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Token   string `json:"token"`
	Secret  string `json:"secret"`
	Enabled bool   `json:"enabled"`
	Expires int64  `json:"expires"`
}
