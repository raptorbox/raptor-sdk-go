package raptor

import (
	"errors"

	"github.com/raptorbox/raptor-sdk-go/models"
)

//Config a client configuration
type Config struct {
	url      string
	username string
	password string
	token    string
}

//GetUsername return username
func (c *Config) GetUsername() string {
	return c.username
}

//GetPassword return password
func (c *Config) GetPassword() string {
	return c.password
}

//GetToken return token
func (c *Config) GetToken() string {
	return c.token
}

//GetURL return URL
func (c *Config) GetURL() string {
	return c.url
}

//Raptor the SDK API wrapper
type Raptor struct {
	config *Config
	client *DefaultClient

	auth      *Auth
	inventory *Inventory
}

//GetConfig return the configuration
func (r *Raptor) GetConfig() *Config {
	return r.config
}

//GetClient return a client instance
func (r *Raptor) GetClient() models.Client {
	if r.client == nil {
		r.client = NewDefaultClient(r)
	}
	return r.client
}

//Auth handles authentication API
func (r *Raptor) Auth() *Auth {
	if r.auth == nil {
		r.auth = CreateAuth(r)
	}
	return r.auth
}

//Inventory handles Inventory API
func (r *Raptor) Inventory() *Auth {
	if r.inventory == nil {
		r.inventory = CreateInventory(r)
	}
	return r.inventory
}

//SetCredentials set username and password
func (r *Raptor) SetCredentials(username string, password string) error {
	return r.setConfig(&Config{
		username: username,
		password: password,
	})
}

//SetToken set token
func (r *Raptor) SetToken(token string) error {
	return r.setConfig(&Config{
		token: token,
	})
}

//setConfig update the configuration
func (r *Raptor) setConfig(config *Config) error {

	if r.config == nil {
		r.config = &Config{}
	}

	if config.GetURL() == "" && r.config.GetURL() == "" {
		return errors.New("URL cannot be empty")
	}
	if config.GetURL() != "" {
		r.config.url = config.GetURL()
	}

	if config.GetToken() != "" {
		r.config.token = config.GetToken()
		r.config.username = ""
		r.config.password = ""
	} else {

		if config.GetUsername() != "" && config.GetPassword() != "" {

			r.config.token = ""
			r.config.username = config.GetUsername()
			r.config.password = config.GetPassword()

			return nil
		}

		return errors.New("Username or Password missing")
	}

	return nil
}

//NewRaptorWithToken create a new Raptor instance using token authentication
func NewRaptorWithToken(url string, token string) (*Raptor, error) {
	r, err := NewRaptor(url)
	if err != nil {
		return nil, err
	}
	err = r.SetToken(token)
	return r, err
}

//NewRaptorWithCredentials create a new Raptor instance using username & password authentication
func NewRaptorWithCredentials(url string, username string, password string) (*Raptor, error) {
	r, err := NewRaptor(url)
	if err != nil {
		return nil, err
	}
	err = r.SetCredentials(username, password)
	return r, err
}

//NewRaptor create a new SDK instance
func NewRaptor(url string) (*Raptor, error) {
	r := Raptor{}
	r.setConfig(&Config{
		url: url,
	})
	return &r, nil
}
