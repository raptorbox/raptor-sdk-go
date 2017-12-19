package raptor

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
)

//CreateApp instantiate a new API client
func CreateApp(r *Raptor) *App {
	return &App{
		Raptor: r,
	}
}

//App API client
type App struct {
	Raptor *Raptor
}

//NewApp return a new App instance
func (i *App) NewApp() *models.App {

	userID := ""
	if i.Raptor.Auth().GetUser() != nil {
		userID = i.Raptor.Auth().GetUser().ID
	}

	app := models.NewApp()
	app.UserID = userID

	return app
}

//NewAppQuery return a new AppQuery instance
func (i *App) NewAppQuery() *models.AppQuery {
	return models.NewAppQuery()
}

//GetConfig return the configuration
func (i *App) GetConfig() models.Config {
	return i.Raptor.GetConfig()
}

//GetClient return a client instance
func (i *App) GetClient() models.Client {
	return i.Raptor.GetClient()
}

//Load a app by ID
func (i *App) Load(ID string) (*models.App, error) {

	raw, err := i.GetClient().Get(fmt.Sprintf(APP_GET, ID), nil)
	if err != nil {
		return nil, err
	}

	app := models.NewApp()
	err = json.Unmarshal(raw, app)
	if err != nil {
		return nil, err
	}

	return app, nil
}

//List apps accessible by an user
func (i *App) List() (*models.AppPager, error) {

	raw, err := i.GetClient().Get(APP_LIST, nil)
	if err != nil {
		return nil, err
	}

	return models.ParseAppPager(raw)
}

//Search for apps
func (i *App) Search(q *models.AppQuery) (*models.AppPager, error) {

	if q == nil {
		return nil, errors.New("Query is missing")
	}

	raw, err := i.GetClient().Post(APP_SEARCH, q, nil)
	if err != nil {
		return nil, err
	}

	return models.ParseAppPager(raw)
}

//Create a app
func (i *App) Create(app *models.App) error {

	raw, err := i.GetClient().Post(APP_CREATE, app, nil)

	if err != nil {
		return err
	}

	res := &models.App{}
	err = i.GetClient().FromJSON(raw, res)
	if err != nil {
		return err
	}

	err = app.Merge(*res)
	if err != nil {
		return err
	}

	return nil
}

//Update a app
func (i *App) Update(app *models.App) error {

	if app.ID == "" {
		return errors.New("App ID is missing")
	}

	raw, err := i.GetClient().Put(fmt.Sprintf(APP_UPDATE, app.ID), app, nil)

	if err != nil {
		return err
	}

	res := &models.App{}
	err = i.GetClient().FromJSON(raw, res)
	if err != nil {
		return err
	}

	err = app.Merge(res)
	if err != nil {
		return err
	}

	return nil
}

//DeleteByID a app by ID
func (i *App) DeleteByID(id string) error {
	err := i.GetClient().Delete(fmt.Sprintf(APP_DELETE, id), nil)
	return err
}

//Delete a app
func (i *App) Delete(app *models.App) error {

	if app.ID == "" {
		return errors.New("App ID is missing")
	}

	return i.DeleteByID(app.ID)
}
