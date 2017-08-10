package models

import "errors"

//Settings for a device
type Settings struct {
	EnableEvents bool `json:"enableEvents"`
	EnableStore  bool `json:"enableStore"`
}

//NewDevice instantiate a new device
func NewDevice() *Device {
	return &Device{
		Settings: &Settings{
			EnableEvents: true,
			EnableStore:  true,
		},
		Properties: make(map[string]interface{}),
		Streams:    make(map[string]Stream),
		Actions:    make(map[string]Action),
	}
}

//Device a definition of a device
type Device struct {
	ID          string                 `json:"id"`
	UserID      string                 `json:"userId"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Settings    *Settings              `json:"settings"`
	Properties  map[string]interface{} `json:"properties"`
	Streams     map[string]Stream      `json:"streams"`
	Actions     map[string]Action      `json:"actions"`
}

//GetID return the Device ID
func (d *Device) GetID() string {
	return d.ID
}

//Merge two device definitions, implements Mergeable
func (d *Device) Merge(raw interface{}) error {

	d1, ok := raw.(Device)
	if !ok {
		return errors.New("Cannot cast to Device model")
	}

	d.ID = d1.ID
	d.UserID = d1.UserID
	d.Name = d1.Name
	d.Description = d1.Description

	if d1.Settings != nil {

		if d.Settings == nil {
			d.Settings = &Settings{}
		}

		d.Settings.EnableEvents = d1.Settings.EnableEvents
		d.Settings.EnableStore = d1.Settings.EnableStore
	}

	if len(d1.Properties) > 0 {
		if d.Properties == nil {
			d.Properties = make(map[string]interface{})
		}
		for key, val := range d1.Properties {
			d.Properties[key] = val
		}
	}

	if len(d1.Streams) > 0 {
		if d.Streams == nil {
			d.Streams = make(map[string]Stream)
		}
		for key, val := range d1.Streams {
			d.Streams[key] = val
			val.SetDevice(d)
		}
	}

	if len(d1.Actions) > 0 {
		if d.Actions == nil {
			d.Actions = make(map[string]Action)
		}
		for key, val := range d1.Actions {
			d.Actions[key] = val
			val.SetDevice(d)
		}
	}

	return nil
}
