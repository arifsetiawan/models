package models

// ClientJSON is a struct representation for Client
type ClientJSON struct {
	Type              string   `json:"_type,omitempty"`
	TenantID          string   `json:"tenant_id,omitempty"`
	ID                string   `json:"id,omitempty"`
	Name              string   `json:"client_name,omitempty" validate:"required"`
	Secret            string   `json:"client_secret,omitempty"`
	RedirectUris      []string `json:"redirect_uris,omitempty" validate:"gt=0,dive,required"`
	GrantTypes        []string `json:"grant_types,omitempty" validate:"gt=0,dive,required"`
	ResponseTypes     []string `json:"response_types,omitempty" validate:"gt=0,dive,required"`
	Scope             string   `json:"scope,omitempty" validate:"required"`
	Owner             string   `json:"owner,omitempty"`
	PolicyURI         string   `json:"policy_uri,omitempty"`
	TermsOfServiceURI string   `json:"tos_uri,omitempty"`
	ClientURI         string   `json:"client_uri,omitempty"`
	LogoURI           string   `json:"logo_uri,omitempty"`
	Contacts          []string `json:"contacts,omitempty"`
	Public            bool     `json:"public,omitempty"`
	MemberUsers       []string `json:"member_users,omitempty"`
	MemberGroups      []string `json:"member_groups,omitempty"`
	CreatedAt         string   `json:"created_at"`
}

// Clear is
func (c *ClientJSON) Clear() {
	c.Secret = ""
}

// AppJSON is
type AppJSON struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ClientURI string `json:"uri,omitempty"`
	LogoURI   string `json:"logo_uri,omitempty"`
}

// NewAppJSON is
func NewAppJSON(client *ClientJSON) *AppJSON {
	return &AppJSON{
		ID:        client.ID,
		Name:      client.Name,
		ClientURI: client.ClientURI,
		LogoURI:   client.LogoURI,
	}
}
