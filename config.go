package raptor

import (
	"io/ioutil"
	"strings"
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

//ConfigFromFile load a config from a JSON file
func ConfigFromFile(src string) (*Config, error) {
	c := &Config{}
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return nil, err
	}
	err = FromJSON(b, c)
	return c, err
}

//ConfigFromString load a config from a JSON string
func ConfigFromString(json string) (*Config, error) {
	c := &Config{}
	r := strings.NewReader(json)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = FromJSON(b, c)
	return c, err
}
