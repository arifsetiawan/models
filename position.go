package models

// PositionJSON is
type PositionJSON struct {
	Type     string `json:"_type,omitempty"`
	TenantID string `json:"tenant_id,omitempty"`
	ID       string `json:"id"`
	GroupID  string `json:"group_id"`
	ParentID string `json:"parent_id"`
}

// PositionGroupJSON is
type PositionGroupJSON struct {
	Type       string `json:"_type,omitempty"`
	TenantID   string `json:"tenant_id,omitempty"`
	ID         string `json:"id"`
	GroupID    string `json:"group_id"`
	GroupName  string `json:"group_name"`
	ParentID   string `json:"parent_id"`
	ParentName string `json:"parent_name"`
	CreatedAt  string `json:"created_at,omitempty"`
}
