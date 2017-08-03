package models

//Stream a definition of a stream of data
type Stream struct {
	Name     string             `json:"name"`
	Channels map[string]Channel `json:"channels"`
	DeviceID string             `json:"deviceId"`
	UserID   string             `json:"userId"`
}
