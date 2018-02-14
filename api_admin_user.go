package raptor

import (
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
)

//CreateUser instantiate a new API client
func CreateUser(r *Raptor) *User {
	return &User{
		Raptor: r,
	}
}

//User API client
type User struct {
	Raptor *Raptor
}

//GetConfig return the configuration
func (s *User) GetConfig() models.Config {
	return s.Raptor.GetConfig()
}

//GetClient return a client instance
func (s *User) GetClient() models.Client {
	return s.Raptor.GetClient()
}

//List the available users
func (s *User) List() (*models.UserPager, error) {
	raw, err := s.GetClient().Get(USER_LIST, nil)
	if err != nil {
		return nil, err
	}

	// pager := models.Pager{}
	// err = s.GetClient().FromJSON(raw, &pager)
	// if err != nil {
	// 	return nil, err
	// }

	return models.ParseUserPager(raw)
}

// SearchByUsername user by name
func (s *User) SearchByUsername(username string) (*models.UserPager, error) {
	raw, err := s.GetClient().Get(USER_LIST+"?username="+username, nil)
	if err != nil {
		return nil, err
	}

	// pager := models.Pager{}
	// err = s.GetClient().FromJSON(raw, &pager)
	// if err != nil {
	// 	return nil, err
	// }

	return models.ParseUserPager(raw)
}

//Read an user
func (s *User) Read(id string) (*models.User, error) {
	raw, err := s.GetClient().Get(fmt.Sprintf(USER_GET, id), nil)
	if err != nil {
		return nil, err
	}

	res := &models.User{}
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Me Read the user identified by the token
func (s *User) Me() (*models.User, error) {
	raw, err := s.GetClient().Get(USER_GET_ME, nil)
	if err != nil {
		return nil, err
	}

	res := &models.User{}
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Create an user
func (s *User) Create(user *models.User) error {

	raw, err := s.GetClient().Post(USER_CREATE, user, nil)
	if err != nil {
		return err
	}

	err = s.GetClient().FromJSON(raw, user)
	if err != nil {
		return err
	}

	return nil
}

//Update an user
func (s *User) Update(user *models.User) error {

	raw, err := s.GetClient().Put(fmt.Sprintf(USER_UPDATE, user.ID), user, nil)
	if err != nil {
		return err
	}

	err = s.GetClient().FromJSON(raw, user)
	if err != nil {
		return err
	}

	return nil
}

//Delete an user
func (s *User) Delete(user *models.User) error {
	err := s.GetClient().Delete(fmt.Sprintf(USER_DELETE, user.ID), nil)
	return err
}
