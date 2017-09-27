package models

// CommonJSON is
type CommonJSON struct {
	Type     string `json:"_type,omitempty"`
	TenantID string `json:"tenant_id,omitempty"`
}
