package models

//Config config params interface
type Config interface {
	GetUsername()
	GetPassword()
	GetToken()
	GetURL()
}
