package models

//NewStream instantiatie a new stream
func NewStream(d *Device) *Stream {
	return &Stream{
		device:   d,
		Channels: make(map[string]*Channel),
	}
}

//Stream a definition of a stream of data
type Stream struct {
	device *Device

	Name     string              `json:"name,omitempty"`
	Channels map[string]*Channel `json:"channels,omitempty"`
	DeviceID string              `json:"deviceId,omitempty"`
	UserID   string              `json:"userId,omitempty"`
}

//GetChannel return the device
func (s *Stream) GetChannel(name string) *Channel {
	if c, ok := s.Channels[name]; ok {
		return c
	}
	return nil
}

//GetDevice return the device
func (s *Stream) GetDevice() *Device {
	return s.device
}

//SetDevice set the device
func (s *Stream) SetDevice(dev *Device) {
	s.device = dev
	s.DeviceID = dev.ID
	s.UserID = dev.UserID
}

//CreateRecord create a data record for this stream
func (s *Stream) CreateRecord() *Record {
	r := NewRecord(s)
	return r
}
