package raptor

import (
	"strconv"
	"testing"
	"time"

	"github.com/raptorbox/raptor-sdk-go/models"
)

func newRole(r *Raptor) *models.Role {
	role := models.NewRole()
	role.Name = "test_" + strconv.Itoa(int(time.Now().UnixNano()))
	return role
}

func createRole(r *Raptor, t *testing.T) *models.Role {
	role := newRole(r)
	role.ID = "bad id"
	err := r.Admin().Role().Create(role)
	if err != nil {
		t.Fatal(err)
	}
	if role.ID == "bad id" {
		t.Fatal("ID has not been overwritten")
	}
	return role
}

func updateRole(role *models.Role, r *Raptor, t *testing.T) {
	err := r.Admin().Role().Update(role)
	if err != nil {
		t.Fatal(err)
	}
}

func loadRole(ID string, r *Raptor, t *testing.T) *models.Role {
	role, err := r.Admin().Role().Read(ID)
	if err != nil {
		t.Fatal(err)
	}
	return role
}

func TestRoleList(t *testing.T) {
	r := doLogin(t)
	createRole(r, t)
	p, err := r.Admin().Role().List()
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Content) == 0 {
		t.Fatal(err)
	}
}

func TestRoleCreate(t *testing.T) {
	r := getTestAdmin(t)
	createRole(r, t)
}

func TestRoleLoad(t *testing.T) {
	r := getTestAdmin(t)
	role := createRole(r, t)
	role1 := loadRole(role.ID, r, t)
	if role.ID != role1.ID {
		t.Fatal("Role id not matching")
	}
}

func TestRoleUpdate(t *testing.T) {
	r := getTestAdmin(t)
	role := createRole(r, t)
	role.Name = role.Name + "__updated"
	updateRole(role, r, t)
	role1 := loadRole(role.ID, r, t)
	if role.Name == role1.Name {
		return
	}
	t.Fatal("Name not updated property")
}
