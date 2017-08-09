package raptor_test

import (
	"log"
	"testing"

	raptor "github.com/raptorbox/raptor-sdk-go"
)

func getTestClient(t *testing.T) *raptor.Raptor {

	c, err := raptor.ConfigFromFile("./test.config.json")
	if err != nil {
		t.Fatal(err)
	}

	raptor, err := raptor.NewFromConfig(c)
	if err != nil {
		t.Fatal(err)
	}

	return raptor
}

func doLogin(t *testing.T) *raptor.Raptor {

	r := getTestClient(t)
	state, err := r.Auth().Login()
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("User logged in %s", state.User.Email)
	return r
}

func TestLogin(t *testing.T) {
	doLogin(t)
}

func TestLogout(t *testing.T) {
	r := doLogin(t)
	r.Auth().Logout()
}
