package models

//ActionStatus the status of an action
type ActionStatus struct {
	action   *Action
	ID       string `json:"id"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Unit     string `json:"unit"`
	DeviceID string `json:"deviceId"`
	ActionID string `json:"actionId"`
}

//GetAction return the action
func (a *ActionStatus) GetAction() *Action {
	return a.action
}

//SetAction set the action
func (a *ActionStatus) SetAction(action *Action) {
	a.action = action
	a.ActionID = action.Name
	if action.GetDevice() != nil {
		a.DeviceID = action.GetDevice().ID
	}
}
