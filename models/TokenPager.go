package models

import "encoding/json"

//TokenPager request information
type TokenPager struct {
	Pager
	Content []Token `json:"content"`
}

// ParseTokenPager parses the pager
func ParseTokenPager(raw []byte) (*TokenPager, error) {

	pager := TokenPager{
		Pager: Pager{},
	}
	err := json.Unmarshal(raw, &pager)
	if err != nil {
		return nil, err
	}

	return &pager, err
}
