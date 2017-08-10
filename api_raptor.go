package raptor

import (
	"errors"
	"net/url"

	log "github.com/Sirupsen/logrus"
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
	return r.SetConfig(&Config{
		Username: username,
		Password: password,
	})
}

//SetToken set token
func (r *Raptor) SetToken(token string) error {
	return r.SetConfig(&Config{
		Token: token,
	})
}

//SetConfig update the configuration merging with previously set
func (r *Raptor) SetConfig(config *Config) error {

	if r.config == nil {
		r.config = &Config{}
	}

	if config.GetURL() == "" && r.config.GetURL() == "" {
		return errors.New("URL cannot be empty")
	}
	if config.GetURL() != "" {
		r.config.URL = config.GetURL()
		lastChar := len(r.config.URL) - 1
		if string(r.config.URL[lastChar]) == "/" {
			r.config.URL = r.config.URL[0:lastChar]
		}

		log.Debugf("Base URL %s", r.config.GetURL())
	}

	_, err := url.Parse(r.config.GetURL())
	if err != nil {
		log.Debugf("Cannot parse URL `%s`", r.config.GetURL())
		return err
	}

	if config.GetToken() != "" {
		r.config.Token = config.GetToken()
		r.config.Username = ""
		r.config.Password = ""
	} else {

		if config.GetUsername() != "" && config.GetPassword() != "" {

			r.config.Token = ""
			r.config.Username = config.GetUsername()
			r.config.Password = config.GetPassword()

			return nil
		}

		return errors.New("Username or Password missing")
	}

	return nil
}
