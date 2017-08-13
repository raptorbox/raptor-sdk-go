package models

//DeviceQuery format a query for device search
type DeviceQuery struct {
	ID          *TextQuery `json:"id,omitempty"`
	Name        *TextQuery `json:"name,omitempty"`
	Description *TextQuery `json:"description,omitempty"`
	Properties  *MapQuery  `json:"properties,omitempty"`
}

//NewDeviceQuery instantiate a device query
func NewDeviceQuery() *DeviceQuery {
	q := &DeviceQuery{
		Name:        NewTextQuery(),
		Description: NewTextQuery(),
		ID:          NewTextQuery(),
		Properties:  NewMapQuery(),
	}
	return q
}
