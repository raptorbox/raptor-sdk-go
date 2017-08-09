package raptor

import (
	"errors"
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
)

//CreateInventory instantiate a new API client
func CreateInventory(r *Raptor) *Inventory {
	return &Inventory{
		Raptor: r,
	}
}

//Inventory API client
type Inventory struct {
	Raptor           *Raptor
	devicePermission Permission
}

//Permission return the Permission API
func (i *Inventory) Permission() Permission {
	if i.devicePermission == nil {
		i.devicePermission = CreateDevicePermission(i.Raptor)
	}
	return i.devicePermission
}

//GetConfig return the configuration
func (i *Inventory) GetConfig() models.Config {
	return i.Raptor.GetConfig()
}

//GetClient return a client instance
func (i *Inventory) GetClient() models.Client {
	return i.Raptor.GetClient()
}

//List devices accessible by an user
func (i *Inventory) List() (*[]models.Device, error) {

	raw, err := i.GetClient().Get(INVENTORY_LIST, nil)

	if err != nil {
		return nil, err
	}

	res := &[]models.Device{}
	err = i.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Search for devices
func (i *Inventory) Search(q *models.DeviceQuery) (*[]models.Device, error) {

	if q == nil {
		return nil, errors.New("Query is missing")
	}

	raw, err := i.GetClient().Post(INVENTORY_SEARCH, q, nil)
	if err != nil {
		return nil, err
	}

	res := &[]models.Device{}
	err = i.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Create a device
func (i *Inventory) Create(dev *models.Device) error {

	raw, err := i.GetClient().Post(INVENTORY_CREATE, dev, nil)

	if err != nil {
		return err
	}

	res := &models.Device{}
	err = i.GetClient().FromJSON(raw, res)
	if err != nil {
		return err
	}

	dev.Merge(res)

	return nil
}

//Update a device
func (i *Inventory) Update(dev *models.Device) error {

	if dev.ID == "" {
		return errors.New("Device ID is missing")
	}

	raw, err := i.GetClient().Put(fmt.Sprintf(INVENTORY_UPDATE, dev.ID), dev, nil)

	if err != nil {
		return err
	}

	res := &models.Device{}
	err = i.GetClient().FromJSON(raw, res)
	if err != nil {
		return err
	}

	dev.Merge(res)

	return nil
}

//DeleteByID a device by ID
func (i *Inventory) DeleteByID(id string) error {
	err := i.GetClient().Delete(fmt.Sprintf(INVENTORY_DELETE, id), nil)
	return err
}

//Delete a device
func (i *Inventory) Delete(dev *models.Device) error {

	if dev.ID == "" {
		return errors.New("Device ID is missing")
	}

	return i.DeleteByID(dev.ID)
}
