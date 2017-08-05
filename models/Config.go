package models

//Config config params interface
type Config interface {
	GetUsername() string
	GetPassword() string
	GetToken() string
	GetURL() string
}
