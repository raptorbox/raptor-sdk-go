package raptor

import (
	"os"
	"testing"

	log "github.com/Sirupsen/logrus"
)

func getTestClient(t *testing.T) *Raptor {

	c, err := NewConfigFromFile("./test.config.json")
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}

	raptor, err := NewFromConfig(c)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}

	return raptor
}

func doLogin(t *testing.T) *Raptor {

	r := getTestClient(t)
	_, err := r.Auth().Login()
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}

	return r
}

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)
	os.Exit(m.Run())
}
