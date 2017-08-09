package raptor

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/prometheus/common/log"
	"github.com/raptorbox/raptor-sdk-go/models"
)

//Raptor the SDK API wrapper
type Raptor struct {
	config *Config
	client *DefaultClient

	auth      *Auth
	inventory *Inventory
	stream    *Stream
	tree      *Tree
	action    *Action
	profile   *Profile
	admin     *Admin
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

//Admin handles authentication API
func (r *Raptor) Admin() *Admin {
	if r.admin == nil {
		r.admin = CreateAdmin(r)
	}
	return r.admin
}

//Auth handles authentication API
func (r *Raptor) Auth() *Auth {
	if r.auth == nil {
		r.auth = CreateAuth(r)
	}
	return r.auth
}

//Profile handles user storage API
func (r *Raptor) Profile() *Profile {
	if r.profile == nil {
		r.profile = CreateProfile(r)
	}
	return r.profile
}

//Inventory handles Inventory API
func (r *Raptor) Inventory() *Inventory {
	if r.inventory == nil {
		r.inventory = CreateInventory(r)
	}
	return r.inventory
}

//Stream handles Stream API
func (r *Raptor) Stream() *Stream {
	if r.stream == nil {
		r.stream = CreateStream(r)
	}
	return r.stream
}

//Action handles Stream API
func (r *Raptor) Action() *Action {
	if r.action == nil {
		r.action = CreateAction(r)
	}
	return r.action
}

//Tree handles Tree API
func (r *Raptor) Tree() *Tree {
	if r.action == nil {
		r.tree = CreateTree(r)
	}
	return r.tree
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

	_, err := url.Parse(r.config.GetURL())
	if err != nil {
		log.Debugf("Cannot parse URL `%s`", r.config.GetURL())
		return err
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

//NewFromConfig create a new Raptor instance from a provided Config
func NewFromConfig(config *Config) (*Raptor, error) {
	r, err := New(config.GetURL())
	if err != nil {
		return nil, err
	}
	err = r.setConfig(config)
	return r, err
}

//NewFromToken create a new Raptor instance using token authentication
func NewFromToken(url string, token string) (*Raptor, error) {
	r, err := New(url)
	if err != nil {
		return nil, err
	}
	err = r.SetToken(token)
	return r, err
}

//NewFromCredentials create a new Raptor instance using username & password authentication
func NewFromCredentials(url string, username string, password string) (*Raptor, error) {
	r, err := New(url)
	if err != nil {
		return nil, err
	}
	err = r.SetCredentials(username, password)
	return r, err
}

//New create a new SDK instance
func New(url string) (*Raptor, error) {
	r := Raptor{}
	r.setConfig(&Config{
		url: url,
	})
	return &r, nil
}

//ToJSON convert the model to JSON string
func ToJSON(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

//FromJSON convert a raw value to a model
func FromJSON(raw []byte, i interface{}) error {
	return json.Unmarshal(raw, i)
}
