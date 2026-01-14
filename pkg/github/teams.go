package github

// Team is based on:
//   - https://docs.github.com/en/rest/teams/teams?apiVersion=2022-11-28
//   - https://docs.github.com/en/webhooks/webhook-events-and-payloads#team
type Team struct {
	ID      int64  `json:"id"`
	NodeID  string `json:"node_id"`
	HTMLURL string `json:"html_url"`

	Slug        string  `json:"slug"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`

	Privacy             string `json:"privacy"`              // "open", "closed", "secret".
	NotificationSetting string `json:"notification_setting"` // "notifications_enabled", "notifications_disabled".
	Permission          string `json:"permission,omitempty"` // "pull", "triage", "push", "maintain", "admin".

	Parent *Team `json:"parent,omitempty"`
}
