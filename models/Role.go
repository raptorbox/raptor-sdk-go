package models

//NewRole create a new Role instance
func NewRole() *Role {
	t := &Role{
		Permissions: []string{},
	}
	return t
}

//Role a Role identifier
type Role struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	Domain      string   `json:"domain"`
}

//Merge a Role with another instance
func (t *Role) Merge(t1 *Role) error {

	if t1.ID != "" {
		t.ID = t1.ID
	}
	if t1.Name != "" {
		t.Name = t1.Name
	}
	if t1.Domain != "" {
		t.Domain = t1.Domain
	}
	t.Permissions = t1.Permissions

	return nil
}
