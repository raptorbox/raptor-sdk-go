package models

//Sort contain sort and order information
type Sort struct {
	Property   string `json:"property"`
	Direction  string `json:"direction"`
	Ascending  bool   `json:"ascending,omitempty"`
	Descending bool   `json:"descending,omitempty"`
}

//Pager request information
type Pager struct {
	First            bool        `json:"first"`
	Last             bool        `json:"last"`
	Number           int64       `json:"number"`
	TotalPages       int64       `json:"totalPages"`
	TotalElements    int64       `json:"totalElements"`
	NumberOfElements int64       `json:"numberOfElements"`
	Length           int64       `json:"length"`
	Sort             []Sort      `json:"sort"`
	Total            int64       `json:"total"`
	Size             int64       `json:"size"`
	Page             int64       `json:"page"`
	Content          interface{} `json:"content"`
}

//Len return the pager length
func (p *Pager) Len() int64 {
	return p.Length
}

//IsEmpty check if there is content
func (p *Pager) IsEmpty() bool {
	return p.Length == 0
}
