package models

// Event interface for event message
type Event interface {
	GetName() string
	GetData() interface{}
}

//ClientOptions ClientOptions for a client request
type ClientOptions struct {
	RetryCount      uint16
	RetryTime       uint16
	RetryStatusCode uint16
	Timeout         uint16
}

//Client restful/pub-sub interface
type Client interface {
	Get(url string, opts *ClientOptions) (interface{}, error)
	Delete(url string, opts *ClientOptions) error
	Post(url string, json interface{}, opts *ClientOptions) (interface{}, error)
	Put(url string, json interface{}, opts *ClientOptions) (interface{}, error)

	Subscribe(topic string, cb func(event Event)) error
	Unsubscribe(topic string, cb func(event Event)) error
}
