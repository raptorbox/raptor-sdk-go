package models

import "encoding/json"

//Pager request information
type AppPager struct {
	Pager
	Content []App `json:"content"`
}

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
