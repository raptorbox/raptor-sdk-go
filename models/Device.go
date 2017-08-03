package models

import "encoding/json"

//Settings for a device
type Settings struct {
	EnableEvents bool `json:"enableEvents"`
	EnableStore  bool `json:"enableStore"`
}

//Device a definition of a device
type Device struct {
	config    *Config
	container Container

	Name        string            `json:"name"`
	Description string            `json:"description"`
	Settings    *Settings         `json:"settings"`
	Streams     map[string]Stream `json:"streams"`
	Actions     map[string]Action `json:"actions"`
}

//ToJSON convert the model to JSON string
func (dev *Device) ToJSON() ([]byte, error) {
	return json.Marshal(dev)
}

//FromJSON convert a raw value to a model
func (dev *Device) FromJSON(raw []byte) error {
	return json.Unmarshal(raw, dev)
}
