package models

// PolicyJSON is a struct representation for Policy
type PolicyJSON struct {
	Type        string      `json:"_type,omitempty"`
	TenantID    string      `json:"tenant_id,omitempty"`
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Subjects    []string    `json:"subjects,omitempty"`
	Effect      string      `json:"effect,omitempty"`
	Resources   []string    `json:"resources,omitempty"`
	Actions     []string    `json:"actions,omitempty"`
	Conditions  interface{} `json:"conditions,omitempty"`
}
