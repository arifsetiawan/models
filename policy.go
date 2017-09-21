package models

// AllowAccess should be used as effect for policies that allow access.
const AllowAccess = "allow"

// DenyAccess should be used as effect for policies that deny access.
const DenyAccess = "deny"

// Context is
type Context map[string]interface{}

// Request is the warden's request object.
type Request struct {
	// Resource is the resource that access is requested to.
	Resource string `json:"resource"`

	// Action is the action that is requested on the resource.
	Action string `json:"action"`

	// Subejct is the subject that is requesting access.
	Subject string `json:"subject"`

	// Context is the request's environmental context.
	Context Context `json:"context"`
}

// Condition either do or do not fulfill an access request.
type Condition interface {
	// GetName returns the condition's name.
	GetName() string

	// Fulfills returns true if the request is fulfilled by the condition.
	Fulfills(interface{}, *Request) bool
}

// Conditions is a collection of conditions.
type Conditions map[string]Condition

// PolicyJSON is a struct representation for Policy
type PolicyJSON struct {
	Type        string     `json:"_type,omitempty"`
	TenantID    string     `json:"tenant_id,omitempty"`
	Name        string     `json:"name,omitempty" validate:"required"`
	ID          string     `json:"id,omitempty"`
	Description string     `json:"description,omitempty"`
	Subjects    []string   `json:"subjects,omitempty"`
	Effect      string     `json:"effect,omitempty" validate:"required"`
	Resources   []string   `json:"resources,omitempty" validate:"gt=0,dive,required"`
	Actions     []string   `json:"actions,omitempty" validate:"gt=0,dive,required"`
	Conditions  Conditions `json:"conditions,omitempty"`
	CreatedAt   string     `json:"created_at,omitempty"`
}

// GetID returns the policies id.
func (p *PolicyJSON) GetID() string {
	return p.ID
}

// GetDescription returns the policies description.
func (p *PolicyJSON) GetDescription() string {
	return p.Description
}

// GetSubjects returns the policies subjects.
func (p *PolicyJSON) GetSubjects() []string {
	return p.Subjects
}

// AllowAccess returns true if the policy effect is allow, otherwise false.
func (p *PolicyJSON) AllowAccess() bool {
	return p.Effect == AllowAccess
}

// GetEffect returns the policies effect which might be 'allow' or 'deny'.
func (p *PolicyJSON) GetEffect() string {
	return p.Effect
}

// GetResources returns the policies resources.
func (p *PolicyJSON) GetResources() []string {
	return p.Resources
}

// GetActions returns the policies actions.
func (p *PolicyJSON) GetActions() []string {
	return p.Actions
}

// GetConditions returns the policies conditions.
func (p *PolicyJSON) GetConditions() Conditions {
	return p.Conditions
}

// GetEndDelimiter returns the delimiter which identifies the end of a regular expression.
func (p *PolicyJSON) GetEndDelimiter() byte {
	return '>'
}

// GetStartDelimiter returns the delimiter which identifies the beginning of a regular expression.
func (p *PolicyJSON) GetStartDelimiter() byte {
	return '<'
}
