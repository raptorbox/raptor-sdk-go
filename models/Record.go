package models

//GeoPoint a geopoint coordinate
type GeoPoint struct{}

//Record a stream record
type Record struct {
	Stream    *Stream
	Channels  map[string]interface{} `json:"channels"`
	Location  *GeoPoint              `json:"location"`
	Timestamp int32                  `json:"timestamp"`
}
