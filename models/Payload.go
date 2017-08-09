package models

// Payload interface for event message
type Payload interface {
	GetType() string
	GetOp() string
}

//DevicePayload device event
type DevicePayload struct {
	Type   string  `json:"type"`
	Op     string  `json:"op"`
	UserID string  `json:"userId"`
	Device *Device `json:"device"`
}

//GetType return the type of event
func (p *DevicePayload) GetType() string {
	return p.Type
}

//GetOp return the operation of the event
func (p *DevicePayload) GetOp() string {
	return p.Op
}

//StreamPayload device event
type StreamPayload struct {
	Type     string  `json:"type"`
	Op       string  `json:"op"`
	UserID   string  `json:"userId"`
	StreamID *Device `json:"streamId"`
	Record   *Record `json:"record"`
}

//GetType return the type of event
func (p *StreamPayload) GetType() string {
	return p.Type
}

//GetOp return the operation of the event
func (p *StreamPayload) GetOp() string {
	return p.Op
}

//TreeNodePayload device event
type TreeNodePayload struct {
	Type    string    `json:"type"`
	Op      string    `json:"op"`
	UserID  string    `json:"userId"`
	Node    *TreeNode `json:"node"`
	Payload Payload   `json:"payload"`
}

//GetType return the type of event
func (p *TreeNodePayload) GetType() string {
	return p.Type
}

//GetOp return the operation of the event
func (p *TreeNodePayload) GetOp() string {
	return p.Op
}

//ActionPayload device event
type ActionPayload struct {
	Type     string `json:"type"`
	Op       string `json:"op"`
	ActionID string `json:"actionId"`
	Data     string `json:"data"`
}

//GetType return the type of event
func (p *ActionPayload) GetType() string {
	return p.Type
}

//GetOp return the operation of the event
func (p *ActionPayload) GetOp() string {
	return p.Op
}
