package models

// PartnerAppCreateJSON is
type PartnerAppCreateJSON struct {
	Type      string `json:"_type,omitempty"`
	TenantID  string `json:"tenant_id,omitempty"`
	ID        string `json:"id,omitempty"`
	UserEmail string `json:"user_email,omitempty"`
	AppName   string `json:"app_name,omitempty"`
	AppScope  string `json:"app_scope,omitempty"`
}

// PartnerAppJSON is
type PartnerAppJSON struct {
	Type         string `json:"_type,omitempty"`
	TenantID     string `json:"tenant_id,omitempty"`
	ID           string `json:"id,omitempty"`
	UserEmail    string `json:"user_email,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
}
