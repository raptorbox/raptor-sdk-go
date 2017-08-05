package models

//Settings for a device
type Settings struct {
	EnableEvents bool `json:"enableEvents"`
	EnableStore  bool `json:"enableStore"`
}

//Device a definition of a device
type Device struct {
	config *Config

	Name        string            `json:"name"`
	Description string            `json:"description"`
	Settings    *Settings         `json:"settings"`
	Streams     map[string]Stream `json:"streams"`
	Actions     map[string]Action `json:"actions"`
}
