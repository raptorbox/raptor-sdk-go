package models

//Stream a definition of a stream of data
type Stream struct {
	device *Device

	Name     string              `json:"name"`
	Channels map[string]*Channel `json:"channels"`
	DeviceID string              `json:"deviceId"`
	UserID   string              `json:"userId"`
}

//GetDevice return the device
func (s *Stream) GetDevice() *Device {
	return s.device
}

//SetDevice set the device
func (s *Stream) SetDevice(dev *Device) {
	s.device = dev
}

//CreateRecord create a data record for this stream
func (s *Stream) CreateRecord() *Record {
	r := NewRecord(s)
	return r
}
