package models

import "github.com/raptorbox/raptor-sdk-go/client"

// Container a struct that can have a container reference
type Container interface {
	GetContainer() *Container
	GetClient() *client.IClient
	GetConfig() *Config
}
