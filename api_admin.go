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
	token  *Token
	role   *Role
}

//User handles Stream API
func (a *Admin) User() *User {
	if a.user == nil {
		a.user = CreateUser(a.Raptor)
	}
	return a.user
}

//Role handles Roles API
func (a *Admin) Role() *Role {
	if a.role == nil {
		a.role = CreateRole(a.Raptor)
	}
	return a.role
}

//Token handles Stream API
func (a *Admin) Token() *Token {
	if a.token == nil {
		a.token = CreateToken(a.Raptor)
	}
	return a.token
}
