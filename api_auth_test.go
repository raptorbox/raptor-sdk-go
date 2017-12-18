package raptor_test

import (
	"testing"
)

func TestLogin(t *testing.T) {
	doLogin(t)
}

func TestLogout(t *testing.T) {
	r := doLogin(t)
	err := r.Auth().Logout()
	if err != nil {
		t.Fatal(err)
	}
}

func TestRefesh(t *testing.T) {
	r := doLogin(t)
	_, err := r.Auth().Refresh()
	if err != nil {
		t.Fatal(err)
	}
}
