package raptor_test

import "testing"

func TestList(t *testing.T) {
	r := doLogin(t)
	_, err := r.Inventory().List()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreate(t *testing.T) {
	r := doLogin(t)
	dev := r.Inventory().NewDevice()
	r.Inventory().Create(dev)
}
