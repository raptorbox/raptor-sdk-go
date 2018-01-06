package raptor

import (
	"strconv"
	"testing"
	"time"

	"github.com/raptorbox/raptor-sdk-go/models"
)

func getTestClient(t *testing.T) *Raptor {

	c, err := NewConfigFromFile("./test.config.json")
	if err != nil {
		t.Fatal(err)
	}

	raptor, err := NewFromConfig(c)
	if err != nil {
		t.Fatal(err)
	}

	return raptor
}

func doLogin(t *testing.T) *Raptor {

	r := getTestClient(t)
	_, err := r.Auth().Login()
	if err != nil {
		t.Fatal(err)
	}

	return r
}

func getTestAdmin(t *testing.T) *Raptor {
	r := doLogin(t)
	return newUser([]string{"admin"}, r, t)
}

func getTestUser(t *testing.T) *Raptor {
	r := doLogin(t)
	return newUser([]string{}, r, t)
}

func newUser(roles []string, r *Raptor, t *testing.T) *Raptor {

	user := models.NewUser()

	username := "test_" + strconv.Itoa(int(time.Now().UnixNano()))
	password := "pass_" + username
	user.Username = username
	user.Password = password
	user.Email = username + "@test.raptor.local"
	user.Roles = roles

	err := r.Admin().User().Create(user)
	if err != nil {
		t.Fatal(err)
	}

	r1, err := NewFromCredentials(r.GetConfig().GetURL(), username, password)
	if err != nil {
		t.Fatal(err)
	}

	_, err = r1.Auth().Login()
	if err != nil {
		t.Fatal(err)
	}

	return r1
}

// func TestMain(m *testing.M) {
// 	os.Exit(m.Run())
// }
