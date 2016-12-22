package models

// Tenant is a struct representation for Tenant Object
type Tenant struct {
	Type        string `json:"_type,omitempty"`
	TenantID    string `json:"tenant_id"`
	Name        string `json:"tenant_name"`
	Description string `json:"tenant_description,omitempty"`
}
