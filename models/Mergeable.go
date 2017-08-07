package models

//Mergeable interface to support merging of models
type Mergeable interface {
	Merge(m interface{}) error
}
