package models

import "encoding/json"

//Pager request information
type PagerDevice struct {
	Pager
	Content []Device `json:"content"`
}

func ParsePagerDevice(raw []byte) (*PagerDevice, error) {

	pager := PagerDevice{
		Pager: Pager{},
	}
	err := json.Unmarshal(raw, &pager)
	if err != nil {
		return nil, err
	}

	return &pager, err
}
