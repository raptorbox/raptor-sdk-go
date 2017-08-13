package models

//NewToken create a new Token instance
func NewToken() *Token {
	t := &Token{}
	t.Enabled = true
	return t
}

//Token a token identifier
type Token struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Token   string `json:"token"`
	Secret  string `json:"secret"`
	Enabled bool   `json:"enabled"`
	Expires int64  `json:"expires"`
}

//Merge a token with another instance
func (t *Token) Merge(t1 *Token) error {

	if t1.ID != 0 {
		t.ID = t1.ID
	}
	if t1.Name != "" {
		t.Name = t1.Name
	}
	if t1.Token != "" {
		t.Token = t1.Token
	}
	if t1.Secret != "" {
		t.Secret = t1.Secret
	}

	t.Expires = t1.Expires
	t.Enabled = t1.Enabled

	return nil
}
