package models

//DeviceQuery format a query for device search
type DeviceQuery struct {
	ID          TextQuery `json:"id"`
	Name        TextQuery `json:"name"`
	Description TextQuery `json:"description"`
	Properties  MapQuery  `json:"properties"`
}

//NewDeviceQuery instantiate a device query
func NewDeviceQuery() *DeviceQuery {
	return &DeviceQuery{}
}
