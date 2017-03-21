package models

// PolicyJSON is a struct representation for Policy
type PolicyJSON struct {
	Type        string      `json:"_type,omitempty"`
	TenantID    string      `json:"tenant_id,omitempty"`
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty" validate:"required"`
	Description string      `json:"description,omitempty"`
	Subjects    []string    `json:"subjects,omitempty"`
	Effect      string      `json:"effect,omitempty" validate:"required"`
	Resources   []string    `json:"resources,omitempty" validate:"gt=0,dive,required"`
	Actions     []string    `json:"actions,omitempty" validate:"gt=0,dive,required"`
	Conditions  interface{} `json:"conditions,omitempty"`
}
