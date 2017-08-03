package api

import (
	"github.com/raptorbox/raptor-sdk-go/client"
	"github.com/raptorbox/raptor-sdk-go/models"
)

//CreateAuth instantiate a new API client
func CreateAuth(c models.Container) *Auth {
	return &Auth{
		container: c,
	}
}

//Auth API client
type Auth struct {
	container models.Container
	state     *models.LoginState
}

//GetContainer return the container
func (a *Auth) GetContainer() models.Container {
	return a.container
}

//GetConfig return the configuration
func (a *Auth) GetConfig() *models.Config {
	return a.container.GetConfig()
}

//GetClient return a client instance
func (a *Auth) GetClient() *client.IClient {
	return a.container.GetClient()
}
