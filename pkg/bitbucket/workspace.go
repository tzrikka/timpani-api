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
type WorkspacesListMembersResponse struct {
	Size     int          `json:"size,omitempty"`
	Page     int          `json:"page,omitempty"`
	Pagelen  int          `json:"pagelen,omitempty"`
	Next     string       `json:"next,omitempty"`
	Previous string       `json:"previous,omitempty"`
	Values   []Membership `json:"values,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-workspace-members-get
func WorkspacesListMembersActivity(ctx workflow.Context, workspace string, emailsFilter []string) ([]User, error) {
	req := WorkspacesListMembersRequest{Workspace: workspace, EmailsFilter: emailsFilter}
	fut := internal.ExecuteTimpaniActivity(ctx, WorkspacesListMembersActivityName, req)

	resp := new(WorkspacesListMembersResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	users := make([]User, len(resp.Values))
	for i, membership := range resp.Values {
		users[i] = membership.User
	}

	return users, nil
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-workspaces/#api-workspaces-workspace-members-get
type Membership struct {
	// Type  string `json:"type"`           // Always "workspace_membership".
	// Links map[string]Link `json:"links"` // Inconsequential.

	User User `json:"user"`

	// Workspace // Inconsequential.
}
