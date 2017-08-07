package models

//Channel a definition of a channel in a stream of data
type Channel struct {
	stream *Stream

	Name string `json:"name"`
	Type string `json:"type"`
	Unit string `json:"unit"`
}

//GetStream return the stream
func (c *Channel) GetStream() *Stream {
	return c.stream
}

//SetStream set the stream
func (c *Channel) SetStream(s *Stream) {
	c.stream = s
}

//GetDevice return the device
func (c *Channel) GetDevice() *Device {
	if c.stream == nil {
		return nil
	}
	return c.stream.GetDevice()
}
