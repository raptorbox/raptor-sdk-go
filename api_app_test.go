package raptor

import (
	"strconv"
	"testing"
	"time"

	"github.com/raptorbox/raptor-sdk-go/models"
)

func newApp(r *Raptor) *models.App {
	app := r.App().NewApp()
	app.Name = "test_" + strconv.Itoa(int(time.Now().UnixNano()))
	return app
}

func createApp(r *Raptor, t *testing.T) *models.App {
	app := newApp(r)
	app.ID = "bad id"
	err := r.App().Create(app)
	if err != nil {
		t.Fatal(err)
	}
	if app.ID == "bad id" {
		t.Fatal("ID has not been overwritten")
	}
	return app
}

func updateApp(app *models.App, r *Raptor, t *testing.T) {
	err := r.App().Update(app)
	if err != nil {
		t.Fatal(err)
	}
}

func loadApp(ID string, r *Raptor, t *testing.T) *models.App {
	app, err := r.App().Load(ID)
	if err != nil {
		t.Fatal(err)
	}
	return app
}

func TestAppList(t *testing.T) {
	r := getTestAdmin(t)
	_, err := r.App().List()
	if err != nil {
		t.Fatal(err)
	}
}

func TestAppCreate(t *testing.T) {
	r := getTestAdmin(t)
	createApp(r, t)
}

func TestAppRead(t *testing.T) {
	r := getTestAdmin(t)
	app := createApp(r, t)
	app1 := loadApp(app.ID, r, t)
	if app.ID != app1.ID {
		t.Fatal("App id not matching")
	}
}

func TestAppUpdate(t *testing.T) {
	r := getTestAdmin(t)
	app := createApp(r, t)
	name2 := app.Name + "Updated#2"
	app.Name = name2
	updateApp(app, r, t)
	app1 := loadApp(app.ID, r, t)
	if name2 != app1.Name {
		t.Fatal("Cannot find updated property")
	}
}
