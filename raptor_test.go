package raptor

import (
	"log"
	"testing"
)

func TestNewRaptor(t *testing.T) {

	raptor, err := NewRaptorWithCredentials("http://raptor.local", "admin", "admin")
	if err != nil {
		t.FailNow()
		return
	}

	state, err := raptor.Auth().Login()
	if err != nil {
		t.FailNow()
		return
	}

	log.Printf("User logged in %s", state.User.Email)

}
