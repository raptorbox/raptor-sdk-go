package models

import "encoding/json"

//DevicePager request information
type DevicePager struct {
	Pager
	Content []Device `json:"content"`
}

//ParseDevicePager parse a device pager
func ParseDevicePager(raw []byte) (*DevicePager, error) {

	pager := DevicePager{
		Pager: Pager{},
	}
	err := json.Unmarshal(raw, &pager)
	if err != nil {
		return nil, err
	}

	return &pager, err
}
