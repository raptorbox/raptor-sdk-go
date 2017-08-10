package raptor_test

import (
	"os"
	"testing"

	log "github.com/Sirupsen/logrus"
	raptor "github.com/raptorbox/raptor-sdk-go"
)

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)
	os.Exit(m.Run())
}

func getTestClient(t *testing.T) *raptor.Raptor {

	c, err := raptor.ConfigFromFile("./test.config.json")
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}

	raptor, err := raptor.NewFromConfig(c)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}

	return raptor
}

func doLogin(t *testing.T) *raptor.Raptor {

	r := getTestClient(t)
	state, err := r.Auth().Login()
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}

	log.Printf("User logged in %s", state.User.Email)
	return r
}

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
