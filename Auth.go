package raptor

import (
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
)

//CreateAuth instantiate a new API client
func CreateAuth(r *Raptor) *Auth {
	return &Auth{
		Raptor: r,
		state:  &models.LoginState{},
	}
}

//Auth API client
type Auth struct {
	Raptor *Raptor
	state  *models.LoginState
}

//GetConfig return the configuration
func (a *Auth) GetConfig() models.Config {
	return a.Raptor.GetConfig()
}

//GetClient return a client instance
func (a *Auth) GetClient() models.Client {
	return a.Raptor.GetClient()
}

//Login login a user with the provided credentials
func (a *Auth) Login() (*models.LoginState, error) {

	var raw []byte
	var err error

	if a.GetConfig().GetToken() != "" {
		a.GetClient().SetAuthorizationHeader(a.GetConfig().GetToken())
		raw, err = a.GetClient().Get(USER_GET_ME, nil)
	} else {

		body := fmt.Sprintf(
			`{ "username": "%s", "password": "%s" }`,
			a.GetConfig().GetUsername(),
			a.GetConfig().GetPassword())

		raw, err = a.GetClient().Post(LOGIN, body, nil)
	}

	if err != nil {
		return nil, err
	}

	state := &models.LoginState{}
	err = a.GetClient().FromJSON(raw, state)

	return state, nil
}
