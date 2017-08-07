package models

//TextQuery format a query for text search in a field
type TextQuery struct {
	StartWith string   `json:"startWith"`
	EndWith   string   `json:"endWith"`
	Contains  string   `json:"contains"`
	Equals    string   `json:"equals"`
	In        []string `json:"in"`
}

//MapQuery format a query for map search in a field
type MapQuery struct {
	ContainsKey   string                 `json:"containsKey"`
	ContainsValue interface{}            `json:"containsValue"`
	Has           map[string]interface{} `json:"has"`
}

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
