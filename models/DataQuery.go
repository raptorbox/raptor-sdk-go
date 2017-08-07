package models

//Metric A spatial metric
type Metric string

const (
	//KILOMETERS metric
	KILOMETERS Metric = "KILOMETERS"
	//MILES metric
	MILES = "MILES"
	//NEUTRAL metric
	NEUTRAL = "NEUTRAL"
)

//IQuery generic query interface
type IQuery interface {
	GetQuery() interface{}
}

//DistanceGeoQuery format a query for geo-spatial distance search in a field
type DistanceGeoQuery struct {
	Center GeoPoint `json:"center"`
	Radius float64  `json:"radius"`
	Unit   Metric   `json:"unit"`
}

//GetQuery return the query
func (q *DistanceGeoQuery) GetQuery() interface{} {
	return q
}

//BoundingBoxGeoQuery format a query for geo-spatial bounding-box search in a field
type BoundingBoxGeoQuery struct {
	NorthWest GeoPoint `json:"northwest"`
	SouthWest GeoPoint `json:"southwest"`
}

//GetQuery return the query
func (q *BoundingBoxGeoQuery) GetQuery() interface{} {
	return q
}

//GeoQuery format a query for geo-spatial search in a field
type GeoQuery struct {
	Distance    *DistanceGeoQuery    `json:"distance"`
	BoundingBox *BoundingBoxGeoQuery `json:"boundingBox"`
}

//GetQuery return the query
func (q *GeoQuery) GetQuery() interface{} {
	return q
}

//NumberQuery format a query for number search in a field
type NumberQuery struct {
	Between []float64 `json:"between"`
}

//GetQuery return the query
func (q *NumberQuery) GetQuery() interface{} {
	return q
}

//TextQuery format a query for text search in a field
type TextQuery struct {
	StartWith string   `json:"startWith"`
	EndWith   string   `json:"endWith"`
	Contains  string   `json:"contains"`
	Equals    string   `json:"equals"`
	In        []string `json:"in"`
}

//GetQuery return the query
func (q *TextQuery) GetQuery() interface{} {
	return q
}

//MapQuery format a query for map search in a field
type MapQuery struct {
	ContainsKey   string                 `json:"containsKey"`
	ContainsValue interface{}            `json:"containsValue"`
	Has           map[string]interface{} `json:"has"`
}

//GetQuery return the query
func (q *MapQuery) GetQuery() interface{} {
	return q
}

//DataQuery format a query for device search
type DataQuery struct {
	Timestamp NumberQuery       `json:"timestamp"`
	Channels  map[string]IQuery `json:"channels"`
	Location  GeoQuery          `json:"location"`
	StreamID  string            `json:"streamId"`
}

//NewDataQuery instantiate a device query
func NewDataQuery() *DataQuery {
	return &DataQuery{}
}
