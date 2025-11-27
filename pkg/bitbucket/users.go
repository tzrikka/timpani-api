package bitbucket

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	UsersGetActivityName = "bitbucket.users.get"
)

// UsersGetRequest is based on:
//   - https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-user-get
//   - https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-users-selected-user-get
type UsersGetRequest struct {
	AccountID string `json:"account_id,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}

// UsersGetResponse is based on:
//   - https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-user-get
//   - https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-users-selected-user-get
type UsersGetResponse = User

// UsersGet is based on:
//   - https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-user-get
//   - https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-users-selected-user-get
func UsersGet(ctx workflow.Context, accountID, uuid string) (*User, error) {
	req := UsersGetRequest{AccountID: accountID, UUID: uuid}
	return internal.ExecuteTimpaniActivity[UsersGetResponse](ctx, UsersGetActivityName, req)
}

// User is based on:
//   - https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-user-get
//   - https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-users-selected-user-get
//   - https://developer.atlassian.com/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-workspace-members-get
type User struct {
	Type string `json:"type"`

	DisplayName string `json:"display_name"`
	Nickname    string `json:"nickname"`
	AccountID   string `json:"account_id"`
	UUID        string `json:"uuid"`

	Links map[string]Link `json:"links"`
}
