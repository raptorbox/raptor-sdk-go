package client

// Event interface for event message
type Event interface {
	GetName() string
	GetData() interface{}
}

//Options options for a client request
type Options struct {
	RepeatTimes uint16
	RepeatWait  uint16
	Timeout     uint16
}

//IClient restful/pub-sub interface
type IClient interface {
	Get(url string, opts *Options) (interface{}, error)
	Delete(url string, opts *Options) error
	Post(url string, json interface{}, opts *Options) (interface{}, error)
	Put(url string, json interface{}, opts *Options) (interface{}, error)

	Subscribe(topic string, cb func(event Event))
	Unsubscribe(topic string, cb func(event Event))
}
