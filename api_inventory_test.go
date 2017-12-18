package raptor_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/raptorbox/raptor-sdk-go"
	"github.com/raptorbox/raptor-sdk-go/models"
)

func newDevice(r *raptor.Raptor) *models.Device {
	dev := r.Inventory().NewDevice()
	dev.Name = "test_" + strconv.Itoa(int(time.Now().UnixNano()))
	return dev
}

func createDevice(r *raptor.Raptor, t *testing.T) *models.Device {
	dev := newDevice(r)
	dev.ID = "bad id"
	err := r.Inventory().Create(dev)
	if err != nil {
		t.Fatal(err)
	}
	if dev.ID == "bad id" {
		t.Fatal("ID has not been overwritten")
	}
	return dev
}

func updateDevice(dev *models.Device, r *raptor.Raptor, t *testing.T) {
	err := r.Inventory().Update(dev)
	if err != nil {
		t.Fatal(err)
	}
}

func loadDevice(ID string, r *raptor.Raptor, t *testing.T) *models.Device {
	dev, err := r.Inventory().Load(ID)
	if err != nil {
		t.Fatal(err)
	}
	return dev
}

func TestList(t *testing.T) {
	r := doLogin(t)
	_, err := r.Inventory().List()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearch(t *testing.T) {
	r := doLogin(t)
	_, err := r.Inventory().List()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreate(t *testing.T) {
	r := doLogin(t)
	createDevice(r, t)
}

func TestLoad(t *testing.T) {
	r := doLogin(t)
	dev := createDevice(r, t)
	dev1 := loadDevice(dev.ID, r, t)
	if dev.ID != dev1.ID {
		t.Fatal("Device id not matching")
	}
}

func TestUpdate(t *testing.T) {
	r := doLogin(t)
	dev := createDevice(r, t)
	dev.Properties["foo"] = "bar"
	updateDevice(dev, r, t)
	dev1 := loadDevice(dev.ID, r, t)
	if v, ok := dev1.Properties["foo"]; ok {
		if v == "bar" {
			return
		}
	}
	t.Fatal("Cannot find updated property")
}
