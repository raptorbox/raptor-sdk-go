package models

//DeviceQuery format a query for device search
type DeviceQuery struct {
	ID          *TextQuery `json:"id,omitempty"`
	AppID       *TextQuery `json:"appId,omitempty"`
	Name        *TextQuery `json:"name,omitempty"`
	Description *TextQuery `json:"description,omitempty"`
	Domain      *TextQuery `json:"domain,omitempty"`
	Properties  *MapQuery  `json:"properties,omitempty"`
}

//NewDeviceQuery instantiate a device query
func NewDeviceQuery() *DeviceQuery {
	q := &DeviceQuery{
		Name:        NewTextQuery(),
		Description: NewTextQuery(),
		Domain:      NewTextQuery(),
		ID:          NewTextQuery(),
		AppID:       NewTextQuery(),
		Properties:  NewMapQuery(),
	}
	return q
}
