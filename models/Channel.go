package models

//Channel a definition of a channel in a stream of data
type Channel struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Unit string `json:"unit"`
}
