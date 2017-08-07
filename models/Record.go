package models

import "time"

//GeoPoint a geopoint coordinate
type GeoPoint struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}

//NewRecord create a new record
func NewRecord(s *Stream) *Record {

	r := &Record{
		Timestamp: time.Now().Unix(),
		Channels:  make(map[string]interface{}),
	}

	if s != nil {
		r.StreamID = s.Name
		if s.GetDevice() != nil {
			d := s.GetDevice()
			r.UserID = d.UserID
			r.DeviceID = d.ID
		} else {
			r.UserID = s.UserID
			r.DeviceID = s.DeviceID
		}
	}

	return r
}

//Record a stream record
type Record struct {
	stream *Stream

	Channels  map[string]interface{} `json:"channels"`
	Location  *GeoPoint              `json:"location"`
	Timestamp int64                  `json:"timestamp"`
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
