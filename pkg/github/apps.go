package github

import "time"

// App is based on:
// https://docs.github.com/en/rest/apps/apps?apiVersion=2022-11-28
type App struct {
	ID          int64  `json:"id"`
	NodeID      string `json:"node_id"`
	HTMLURL     string `json:"html_url"`
	ExternalURL string `json:"external_url"`

	// ClientID string `json:"client_id"` // Unnecessary.

	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	Owner       User              `json:"owner"`
	Permissions map[string]string `json:"permissions,omitempty"`
	Events      []string          `json:"events,omitempty"`

	// InstallationsCount int `json:"installations_count,omitempty"` // Unnecessary.

	CreatedAt time.Time `json:"created_at,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
}
