package raptor

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

//NewConfigFromFile load a config from a JSON file
func NewConfigFromFile(src string) (*Config, error) {
	c := &Config{}
	err := LoadModelFromFile(src, c)
	return c, err
}

//NewConfigFromString load a config from a JSON string
func NewConfigFromString(json string) (*Config, error) {
	c := &Config{}
	err := LoadModelFromString(json, c)
	return c, err
}
