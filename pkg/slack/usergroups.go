package slack

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

//revive:disable:exported
const (
	UserGroupsListActivityName      = "slack.usergroups.list"
	UserGroupsUsersListActivityName = "slack.usergroups.users.list"
) //revive:enable:exported

// UserGroupsListRequest is based on:
// https://docs.slack.dev/reference/methods/usergroups.list/
type UserGroupsListRequest struct {
	IncludeCount    bool `json:"include_count,omitempty"`
	IncludeDisabled bool `json:"include_disabled,omitempty"`
	IncludeUsers    bool `json:"include_users,omitempty"`

	TeamID string `json:"team_id,omitempty"`
}

// UserGroupsListResponse is based on:
// https://docs.slack.dev/reference/methods/usergroups.list/
type UserGroupsListResponse struct {
	Response

	Usergroups []UserGroup `json:"usergroups,omitempty"`
}

// UserGroupsList is based on:
// https://docs.slack.dev/reference/methods/usergroups.list/
func UserGroupsList(ctx workflow.Context, count, disabled, users bool) ([]UserGroup, error) {
	req := UserGroupsListRequest{IncludeCount: count, IncludeDisabled: disabled, IncludeUsers: users}
	resp, err := internal.ExecuteTimpaniActivity[UserGroupsListResponse](ctx, UserGroupsListActivityName, req)
	if err != nil {
		return nil, err
	}
	return resp.Usergroups, nil
}

// UserGroupsUsersListRequest is based on:
// https://docs.slack.dev/reference/methods/usergroups.users.list/
type UserGroupsUsersListRequest struct {
	Usergroup string `json:"usergroup"`

	IncludeDisabled bool   `json:"include_disabled,omitempty"`
	TeamID          string `json:"team_id,omitempty"`
}

// UserGroupsUsersListResponse is based on:
// https://docs.slack.dev/reference/methods/usergroups.users.list/
type UserGroupsUsersListResponse struct {
	Response

	Users []string `json:"users,omitempty"`
}

// UserGroupsUsersList is based on:
// https://docs.slack.dev/reference/methods/usergroups.users.list/
func UserGroupsUsersList(ctx workflow.Context, usergroup string, includeDisabled bool) ([]string, error) {
	req := UserGroupsUsersListRequest{Usergroup: usergroup, IncludeDisabled: includeDisabled}
	resp, err := internal.ExecuteTimpaniActivity[UserGroupsUsersListResponse](ctx, UserGroupsUsersListActivityName, req)
	if err != nil {
		return nil, err
	}
	return resp.Users, nil
}

// UserGroup is based on:
// https://docs.slack.dev/reference/objects/usergroup-object/
type UserGroup struct {
	ID                  string `json:"id,omitempty"`
	TeamID              string `json:"team_id,omitempty"`
	EnterpriseSubteamID string `json:"enterprise_subteam_id,omitempty"`

	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Handle      string `json:"handle,omitempty"`
	AutoType    string `json:"auto_type,omitempty"`

	DateCreate int64 `json:"date_create,omitempty"`
	DateUpdate int64 `json:"date_update,omitempty"`
	DateDelete int64 `json:"date_delete,omitempty"`

	// CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy string `json:"updated_by,omitempty"`
	// DeletedBy string `json:"deleted_by,omitempty"`.

	// IsUsergroup         bool `json:"is_usergroup,omitempty"`
	// IsSubteam           bool `json:"is_subteam,omitempty"`
	// IsExternal          bool `json:"is_external,omitempty"`
	// IsIDPGroup          bool `json:"is_idp_group,omitempty"`
	// IsOrgLevel          bool `json:"is_org_level,omitempty"`
	// IsEditingRestricted bool `json:"is_editing_restricted,omitempty"`
	// IsMembershipLocked  bool `json:"is_membership_locked,omitempty"`
	// IsSection           bool `json:"is_section,omitempty"`
	// IsVisible           bool `json:"is_visible,omitempty"`.

	Prefs Prefs    `json:"prefs"`
	Users []string `json:"users,omitempty"`

	UserCount    int `json:"user_count,omitempty"`
	ChannelCount int `json:"channel_count,omitempty"`
}

// Prefs is based on:
// https://docs.slack.dev/reference/objects/usergroup-object/
type Prefs struct {
	Channels []string `json:"channels,omitempty"`
	Groups   []string `json:"groups,omitempty"`
}
