package models

import (
	"errors"
)

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
		Streams:    make(map[string]*Stream),
		Actions:    make(map[string]*Action),
	}
}

//Device a definition of a device
type Device struct {
	ID          string                 `json:"id,omitempty"`
	UserID      string                 `json:"userId,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description,omitempty"`
	Domain      string                 `json:"domain,omitempty"`
	Settings    *Settings              `json:"settings,omitempty"`
	Properties  map[string]interface{} `json:"properties,omitempty"`
	Streams     map[string]*Stream     `json:"streams,omitempty"`
	Actions     map[string]*Action     `json:"actions,omitempty"`
}

//EnsureReferences ensure structs proeperly point to parent structs
func (d *Device) EnsureReferences() {

	// apply internal references
	for _, s := range d.Streams {
		s.SetDevice(d)
		for _, c := range s.Channels {
			c.SetStream(s)
		}
	}
	for _, a := range d.Actions {
		a.SetDevice(d)
	}

}

//GetID return the Device ID
func (d *Device) GetID() string {
	return d.ID
}

//GetStream return a Stream by name
func (d *Device) GetStream(name string) *Stream {
	if v, ok := d.Streams[name]; ok {
		return v
	}
	return nil
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
	d.Domain = d1.Domain

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
			d.Streams = make(map[string]*Stream)
		}
		for key, val := range d1.Streams {
			d.Streams[key] = val
			val.SetDevice(d)
		}
	}

	if len(d1.Actions) > 0 {
		if d.Actions == nil {
			d.Actions = make(map[string]*Action)
		}
		for key, val := range d1.Actions {
			d.Actions[key] = val
			val.SetDevice(d)
		}
	}

	return nil
}
