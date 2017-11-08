package raptor

import (
	"errors"
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
)

//CreateProfile instantiate a new API client
func CreateProfile(r *Raptor) *Profile {
	return &Profile{
		Raptor: r,
	}
}

//Profile API client
type Profile struct {
	Raptor *Raptor
}

//GetConfig return the configuration
func (s *Profile) GetConfig() models.Config {
	return s.Raptor.GetConfig()
}

//GetClient return a client instance
func (s *Profile) GetClient() models.Client {
	return s.Raptor.GetClient()
}

//Get a stored value by key
func (s *Profile) Get(key string) ([]byte, error) {

	user := s.Raptor.Auth().GetUser()
	if user == nil {
		return nil, errors.New("Missing user")
	}

	raw, err := s.GetClient().Get(fmt.Sprintf(PROFILE_GET, user.ID, key), nil)
	return raw, err
}

//GetAll retrieve all stored value for an user
func (s *Profile) GetAll() ([]byte, error) {

	user := s.Raptor.Auth().GetUser()
	if user == nil {
		return nil, errors.New("Missing user")
	}

	raw, err := s.GetClient().Get(fmt.Sprintf(PROFILE_GET_ALL, user.ID), nil)
	return raw, err
}

//Set store a value by key
func (s *Profile) Set(key string, val []byte) ([]byte, error) {

	user := s.Raptor.Auth().GetUser()
	if user == nil {
		return nil, errors.New("Missing user")
	}

	raw, err := s.GetClient().Put(fmt.Sprintf(PROFILE_SET, user.ID, key), val, nil)
	return raw, err
}
