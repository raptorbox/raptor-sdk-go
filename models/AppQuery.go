package models

//AppQuery an app search query
type AppQuery struct {
	ID         *TextQuery `json:"id,omitempty"`
	UserID     *TextQuery `json:"userId,omitempty"`
	Users      *MapQuery  `json:"users,omitempty"`
	Name       *TextQuery `json:"name,omitempty"`
	Properties *MapQuery  `json:"properties,omitempty"`
}

//NewAppQuery instantiate a new device
func NewAppQuery() *AppQuery {
	q := &AppQuery{
		Name:       NewTextQuery(),
		UserID:     NewTextQuery(),
		ID:         NewTextQuery(),
		Properties: NewMapQuery(),
		Users:      NewMapQuery(),
	}
	return q
}
