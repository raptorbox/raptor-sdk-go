package raptor

import "github.com/raptorbox/raptor-sdk-go/models"

//CreateAdmin instantiate a new API client
func CreateAdmin(r *Raptor) *Admin {
	return &Admin{
		Raptor: r,
	}
}

//Admin API client
type Admin struct {
	Raptor *Raptor
	state  *models.LoginState
}
