package raptor

import (
	"fmt"

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
func (a *Auth) GetConfig() models.Config {
	return a.container.GetConfig()
}

//GetClient return a client instance
func (a *Auth) GetClient() *models.Client {
	return a.container.GetClient()
}

//Login login a user with the provided credentials
func (a *Auth) Login() (*models.LoginState, error) {

	if a.GetConfig().GetToken() {
		a.GetClient().SetAuthorizationHeader(a.GetConfig().GetToken())
		raw, err := a.container.GetClient().Post(LOGIN, body, nil)
	} else {

		body := fmt.Sprintf(
			`{ "username": "%s", "password": "%s" }`,
			a.GetConfig().GetUsername(),
			a.GetConfig().GetPassword())

		raw, err := a.container.GetClient().Post(LOGIN, body, nil)
		if err != nil {
			return nil, err
		}

		state := models.LoginState{}
		err := a.GetClient().FromJSON(raw, &state)

		return
	}

}
