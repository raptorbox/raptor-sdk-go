package raptor

import (
	"fmt"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/raptorbox/raptor-sdk-go/models"
)

// DefaultClientOptions create default client options
func DefaultClientOptions() *models.ClientOptions {
	return &models.ClientOptions{
		NewClient:       false,
		RetryCount:      3,
		RetryStatusCode: 400,
		RetryTime:       200,
		Timeout:         time.Second * 10,
	}
}

//NewDefaultClient initialize a default client
func NewDefaultClient(c *Raptor) *DefaultClient {
	return &DefaultClient{
		Raptor:           c,
		brokerConnection: &BrokerConnection{},
	}
}

//DefaultClient IClient default implementation
type DefaultClient struct {
	Raptor              *Raptor
	req                 *gorequest.SuperAgent
	authorizationHeader string
	brokerConnection    *BrokerConnection
}

//ToJSON convert the model to JSON string
func (c *DefaultClient) ToJSON(i interface{}) ([]byte, error) {
	return ToJSON(i)
}

//FromJSON convert a raw value to a model
func (c *DefaultClient) FromJSON(raw []byte, i interface{}) error {
	return FromJSON(raw, i)
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

	if opts == nil {
		opts = DefaultClientOptions()
	}

	var r *gorequest.SuperAgent
	if opts.NewClient {
		r = gorequest.New()
	} else {

		if c.req == nil {
			c.req = gorequest.New()
		}

		r = c.req
	}

	if opts.TextPlain {
		r.Set("Content-Type", "text/plain")
	} else {
		r.Set("Content-Type", "application/json")
	}

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

//url generate an url from basepath
func (c *DefaultClient) url(url string) string {
	return fmt.Sprintf("%s/%s", c.GetConfig().GetURL(), url)
}

//Get request
func (c *DefaultClient) Get(url string, opts *models.ClientOptions) ([]byte, error) {
	_, responseBody, errs := c.request(opts).Get(c.url(url)).EndBytes()
	return responseBody, handleErrors(errs)
}

//Delete request
func (c *DefaultClient) Delete(url string, opts *models.ClientOptions) error {
	_, _, errs := c.request(opts).Delete(c.url(url)).EndBytes()
	return handleErrors(errs)
}

//Post request
func (c *DefaultClient) Post(url string, json interface{}, opts *models.ClientOptions) ([]byte, error) {
	_, responseBody, errs := c.request(opts).Post(c.url(url)).Send(json).EndBytes()
	return responseBody, handleErrors(errs)
}

//Put request
func (c *DefaultClient) Put(url string, json interface{}, opts *models.ClientOptions) ([]byte, error) {
	_, responseBody, errs := c.request(opts).Put(c.url(url)).Send(json).EndBytes()
	return responseBody, handleErrors(errs)
}
