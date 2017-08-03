package client

import (
	"github.com/parnurzeal/gorequest"
	"github.com/raptorbox/raptor-sdk-go/client"
	"github.com/raptorbox/raptor-sdk-go/models"
)

//NewDefaultClient initialize a default client
func NewDefaultClient(c *models.Container) *DefaultClient {
	return &DefaultClient{
		container: c,
	}
}

//DefaultClient IClient default implementation
type DefaultClient struct {
	container           *models.Container
	request             *gorequest.SuperAgent
	authorizationHeader string
}

//GetContainer return the container
func (c *DefaultClient) GetContainer() models.Container {
	return a.container
}

//GetConfig return the configuration
func (c *DefaultClient) GetConfig() *models.Config {
	return a.GetContainer().GetConfig()
}

//GetClient return a client instance
func (c *DefaultClient) GetClient() *client.IClient {
	return c
}

//SetAuthorizationHeader sign the request with provided token
func (c *DefaultClient) SetAuthorizationHeader(token string) {
	c.authorizationHeader = token
}

//request
func (c *DefaultClient) request(opts *Options) *SuperAgent {

	r := gorequest.New()

	r.Set("Content-Type", "application/json")
	if c.authorizationHeader != "" {
		r.Set("Authorization", c.authorizationHeader)
	}

	if opts.Timeout > 0 {
		r.Timeout(opts.Timeout)
	}

	if opts.RepeatTimes > 0 {
		r.Retry(opts.RetryCount, opts.RetryTime, opts.RetryStatusCode)
	}

	return r
}

//Get request
func (c *DefaultClient) Get(url string, opts *Options) ([]bytes, error) {
	resp, bodyBytes, errs := c.request(opts).Get(url).EndBytes()
	return bodyBytes, errs
}

//Delete request
func (c *DefaultClient) Delete(url string, opts *Options) error {

}

//Post request
func (c *DefaultClient) Post(url string, json interface{}, opts *Options) (interface{}, error) {

}

//Put request
func (c *DefaultClient) Put(url string, json interface{}, opts *Options) (interface{}, error) {

}

//Subscribe to topic
func (c *DefaultClient) Subscribe(topic string, cb func(event Event)) {

}

//Unsubscribe from topic
func (c *DefaultClient) Unsubscribe(topic string, cb func(event Event)) {

}
