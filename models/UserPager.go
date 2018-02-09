package models

import "encoding/json"

//UserPager request information
type UserPager struct {
	Pager
	Content []User `json:"content"`
}

// ParseUserPager parses the pager
func ParseUserPager(raw []byte) (*UserPager, error) {

	pager := UserPager{
		Pager: Pager{},
	}
	err := json.Unmarshal(raw, &pager)
	if err != nil {
		return nil, err
	}

	return &pager, err
}
