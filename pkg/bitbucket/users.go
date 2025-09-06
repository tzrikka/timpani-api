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
func UsersGetActivity(ctx workflow.Context, accountID, uuid string) (map[string]any, error) {
	req := UsersGetRequest{AccountID: accountID, UUID: uuid}
	fut := internal.ExecuteTimpaniActivity(ctx, UsersGetActivityName, req)

	resp := &map[string]any{}
	if err := fut.Get(ctx, resp); err != nil {
		return *resp, err
	}

	return *resp, nil
}
