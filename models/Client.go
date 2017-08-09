package models

import "time"

//ClientOptions ClientOptions for a client request
type ClientOptions struct {
	RetryCount      int
	RetryTime       time.Duration
	RetryStatusCode int
	Timeout         time.Duration
	NewClient       bool
	TextPlain       bool
}

//Client restful/pub-sub interface
type Client interface {
	Get(url string, opts *ClientOptions) ([]byte, error)
	Delete(url string, opts *ClientOptions) error
	Post(url string, json interface{}, opts *ClientOptions) ([]byte, error)
	Put(url string, json interface{}, opts *ClientOptions) ([]byte, error)
	SetAuthorizationHeader(token string)
	Subscribe(topic string, cb func(event Payload)) error
	Unsubscribe(topic string, cb func(event Payload)) error
	FromJSON(raw []byte, i interface{}) error
	ToJSON(i interface{}) ([]byte, error)
}
