package models

import "encoding/json"

//RolePager request information
type RolePager struct {
	Pager
	Content []Role `json:"content"`
}

// ParseRolePager parses the pager
func ParseRolePager(raw []byte) (*RolePager, error) {

	pager := RolePager{
		Pager: Pager{},
	}
	err := json.Unmarshal(raw, &pager)
	if err != nil {
		return nil, err
	}

	return &pager, err
}
