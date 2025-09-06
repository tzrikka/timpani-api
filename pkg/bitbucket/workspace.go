package bitbucket

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	WorkspacesListMembersActivityName = "bitbucket.workspaces.listMembers"
)

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-workspace-members-get
type WorkspacesListMembersRequest struct {
	Workspace    string   `json:"workspace"`
	EmailsFilter []string `json:"emails_filter"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-workspace-members-get
func WorkspacesListMembersActivity(ctx workflow.Context, workspace string, emailsFilter []string) (map[string]any, error) {
	req := WorkspacesListMembersRequest{Workspace: workspace, EmailsFilter: emailsFilter}
	fut := internal.ExecuteTimpaniActivity(ctx, WorkspacesListMembersActivityName, req)

	resp := &map[string]any{}
	if err := fut.Get(ctx, resp); err != nil {
		return *resp, err
	}

	return *resp, nil
}
