package raptor

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
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
	req                 *gorequest.SuperAgent
	authorizationHeader string
}

//ToJSON convert the model to JSON string
func (c *DefaultClient) ToJSON(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

//FromJSON convert a raw value to a model
func (c *DefaultClient) FromJSON(raw []byte, i interface{}) error {
	return json.Unmarshal(raw, i)
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
func (c *DefaultClient) request(opts *models.ClientOptions) *SuperAgent {

	var r *gorequest.SuperAgent
	if opts.NewClient {
		r := gorequest.New()
	} else {
		r := c.req
	}

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
func (c *DefaultClient) Get(url string, opts *models.ClientOptions) ([]bytes, error) {
	resp, bodyBytes, errs := c.request(opts).Get(url).EndBytes()
	return bodyBytes, errs
}

//Delete request
func (c *DefaultClient) Delete(url string, opts *models.ClientOptions) error {
	resp, bodyBytes, errs := c.request(opts).Delete(url).EndBytes()
	return bodyBytes, errs
}

//Post request
func (c *DefaultClient) Post(url string, json interface{}, opts *models.ClientOptions) (interface{}, error) {
	resp, bodyBytes, errs := c.request(opts).Post(url, json).EndBytes()
	return bodyBytes, errs
}

//Put request
func (c *DefaultClient) Put(url string, json interface{}, opts *models.ClientOptions) (interface{}, error) {
	resp, bodyBytes, errs := c.request(opts).Put(url, json).EndBytes()
	return bodyBytes, errs
}

//Subscribe to topic
func (c *DefaultClient) Subscribe(topic string, cb func(event Event)) error {

}

//Unsubscribe from topic
func (c *DefaultClient) Unsubscribe(topic string, cb func(event Event)) error {

}
