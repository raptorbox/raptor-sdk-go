package raptor

import (
	"encoding/json"
	"os"
)

var debugEnabled bool

func init() {
	if os.Getenv("DEBUG") != "" {
		debugEnabled = true
	}
}

//NewFromConfig create a new Raptor instance from a provided Config
func NewFromConfig(config *Config) (*Raptor, error) {
	r := &Raptor{}
	err := r.SetConfig(config)
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
	r.SetConfig(&Config{
		URL: url,
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
