package slack

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	UsersConversationsActivityName = "slack.users.conversations"
	UsersGetPresenceActivityName   = "slack.users.getPresence"
	UsersIdentityActivityName      = "slack.users.identity"
	UsersInfoActivityName          = "slack.users.info"
	UsersListActivityName          = "slack.users.list"
	UsersLookupByEmailActivityName = "slack.users.lookupByEmail"
	UsersProfileGetActivityName    = "slack.users.profile.get"
)

// https://docs.slack.dev/reference/methods/users.conversations/
type UsersConversationsRequest struct {
	Types           string `json:"types,omitempty"`
	User            string `json:"user,omitempty"`
	ExcludeArchived bool   `json:"exclude_archived,omitempty"`

	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`

	TeamID string `json:"team_id,omitempty"`
}

// https://docs.slack.dev/reference/methods/users.conversations/
type UsersConversationsResponse struct {
	Response

	Channels []map[string]any `json:"channels,omitempty"`
}

// https://docs.slack.dev/reference/methods/users.getPresence/
type UsersGetPresenceRequest struct {
	User string `json:"user,omitempty"`
}

// https://docs.slack.dev/reference/methods/users.getPresence/
type UsersGetPresenceResponse struct {
	Response

	Presence string `json:"presence,omitempty"`
}

// https://docs.slack.dev/reference/methods/users.info/
type UsersInfoRequest struct {
	User string `json:"user"`

	IncludeLocale bool `json:"include_locale,omitempty"`
}

// https://docs.slack.dev/reference/methods/users.info/
type UsersInfoResponse struct {
	Response

	User *User `json:"user,omitempty"`
}

func UsersInfoActivity(ctx workflow.Context, userID string) (*User, error) {
	req := UsersInfoRequest{User: userID}
	fut := internal.ExecuteTimpaniActivity(ctx, UsersInfoActivityName, req)

	resp := new(UsersInfoResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp.User, nil
}

// https://docs.slack.dev/reference/methods/users.list/
type UsersListRequest struct {
	IncludeLocale bool `json:"include_locale,omitempty"`

	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`

	TeamID string `json:"team_id,omitempty"`
}

// https://docs.slack.dev/reference/methods/users.list/
type UsersListResponse struct {
	Response

	Members []map[string]any `json:"members,omitempty"`
	CacheTS int64            `json:"cache_ts,omitempty"`
}

// https://docs.slack.dev/reference/methods/users.lookupByEmail/
type UsersLookupByEmailRequest struct {
	Email string `json:"email"`
}

// https://docs.slack.dev/reference/methods/users.lookupByEmail/
type UsersLookupByEmailResponse struct {
	Response

	User *User `json:"user,omitempty"`
}

func UsersLookupByEmailActivity(ctx workflow.Context, email string) (*User, error) {
	req := UsersLookupByEmailRequest{Email: email}
	fut := internal.ExecuteTimpaniActivity(ctx, UsersLookupByEmailActivityName, req)

	resp := new(UsersLookupByEmailResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp.User, nil
}

// https://docs.slack.dev/reference/methods/users.profile.get/
type UsersProfileGetRequest struct {
	User string `json:"user"`

	IncludeLabels bool `json:"include_labels,omitempty"`
}

// https://docs.slack.dev/reference/methods/users.profile.get/
type UsersProfileGetResponse struct {
	Response

	Profile *Profile `json:"profile,omitempty"`
}

func UsersProfileGetActivity(ctx workflow.Context, userID string) (*Profile, error) {
	req := UsersProfileGetRequest{User: userID}
	fut := internal.ExecuteTimpaniActivity(ctx, UsersProfileGetActivityName, req)

	resp := new(UsersProfileGetResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp.Profile, nil
}

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
