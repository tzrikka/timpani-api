package jira

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	UsersGetActivityName = "jira.users.get"
)

// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-users/#api-rest-api-3-user-get
type UsersGetRequest struct {
	AccountID string `json:"account_id"`
}

// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-users/#api-rest-api-3-user-get
type UsersGetResponse User

// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-users/#api-rest-api-3-user-get
func UsersGetActivity(ctx workflow.Context, accountID string) (*UsersGetResponse, error) {
	req := UsersGetRequest{AccountID: accountID}
	fut := internal.ExecuteTimpaniActivity(ctx, UsersGetActivityName, req)

	resp := new(UsersGetResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-users/#api-rest-api-3-user-get
type User struct {
	Self        string `json:"self"`
	AccountID   string `json:"accountId"`
	AccountType string `json:"accountType"`
	Active      bool   `json:"active"`

	DisplayName string `json:"displayName"`
	Email       string `json:"emailAddress"`
	TimeZone    string `json:"timeZone"`
	Locale      string `json:"locale"`

	AvatarURLs map[string]string `json:"avatarUrls"`
}
