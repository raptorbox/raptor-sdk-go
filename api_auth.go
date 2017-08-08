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

	if err != nil {
		return nil, err
	}

	return state, nil
}

//Logout logout an user
func (a *Auth) Logout() error {

	_, err := a.GetClient().Post(LOGOUT, nil, nil)
	if err != nil {
		return err
	}

	a.state = nil

	return nil
}

//Refresh a user token
func (a *Auth) Refresh() (*models.LoginState, error) {

	raw, err := a.GetClient().Get(REFRESH_TOKEN, nil)
	if err != nil {
		return nil, err
	}

	state := &models.LoginState{}
	err = a.GetClient().FromJSON(raw, state)

	if err != nil {
		return nil, err
	}

	a.state = state

	return a.state, nil
}
