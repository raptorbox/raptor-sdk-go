package models

import "errors"

//NewApp instantiate a new device
func NewApp() *App {
	return &App{
		Roles: make([]AppRole, 0),
		Users: make([]AppUser, 0),
	}
}

//AppRole an app role
type AppRole struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions,omitempty"`
}

//AppUser an app user
type AppUser struct {
	ID      string   `json:"id,omitempty"`
	Roles   []string `json:"roles"`
	Enabled bool     `json:"enabled"`
}

//App a definition of an app
type App struct {
	ID      string    `json:"id,omitempty"`
	Name    string    `json:"name"`
	Enabled bool      `json:"enabled,omitempty"`
	UserID  string    `json:"userId"`
	Roles   []AppRole `json:"roles,omitempty"`
	Users   []AppUser `json:"users,omitempty"`
}

//GetID return the App ID
func (a *App) GetID() string {
	return a.ID
}

//Merge two instance of App
func (a *App) Merge(raw interface{}) error {

	a1, ok := raw.(App)
	if !ok {
		return errors.New("Cannot cast to App model")
	}

	a.ID = a1.ID
	a.UserID = a1.UserID
	a.Name = a1.Name
	a.Enabled = a1.Enabled

	if len(a1.Roles) > 0 {
		a.Roles = make([]AppRole, 0)
		for _, val := range a1.Roles {
			a.Roles = append(a.Roles, val)
		}
	}
	if len(a1.Users) > 0 {
		a.Users = make([]AppUser, 0)
		for _, val := range a1.Users {
			a.Users = append(a.Users, val)
		}
	}

	return nil
}
