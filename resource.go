package models

// ResourceJSON is a strut representation for Resource
type ResourceJSON struct {
	Type        string   `json:"_type"`
	ID          string   `json:"id"`
	URI         []string `json:"uri"`
	URL         []string `json:"url"`
	Description string   `json:"description,omitempty"`
}
