package slack

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	UsersConversationsActivityName = "slack.users.conversations"
	UsersGetPresenceActivityName   = "slack.users.getPresence"
	UsersInfoActivityName          = "slack.users.info"
	UsersListActivityName          = "slack.users.list"
	UsersLookupByEmailActivityName = "slack.users.lookupByEmail"
	UsersProfileGetActivityName    = "slack.users.profile.get"
)

// UsersConversationsRequest is based on:
// https://docs.slack.dev/reference/methods/users.conversations/
type UsersConversationsRequest struct {
	Types           string `json:"types,omitempty"`
	User            string `json:"user,omitempty"`
	ExcludeArchived bool   `json:"exclude_archived,omitempty"`

	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`

	TeamID string `json:"team_id,omitempty"`
}

// UsersConversationsResponse is based on:
// https://docs.slack.dev/reference/methods/users.conversations/
type UsersConversationsResponse struct {
	Response

	Channels []map[string]any `json:"channels,omitempty"`
}

// UsersGetPresenceRequest is based on:
// https://docs.slack.dev/reference/methods/users.getPresence/
type UsersGetPresenceRequest struct {
	User string `json:"user,omitempty"`
}

// UsersGetPresenceResponse is based on:
// https://docs.slack.dev/reference/methods/users.getPresence/
type UsersGetPresenceResponse struct {
	Response

	Presence string `json:"presence,omitempty"`
}

// UsersInfoRequest is based on:
// https://docs.slack.dev/reference/methods/users.info/
type UsersInfoRequest struct {
	User string `json:"user"`

	IncludeLocale bool `json:"include_locale,omitempty"`
}

// UsersInfoResponse is based on:
// https://docs.slack.dev/reference/methods/users.info/
type UsersInfoResponse struct {
	Response

	User *User `json:"user,omitempty"`
}

// UsersInfo is based on:
// https://docs.slack.dev/reference/methods/users.info/
func UsersInfo(ctx workflow.Context, userID string) (*User, error) {
	req := UsersInfoRequest{User: userID}
	resp, err := internal.ExecuteTimpaniActivity[UsersInfoResponse](ctx, UsersInfoActivityName, req)
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}

// UsersListRequest is based on:
// https://docs.slack.dev/reference/methods/users.list/
type UsersListRequest struct {
	IncludeLocale bool `json:"include_locale,omitempty"`

	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`

	TeamID string `json:"team_id,omitempty"`
}

// UsersListResponse is based on:
// https://docs.slack.dev/reference/methods/users.list/
type UsersListResponse struct {
	Response

	Members []map[string]any `json:"members,omitempty"`
	CacheTS int64            `json:"cache_ts,omitempty"`
}

// UsersLookupByEmailRequest is based on:
// https://docs.slack.dev/reference/methods/users.lookupByEmail/
type UsersLookupByEmailRequest struct {
	Email string `json:"email"`
}

// UsersLookupByEmailResponse is based on:
// https://docs.slack.dev/reference/methods/users.lookupByEmail/
type UsersLookupByEmailResponse struct {
	Response

	User *User `json:"user,omitempty"`
}

// UsersLookupByEmail is based on:
// https://docs.slack.dev/reference/methods/users.lookupByEmail/
func UsersLookupByEmail(ctx workflow.Context, email string) (*User, error) {
	req := UsersLookupByEmailRequest{Email: email}
	resp, err := internal.ExecuteTimpaniActivity[UsersLookupByEmailResponse](ctx, UsersLookupByEmailActivityName, req)
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}

// UsersProfileGetRequest is based on:
// https://docs.slack.dev/reference/methods/users.profile.get/
type UsersProfileGetRequest struct {
	User string `json:"user"`

	IncludeLabels bool `json:"include_labels,omitempty"`
}

// UsersProfileGetResponse is based on:
// https://docs.slack.dev/reference/methods/users.profile.get/
type UsersProfileGetResponse struct {
	Response

	Profile *Profile `json:"profile,omitempty"`
}

// UsersProfileGet is based on:
// https://docs.slack.dev/reference/methods/users.profile.get/
func UsersProfileGet(ctx workflow.Context, userID string) (*Profile, error) {
	req := UsersProfileGetRequest{User: userID}
	resp, err := internal.ExecuteTimpaniActivity[UsersProfileGetResponse](ctx, UsersProfileGetActivityName, req)
	if err != nil {
		return nil, err
	}
	return resp.Profile, nil
}

// User is based on:
// https://docs.slack.dev/reference/objects/user-object/
type User struct {
	ID       string `json:"id"`
	TeamID   string `json:"team_id"`
	RealName string `json:"real_name"`
	IsBot    bool   `json:"is_bot"`

	TZ       string `json:"tz"`
	TZLabel  string `json:"tz_label"`
	TZOffset int    `json:"tz_offset"`

	Updated int64 `json:"updated"`

	Profile Profile `json:"profile"`
}

// Profile is based on:
// https://docs.slack.dev/reference/objects/user-object/#profile
type Profile struct {
	DisplayName           string `json:"display_name"`
	DisplayNameNormalized string `json:"display_name_normalized"`

	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	RealName           string `json:"real_name"`
	RealNameNormalized string `json:"real_name_normalized"`

	Email string `json:"email"`
	Team  string `json:"team"`

	Image24  string `json:"image_24"`
	Image32  string `json:"image_32"`
	Image48  string `json:"image_48"`
	Image72  string `json:"image_72"`
	Image192 string `json:"image_192"`
	Image512 string `json:"image_512"`

	APIAppID     string `json:"api_app_id,omitempty"`
	BotID        string `json:"bot_id,omitempty"`
	AlwaysActive bool   `json:"always_active,omitempty"`

	// https://docs.slack.dev/reference/methods/users.profile.set#custom_profile
}
