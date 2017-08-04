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

	auth *Auth
}

//GetContainer return the container
func (r *Raptor) GetContainer() Raptor {
	return r
}

//GetConfig return the configuration
func (r *Raptor) GetConfig() *Config {
	return r.config
}

//GetClient return a client instance
func (r *Raptor) GetClient() *models.Client {
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
}

//CreateWithToken create a new Raptor instance using token authentication
func CreateWithToken(url string, token string) (*Raptor, error) {
	return NewRaptor(Config{
		URL:   url,
		Token: token,
	})
}

//CreateWithUsername create a new Raptor instance using username & password authentication
func CreateWithUsername(url string, username string, password string) (*Raptor, error) {
	return NewRaptor(Config{
		URL:      url,
		Username: username,
		Password: password,
	})
}

//NewRaptor create a new SDK instance
func NewRaptor(config Config) (*Raptor, error) {

	if config.URL == "" {
		return nil, errors.New("URL cannot be empty")
	}
	if config.Token == "" {
		config.Username = ""
		config.Password = ""
	} else {
		if config.Username == "" || config.Password == "" {
			return nil, errors.New("Username or Password missing")
		}
	}

	r := Raptor{
		config: config,
	}
	return &r, nil
}
