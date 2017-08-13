package raptor_test

import (
	"os"
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/raptorbox/raptor-sdk-go"
)

func getTestClient(t *testing.T) *raptor.Raptor {

	c, err := raptor.NewConfigFromFile("./test.config.json")
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

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)
	os.Exit(m.Run())
}
