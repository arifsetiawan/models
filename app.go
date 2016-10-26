package models

// ApplicationJSON is a struct representation for Application Object
type ApplicationJSON struct {
	Type           string   `json:"_type,omitempty"`
	ID             string   `json:"id"`
	AppName        string   `json:"app_name,omitempty"`
	AppDescription string   `json:"app_description,omitempty"`
	AppURL         string   `json:"app_url"`
	AppIcon        string   `json:"app_icon,omitempty"`
	Members        []string `json:"members"`
}
