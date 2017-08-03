package models

//Action a definition of a triggerable action
type Action struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Unit     string `json:"unit"`
	DeviceID string `json:"deviceId"`
	UserID   string `json:"userId"`
}
