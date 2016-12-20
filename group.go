package models

// GroupJSON is a struct representation for Group Data
type GroupJSON struct {
	ID               string   `json:"_id"`
	Type             string   `json:"_type,omitempty"`
	TenantID         string   `json:"tenant_id,omitempty"`
	GroupID          string   `json:"group_id"`
	GroupName        string   `json:"group_name"`
	GroupDescription string   `json:"group_description,omitempty"`
	Members          []string `json:"members"`
}
