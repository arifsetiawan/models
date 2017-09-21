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
	TermsOfServiceURI string   `json:"tos_uri" gorethink:"tos_uri"`
	ClientURI         string   `json:"client_uri,omitempty"`
	LogoURI           string   `json:"logo_uri,omitempty"`
	Contacts          []string `json:"contacts,omitempty"`
	Public            bool     `json:"public" gorethink:"public"`
	Members           []string `json:"members,omitempty"`
	CreatedAt         string   `json:"created_at"`
}

// Client is
type Client struct {
	ID                string   `json:"id" gorethink:"id"`
	Name              string   `json:"client_name" gorethink:"client_name"`
	Secret            string   `json:"client_secret,omitempty" gorethink:"client_secret"`
	RedirectURIs      []string `json:"redirect_uris" gorethink:"redirect_uris"`
	GrantTypes        []string `json:"grant_types" gorethink:"grant_types"`
	ResponseTypes     []string `json:"response_types" gorethink:"response_types"`
	Scope             string   `json:"scope" gorethink:"scope"`
	Owner             string   `json:"owner" gorethink:"owner"`
	PolicyURI         string   `json:"policy_uri" gorethink:"policy_uri"`
	TermsOfServiceURI string   `json:"tos_uri" gorethink:"tos_uri"`
	ClientURI         string   `json:"client_uri" gorethink:"client_uri"`
	LogoURI           string   `json:"logo_uri" gorethink:"logo_uri"`
	Contacts          []string `json:"contacts" gorethink:"contacts"`
	Public            bool     `json:"public" gorethink:"public"`
}
