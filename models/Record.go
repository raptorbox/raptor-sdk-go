package models

//GeoPoint a geopoint coordinate
type GeoPoint struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}

//Record a stream record
type Record struct {
	stream *Stream

	Channels  map[string]interface{} `json:"channels"`
	Location  *GeoPoint              `json:"location"`
	Timestamp int32                  `json:"timestamp"`
	StreamID  string                 `json:"streamId"`
	DeviceID  string                 `json:"deviceId"`
	UserID    string                 `json:"userId"`
}

//GetStream return the reference stream
func (r *Record) GetStream() *Stream {
	return r.stream
}

//GetChannel return the reference channel
func (r *Record) GetChannel(name string) *Channel {

	if r.GetStream() == nil {
		return nil
	}

	c, ok := r.GetStream().Channels[name]
	if !ok {
		return nil
	}

	return c
}
