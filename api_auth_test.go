package raptor_test

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

func TestLogin(t *testing.T) {
	doLogin(t)
}

func TestLogout(t *testing.T) {
	r := doLogin(t)
	err := r.Auth().Logout()
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}
}

func TestRefesh(t *testing.T) {
	r := doLogin(t)
	_, err := r.Auth().Refresh()
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}

}
