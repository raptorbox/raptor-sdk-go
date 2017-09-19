package raptor

import (
	"errors"
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/parnurzeal/gorequest"
	"github.com/raptorbox/raptor-sdk-go/models"
	debug "github.com/tj/go-debug"
)

var d = debug.Debug("raptor:client:http")

// DefaultClientOptions create default client options
func DefaultClientOptions() *models.ClientOptions {
	return &models.ClientOptions{
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
	Raptor           *Raptor
	brokerConnection *BrokerConnection
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

//request
func (c *DefaultClient) prepareRequest(method string, url string, opts *models.ClientOptions) *gorequest.SuperAgent {

	if opts == nil {
		opts = DefaultClientOptions()
	}

	r := gorequest.New()

	r.Method = method
	r.Url = c.url(url)

	log.Debugf("Performing request %s %s", r.Method, r.Url)

	if opts.TextPlain {
		r.Set("Content-Type", "text/plain")
	} else {
		r.Set("Content-Type", "application/json")
	}

	if _, ok := r.Header["Authorization"]; ok {
		delete(r.Header, "Authorization")
	}

	if !opts.SkipAuthHeader {

		authorizationToken := ""
		if c.GetConfig().GetToken() != "" {
			authorizationToken = c.GetConfig().GetToken()
		}
		if c.Raptor.Auth().GetToken() != "" {
			authorizationToken = c.Raptor.Auth().GetToken()
		}

		if authorizationToken != "" {
			d("Using token %s", authorizationToken)
			r.Set("Authorization", authorizationToken)
		}
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

	if len(errs) == 0 {
		return nil
	}

	log.Error("Request errors")
	for _, err := range errs {
		log.Warnf("- %s", err.Error())
	}

	return errs[0]
}

func (c *DefaultClient) afterRequest(opts *models.ClientOptions, response gorequest.Response, body []byte, errs []error) ([]byte, error) {

	if response == nil {
		log.Debug("Response is missing")
		return nil, errors.New("Response is missing")
	}

	d("Response %d", response.StatusCode)
	d(string(body))

	err := handleErrors(errs)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		err = fmt.Errorf("Request failed with %s: %s", response.Status, string(body))
	}

	return body, err
}

//url generate an url from basepath
func (c *DefaultClient) url(url string) string {
	return fmt.Sprintf("%s%s", c.GetConfig().GetURL(), url)
}

//Get request
func (c *DefaultClient) Get(url string, opts *models.ClientOptions) ([]byte, error) {
	response, body, errs := c.prepareRequest(gorequest.GET, url, opts).EndBytes()
	return c.afterRequest(opts, response, body, errs)
}

//Delete request
func (c *DefaultClient) Delete(url string, opts *models.ClientOptions) error {
	response, body, errs := c.prepareRequest(gorequest.DELETE, url, opts).EndBytes()
	_, err := c.afterRequest(opts, response, body, errs)
	return err
}

//Post request
func (c *DefaultClient) Post(url string, json interface{}, opts *models.ClientOptions) ([]byte, error) {
	if log.GetLevel() == log.DebugLevel {
		b, err := c.ToJSON(json)
		if err == nil {
			d("Data: %v", string(b))
		} else {
			d("Data: [ERR: %s]", err.Error())
		}
	}
	response, body, errs := c.prepareRequest(gorequest.POST, url, opts).Send(json).EndBytes()
	d("Data: %v", json)
	res, err := c.afterRequest(opts, response, body, errs)
	return res, err
}

//Put request
func (c *DefaultClient) Put(url string, json interface{}, opts *models.ClientOptions) ([]byte, error) {
	if log.GetLevel() == log.DebugLevel {
		b, err := c.ToJSON(json)
		if err == nil {
			d("Data: %v", string(b))
		} else {
			d("Data: [ERR: %s]", err.Error())
		}
	}
	response, body, errs := c.prepareRequest(gorequest.PUT, url, opts).Send(json).EndBytes()
	res, err := c.afterRequest(opts, response, body, errs)
	return res, err
}
