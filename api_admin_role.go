package raptor

import (
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
)

//CreateRole instantiate a new API client
func CreateRole(r *Raptor) *Role {
	return &Role{
		Raptor: r,
	}
}

//Role API client
type Role struct {
	Raptor         *Raptor
	RolePermission Permission
}

//GetConfig return the configuration
func (s *Role) GetConfig() models.Config {
	return s.Raptor.GetConfig()
}

//GetClient return a client instance
func (s *Role) GetClient() models.Client {
	return s.Raptor.GetClient()
}

//List the available Role for the current user
func (s *Role) List() (*models.RolePager, error) {

	raw, err := s.GetClient().Get(ROLE_LIST, nil)
	if err != nil {
		return nil, err
	}

	return models.ParseRolePager(raw)
}

//Read a Role
func (s *Role) Read(id string) (*models.Role, error) {
	raw, err := s.GetClient().Get(fmt.Sprintf(ROLE_READ, id), nil)
	if err != nil {
		return nil, err
	}

	res := &models.Role{}
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Create a Role
func (s *Role) Create(Role *models.Role) error {
	raw, err := s.GetClient().Post(ROLE_CREATE, Role, nil)
	if err != nil {
		return err
	}

	err = s.GetClient().FromJSON(raw, Role)
	if err != nil {
		return err
	}

	return nil
}

//Update a Role
func (s *Role) Update(Role *models.Role) error {
	raw, err := s.GetClient().Put(fmt.Sprintf(ROLE_UPDATE, Role.ID), Role, nil)
	if err != nil {
		return err
	}

	err = s.GetClient().FromJSON(raw, Role)
	if err != nil {
		return err
	}

	return nil
}

//Delete an user
func (s *Role) Delete(Role *models.Role) error {
	err := s.GetClient().Delete(fmt.Sprintf(ROLE_DELETE, Role.ID), nil)
	return err
}
