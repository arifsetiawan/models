package models

// GroupJSON is a struct representation for Group Data
type GroupJSON struct {
	Type             string   `json:"_type,omitempty"`
	ID               string   `json:"id"`
	GroupName        string   `json:"name"`
	GroupDescription string   `json:"description,omitempty"`
	Members          []string `json:"members"`
}
