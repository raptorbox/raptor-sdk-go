package raptor

import (
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
)

//CreateTree instantiate a new API client
func CreateTree(r *Raptor) *Tree {
	return &Tree{
		Raptor: r,
	}
}

//Tree API client
type Tree struct {
	Raptor *Raptor
}

//GetConfig return the configuration
func (s *Tree) GetConfig() models.Config {
	return s.Raptor.GetConfig()
}

//GetClient return a client instance
func (s *Tree) GetClient() models.Client {
	return s.Raptor.GetClient()
}

//List the available trees
func (s *Tree) List() ([]models.TreeNode, error) {
	raw, err := s.GetClient().Get(TREE_LIST, nil)
	if err != nil {
		return nil, err
	}

	res := make([]models.TreeNode, 0)
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Tree return the structure for a node
func (s *Tree) Tree(node *models.TreeNode) (*models.TreeNode, error) {
	raw, err := s.GetClient().Get(fmt.Sprintf(TREE_GET, node.ID), nil)
	if err != nil {
		return nil, err
	}

	res := models.TreeNode{}
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

//Children return the children for a node
func (s *Tree) Children(node *models.TreeNode) (*models.TreeNode, error) {
	raw, err := s.GetClient().Get(fmt.Sprintf(TREE_CHILDREN, node.ID), nil)
	if err != nil {
		return nil, err
	}

	res := models.TreeNode{}
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

//Add a leaf nodes to a branch node
func (s *Tree) Add(node *models.TreeNode, childs []models.TreeNode) error {
	_, err := s.GetClient().Post(fmt.Sprintf(TREE_ADD, node.ID), childs, nil)
	return err
}

//Create a new root node
func (s *Tree) Create(node *models.TreeNode) (*models.TreeNode, error) {

	raw, err := s.GetClient().Post(TREE_CREATE, node, nil)
	if err != nil {
		return nil, err
	}

	res := models.TreeNode{}
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

//Delete a node
func (s *Tree) Delete(node *models.TreeNode) error {
	err := s.GetClient().Delete(fmt.Sprintf(TREE_REMOVE, node.ID), nil)
	return err
}
