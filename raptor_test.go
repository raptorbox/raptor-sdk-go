package raptor

import "testing"

func NewRaptorTest(t *testing.T) {

	raptor, err := NewRaptorWithCredentials("http://raptor.local", "admin", "admin")
	if err != nil {
		t.FailNow()
		return
	}

	raptor.Auth().Login()

}
