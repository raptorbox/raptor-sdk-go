package models

// Container a struct that can have a container reference
type Container interface {
	GetContainer() *Container
	GetClient() *Client
	GetConfig() Config
}
