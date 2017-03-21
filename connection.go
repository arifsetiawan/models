package models

// ConnectionJSON is a struct representation for Connection
type ConnectionJSON struct {
	Type           string             `json:"_type,omitempty"`
	TenantID       string             `json:"tenant_id"`
	ClientID       []string           `json:"client_id"`
	ConnectionID   string             `json:"connection_id"`
	ConnectionName string             `json:"connection_name,omitempty" validate:"required"`
	ConnectionType string             `json:"connection_type" validate:"required"` // internal, ad, ldap, microsoft, google, mysql, postgre, etc
	ConnectionData ConnectionDataJSON `json:"connection_data,omitempty"`
}

// ConnectionDataJSON is a struct representation for Connection Data
type ConnectionDataJSON struct {
	ClientID     string   `json:"client_id,omitempty"`
	ClientSecret string   `json:"client_secret,omitempty"`
	AuthURL      string   `json:"auth_url,omitempty"`
	TokenURL     string   `json:"token_url,omitempty"`
	RedirectURL  string   `json:"redirect_url,omitempty"`
	Scopes       []string `json:"scopes,omitempty"`
	Host         string   `json:"host,omitempty"`
	Port         int      `json:"port,omitempty"`
	Username     string   `json:"username,omitempty"`
	Password     string   `json:"password,omitempty"`
	BindUsername string   `json:"bind_username,omitempty"`
	BindPassword string   `json:"bind_password,omitempty"`
	Database     string   `json:"database,omitempty"`
	BaseDN       string   `json:"base_dn,omitempty"`
}
