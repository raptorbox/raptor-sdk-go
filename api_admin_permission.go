package raptor

import (
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
)

//SubjectType a subject for ACL enforcement
type SubjectType string

const (
	//SubjectTypeDevice device type subject
	SubjectTypeDevice SubjectType = "device"
	//SubjectTypeToken token type subject
	SubjectTypeToken = "token"
	//SubjectTypeTreeNode tree node type subject
	SubjectTypeTreeNode = "tree"
)

//CreateDevicePermission instantiate a new API client
func CreateDevicePermission(r *Raptor) Permission {
	return CreatePermission(r, SubjectTypeDevice)
}

//CreateTokenPermission instantiate a new API client
func CreateTokenPermission(r *Raptor) Permission {
	return CreatePermission(r, SubjectTypeToken)
}

//CreateTreeNodePermission instantiate a new API client
func CreateTreeNodePermission(r *Raptor) Permission {
	return CreatePermission(r, SubjectTypeTreeNode)
}

//CreatePermission instantiate a new API client specifying a subject
func CreatePermission(r *Raptor, subjectType SubjectType) Permission {
	return &GenericPermission{
		Raptor:      r,
		subjectType: subjectType,
	}
}

//Permission API client interface
type Permission interface {
	GetConfig() models.Config
	GetClient() models.Client
	Get(subjectID string) ([]string, error)
	Set(subjectID string, userID string, permissions []string) ([]string, error)
}

//GenericPermission API client abstract per subject ACL permission management
type GenericPermission struct {
	Raptor      *Raptor
	subjectType SubjectType
}

//GetConfig return the configuration
func (s *GenericPermission) GetConfig() models.Config {
	return s.Raptor.GetConfig()
}

//GetClient return a client instance
func (s *GenericPermission) GetClient() models.Client {
	return s.Raptor.GetClient()
}

//Get the available device permissions
func (s *GenericPermission) Get(subjectID string) ([]string, error) {

	raw, err := s.GetClient().Get(fmt.Sprintf(PERMISSION_GET, s.subjectType, subjectID), nil)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0)
	err = s.GetClient().FromJSON(raw, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Set the token permissions
func (s *GenericPermission) Set(subjectID string, userID string, permissions []string) ([]string, error) {

	body := make(map[string]interface{})
	body["user"] = userID
	body["permissions"] = permissions

	raw, err := s.GetClient().Put(fmt.Sprintf(PERMISSION_SET, s.subjectType, subjectID), body, nil)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0)
	err = s.GetClient().FromJSON(raw, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
