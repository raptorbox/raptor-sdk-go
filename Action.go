package raptor

import (
	"errors"
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
)

//CreateAction instantiate a new API client
func CreateAction(r *Raptor) *Action {
	return &Action{
		Raptor: r,
	}
}

//Action API client
type Action struct {
	Raptor *Raptor
}

//GetConfig return the configuration
func (a *Action) GetConfig() models.Config {
	return a.Raptor.GetConfig()
}

//GetClient return a client instance
func (a *Action) GetClient() models.Client {
	return a.Raptor.GetClient()
}

//GetStatus fetch the action status
func (a *Action) GetStatus(action *models.Action) (*models.ActionStatus, error) {
	raw, err := a.GetClient().Get(fmt.Sprintf(STREAM_LAST_UPDATE, action.GetDevice().ID, action.Name), nil)
	if err != nil {
		return nil, err
	}
	res := &models.ActionStatus{}
	err = a.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}
	res.SetAction(action)
	return res, nil
}

//SetStatus set the action status
func (a *Action) SetStatus(status *models.ActionStatus) error {
	action := status.GetAction()
	if action == nil {
		return errors.New("Action is missing, use Action.CreateStatus to initialize an ActionStatus")
	}
	_, err := a.GetClient().Post(fmt.Sprintf(STREAM_LAST_UPDATE, action.GetDevice().ID, action.Name), status, nil)
	return err
}

//Invoke trigger an action
func (a *Action) Invoke(action *models.Action, body string) error {
	_, err := a.GetClient().Post(fmt.Sprintf(ACTION_INVOKE, action.GetDevice().ID, action.Name), body, &models.ClientOptions{
		TextPlain: true,
	})
	if err != nil {
		return err
	}
	return nil
}
