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

//GetUser return the authenticated user
func (a *Auth) GetUser() *models.User {
	if a.state == nil {
		return nil
	}
	return a.state.User
}

//GetToken return the authentication token
func (a *Auth) GetToken() string {
	if a.state == nil {
		return ""
	}
	return a.state.Token
}

//Login login a user with the provided credentials
func (a *Auth) Login() (*models.LoginState, error) {

	var raw []byte
	var err error

	if a.GetConfig().GetToken() != "" {
		user, err1 := a.Raptor.Admin().User().Me()
		if err1 != nil {
			return nil, err1
		}

		a.state = &models.LoginState{
			Token: a.GetConfig().GetToken(),
			User:  user,
		}

	} else {
		body := fmt.Sprintf(
			`{ "username": "%s", "password": "%s" }`,
			a.GetConfig().GetUsername(),
			a.GetConfig().GetPassword())

		raw, err = a.GetClient().Post(LOGIN, body, &models.ClientOptions{
			SkipAuthHeader: true,
		})

		if err != nil {
			return nil, err
		}

		state := &models.LoginState{}
		err = a.GetClient().FromJSON(raw, state)

		if err != nil {
			return nil, err
		}

		a.state = state

	}

	if err != nil {
		return nil, err
	}

	return a.state, nil
}

//Logout logout an user
func (a *Auth) Logout() error {
	err := a.GetClient().Delete(LOGOUT, nil)
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
