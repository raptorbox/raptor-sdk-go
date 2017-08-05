package raptor

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
	"github.com/raptorbox/raptor-sdk-go/models"
)

//NewDefaultClient initialize a default client
func NewDefaultClient(c *Raptor) *DefaultClient {
	return &DefaultClient{
		Raptor: c,
	}
}

//DefaultClient IClient default implementation
type DefaultClient struct {
	Raptor              *Raptor
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

//GetConfig return the configuration
func (c *DefaultClient) GetConfig() *Config {
	return c.Raptor.GetConfig()
}

//GetClient return a client instance
func (c *DefaultClient) GetClient() models.Client {
	return c
}

//SetAuthorizationHeader sign the request with provided token
func (c *DefaultClient) SetAuthorizationHeader(token string) {
	c.authorizationHeader = token
}

//request
func (c *DefaultClient) request(opts *models.ClientOptions) *gorequest.SuperAgent {

	var r *gorequest.SuperAgent
	if opts.NewClient {
		r = gorequest.New()
	} else {
		r = c.req
	}

	r.Set("Content-Type", "application/json")
	if c.authorizationHeader != "" {
		r.Set("Authorization", c.authorizationHeader)
	}

	if opts.Timeout > 0 {
		r.Timeout(opts.Timeout)
	}

	if opts.RetryTime > 0 {
		r.Retry(opts.RetryCount, opts.RetryTime, opts.RetryStatusCode)
	}

	return r
}

func handleErrors(errs []error) error {
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}

//Get request
func (c *DefaultClient) Get(url string, opts *models.ClientOptions) ([]byte, error) {
	_, responseBody, errs := c.request(opts).Get(url).EndBytes()
	return responseBody, handleErrors(errs)
}

//Delete request
func (c *DefaultClient) Delete(url string, opts *models.ClientOptions) error {
	_, _, errs := c.request(opts).Delete(url).EndBytes()
	return handleErrors(errs)
}

//Post request
func (c *DefaultClient) Post(url string, json interface{}, opts *models.ClientOptions) ([]byte, error) {
	_, responseBody, errs := c.request(opts).Post(url).Send(json).EndBytes()
	return responseBody, handleErrors(errs)
}

//Put request
func (c *DefaultClient) Put(url string, json interface{}, opts *models.ClientOptions) ([]byte, error) {
	_, responseBody, errs := c.request(opts).Put(url).Send(json).EndBytes()
	return responseBody, handleErrors(errs)
}

//Subscribe to topic
func (c *DefaultClient) Subscribe(topic string, cb func(event models.Event)) error {
	return nil
}

//Unsubscribe from topic
func (c *DefaultClient) Unsubscribe(topic string, cb func(event models.Event)) error {
	return nil
}
