package models

import "encoding/json"

//AppPager request information
type AppPager struct {
	Pager
	Content []App `json:"content"`
}

// ParseAppPager parses the pager
func ParseAppPager(raw []byte) (*AppPager, error) {

	pager := AppPager{
		Pager: Pager{},
	}
	err := json.Unmarshal(raw, &pager)
	if err != nil {
		return nil, err
	}

	return &pager, err
}
