package raptor

//CreateAdmin instantiate a new API client
func CreateAdmin(r *Raptor) *Admin {
	return &Admin{
		Raptor: r,
	}
}

//Admin API client
type Admin struct {
	Raptor *Raptor
	user   *User
}

//User handles Stream API
func (a *Admin) User() *User {
	if a.user == nil {
		a.user = CreateUser(a.Raptor)
	}
	return a.user
}
