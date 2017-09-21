package models

// GroupJSON is a struct representation for Group Data
type GroupJSON struct {
	Type             string   `json:"_type,omitempty"`
	TenantID         string   `json:"tenant_id,omitempty"`
	GroupID          string   `json:"group_id"`
	GroupName        string   `json:"group_name" validate:"required"`
	GroupDescription string   `json:"group_description,omitempty"`
	Members          []string `json:"members"`
	CreatedAt        string   `json:"created_at,omitempty"`
}
