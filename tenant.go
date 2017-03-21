package models

// Tenant is a struct representation for Tenant Object
type Tenant struct {
	Type                  string               `json:"_type,omitempty"`
	TenantID              string               `json:"tenant_id" validate:"required"`
	Name                  string               `json:"tenant_name" validate:"required"`
	Description           string               `json:"tenant_description,omitempty"`
	DefaultAppCollections []TenantApplications `json:"default_app_collection"`
}

// TenantApplications is a representation for tenant's default app
type TenantApplications struct {
	AppType string `json:"app_type"`
	AppID   string `json:"app_id"`
}
