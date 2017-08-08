package models

//Action a definition of a triggerable action
type Action struct {
	device   *Device
	Name     string `json:"name"`
	Type     string `json:"type"`
	Unit     string `json:"unit"`
	DeviceID string `json:"deviceId"`
	UserID   string `json:"userId"`
}

//CreateStatus create a wrapper to be sent as status definition
func (a *Action) CreateStatus(status string) *ActionStatus {
	return &ActionStatus{
		action:   a,
		ActionID: a.Name,
		DeviceID: a.GetDevice().ID,
		Status:   status,
	}
}

//GetDevice return the device
func (a *Action) GetDevice() *Device {
	return a.device
}

//SetDevice set the device
func (a *Action) SetDevice(dev *Device) {
	a.device = dev
}
