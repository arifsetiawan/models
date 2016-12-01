package models

// Tenant is a struct representation for Tenant Object
type Tenant struct {
	Type        string `json:"_type"`
	TenantID    string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
