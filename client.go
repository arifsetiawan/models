package models

// ClientJSON is a struct representation for Client
type ClientJSON struct {
	Type          string   `json:"_type,omitempty"`
	ID            string   `json:"id,omitempty"`
	TenantID      string   `json:"tenant_id,omitempty"`
	ClientName    string   `json:"client_name,omitempty"`
	ClientSecret  string   `json:"client_secret,omitempty"`
	RedirectUris  []string `json:"redirect_uris,omitempty"`
	GrantTypes    []string `json:"grant_types,omitempty"`
	ResponseTypes []string `json:"response_types,omitempty"`
	GrantedScopes []string `json:"granted_scopes,omitempty"`
	Owner         string   `json:"owner,omitempty"`
	PolicyURI     string   `json:"policy_uri,omitempty"`
	TosURI        string   `json:"tos_uri,omitempty"`
	ClientURI     string   `json:"client_uri,omitempty"`
	LogoURI       string   `json:"logo_uri,omitempty"`
	Contacts      []string `json:"contacts,omitempty"`
	Scope         string   `json:"scope,omitempty"`
	Members       []string `json:"members,omitempty"`
}
