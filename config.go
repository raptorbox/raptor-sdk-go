package raptor

import (
	"io/ioutil"
	"os"
	"strings"
)

//Config a client configuration
type Config struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

//GetUsername return username
func (c *Config) GetUsername() string {
	return c.Username
}

//GetPassword return password
func (c *Config) GetPassword() string {
	return c.Password
}

//GetToken return token
func (c *Config) GetToken() string {
	return c.Token
}

//GetURL return URL
func (c *Config) GetURL() string {
	return c.URL
}

//ConfigFromFile load a config from a JSON file
func ConfigFromFile(src string) (*Config, error) {

	if _, err := os.Stat(src); os.IsNotExist(err) {
		return nil, err
	}

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
