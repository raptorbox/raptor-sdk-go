package raptor

import (
	"errors"
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
)

//CreateToken instantiate a new API client
func CreateToken(r *Raptor) *Token {
	return &Token{
		Raptor: r,
	}
}

//Token API client
type Token struct {
	Raptor          *Raptor
	tokenPermission Permission
}

//Permission return the Permission API
func (s *Token) Permission() Permission {
	if s.tokenPermission == nil {
		s.tokenPermission = CreateTokenPermission(s.Raptor)
	}
	return s.tokenPermission
}

//GetConfig return the configuration
func (s *Token) GetConfig() models.Config {
	return s.Raptor.GetConfig()
}

//GetClient return a client instance
func (s *Token) GetClient() models.Client {
	return s.Raptor.GetClient()
}

//List the available token for the current user
func (s *Token) List() ([]models.Token, error) {

	user := s.Raptor.Auth().GetUser()
	if user == nil {
		return nil, errors.New("Missing user")
	}

	return s.ListByUUID(user.UUID)
}

//ListByUUID the available token for a user
func (s *Token) ListByUUID(uuid string) ([]models.Token, error) {

	raw, err := s.GetClient().Get(fmt.Sprintf(TOKEN_LIST, uuid), nil)
	if err != nil {
		return nil, err
	}

	res := make([]models.Token, 0)
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Read a token
func (s *Token) Read(id string) (*models.Token, error) {
	raw, err := s.GetClient().Get(fmt.Sprintf(TOKEN_GET, id), nil)
	if err != nil {
		return nil, err
	}

	res := &models.Token{}
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Create a token
func (s *Token) Create(token *models.Token) (*models.Token, error) {
	raw, err := s.GetClient().Post(TOKEN_CREATE, token, nil)
	if err != nil {
		return nil, err
	}

	err = s.GetClient().FromJSON(raw, token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

//Update a token
func (s *Token) Update(token *models.Token) (*models.Token, error) {
	raw, err := s.GetClient().Put(fmt.Sprintf(TOKEN_UPDATE, token.ID), token, nil)
	if err != nil {
		return nil, err
	}

	err = s.GetClient().FromJSON(raw, token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

//Delete an user
func (s *Token) Delete(token *models.Token) error {
	err := s.GetClient().Delete(fmt.Sprintf(TOKEN_DELETE, token.ID), nil)
	return err
}
