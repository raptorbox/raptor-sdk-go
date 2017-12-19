package models

//AppQuery an app search query
type AppQuery struct {
	UserID string   `json:"userId,omitempty"`
	Users  []string `json:"users,omitempty"`
}

//NewApp instantiate a new device
func NewAppQuery() *AppQuery {
	return &AppQuery{}
}
