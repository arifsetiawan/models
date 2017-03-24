package models

// User is representation of data for IDP
type User struct {
	Type           string            `json:"_type,omitempty" csv:"_type,omitempty"`
	UserID         string            `json:"user_id" validate:"required"`
	TenantID       string            `json:"tenant_id" csv:"tenant_id,omitempty"`
	Username       string            `json:"username" csv:"username,omitempty" validate:"required"`
	Email          string            `json:"email" csv:"email,omitempty" validate:"required,email"`
	Name           string            `json:"name" csv:"name,omitempty" validate:"required"`
	PositionID     string            `json:"position_id,omitempty" csv:"position_id,omitempty"`
	Position       string            `json:"position,omitempty" csv:"position,omitempty"`
	Organization   string            `json:"organization,omitempty" csv:"organization,omitempty"`
	Groups         []string          `json:"groups,omitempty"`
	Apps           []ApplicationJSON `json:"apps,omitempty"`
	Identities     []Identity        `json:"identities,omitempty"`
	ProfilePicture string            `json:"profile_picture,omitempty"`
	CreatedAt      string            `json:"created_at,omitempty"`
	UpdatedAt      string            `json:"updated_at,omitempty"`
	LastIP         string            `json:"last_ip,omitempty"`
	LastLogin      string            `json:"last_login,omitempty"`
	UserMetadata   interface{}       `json:"user_metadata"`
	AppMetadata    interface{}       `json:"app_metadata"`
	IsSuspended    bool              `json:"is_suspended"`
}

// UserInternal is representation for Internal Database
type UserInternal struct {
	Type           string   `json:"_type,omitempty" csv:"_type,omitempty"`
	UserID         string   `json:"user_id"`
	TenantID       string   `json:"tenant_id,omitempty" csv:"tenant_id,omitempty"`
	Name           string   `json:"name" csv:"name,omitempty"`
	Username       string   `json:"username" csv:"username,omitempty"`
	Email          string   `json:"email" csv:"email,omitempty"`
	PositionID     string   `json:"position_id,omitempty" csv:"position_id,omitempty"`
	Position       string   `json:"position,omitempty" csv:"position,omitempty"`
	Organization   string   `json:"organization,omitempty" csv:"organization,omitempty"`
	Group          []string `json:"groups,omitempty"`
	Password       string   `json:"password" csv:"password,omitempty"`
	ProfilePicture string   `json:"profile_picture,omitempty"`
	CreatedAt      string   `json:"created_at,omitempty"`
	UpdatedAt      string   `json:"updated_at,omitempty"`
	Provider       string   `json:"provider"`
	IsSuspended    bool     `json:"is_suspended"`
}

// UserProfile is common object to store user data from provider
type UserProfile struct {
	Email        string
	Username     string
	Name         string
	Organization string
}

// Identity is extra info for user idp
type Identity struct {
	Provider     string `json:"provider"`
	Connection   string `json:"connection"`
	UserID       string `json:"user_id"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpireIn     string `json:"expire_in,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
}
