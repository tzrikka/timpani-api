package bitbucket

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	UsersGetActivityName = "bitbucket.users.get"
)

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-user-get
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-users-selected-user-get
type UsersGetRequest struct {
	AccountID string `json:"account_id,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-user-get
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-users-selected-user-get
type UsersGetResponse User

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-user-get
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-users-selected-user-get
func UsersGetActivity(ctx workflow.Context, accountID, uuid string) (*UsersGetResponse, error) {
	req := UsersGetRequest{AccountID: accountID, UUID: uuid}
	fut := internal.ExecuteTimpaniActivity(ctx, UsersGetActivityName, req)

	resp := new(UsersGetResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-user-get
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-users/#api-users-selected-user-get
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-workspace-members-get
type User struct {
	Type string `json:"type"`

	DisplayName string `json:"display_name"`
	Nickname    string `json:"nickname"`
	AccountID   string `json:"account_id"`
	UUID        string `json:"uuid"`

	Links map[string]Link `json:"links"`
}

type Link struct {
	HRef string `json:"href"`
}
