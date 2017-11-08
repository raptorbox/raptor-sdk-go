package models

//NewUser create a new User instance
func NewUser() *User {
	u := new(User)
	u.Enabled = true
	return u
}

//User a user account
type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Enabled  bool     `json:"enabled"`
	Roles    []string `json:"roles"`
}

//Merge properties of two instances of an User
func (u *User) Merge(u2 *User) error {

	u.Enabled = u2.Enabled

	if u2.Email != "" {
		u.Email = u2.Email
	}
	if u2.Email != "" {
		u.Password = u2.Password
	}
	if len(u2.Roles) > 0 {
		u.Roles = u2.Roles
	}
	if u2.ID != "" {
		u.ID = u2.ID
	}
	if u2.Username != "" {
		u.Username = u2.Username
	}

	return nil
}
