package bitbucket

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	CommitsDiffActivityName = "bitbucket.commits.diff"
)

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diff-spec-get
type CommitsDiffRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace string `json:"workspace"`
	RepoSlug  string `json:"repo_slug"`
	Spec      string `json:"spec"`
}

func CommitsDiffActivity(ctx workflow.Context, req CommitsDiffRequest) (string, error) {
	fut := internal.ExecuteTimpaniActivity(ctx, CommitsDiffActivityName, req)

	resp := new(string)
	if err := fut.Get(ctx, resp); err != nil {
		return "", err
	}

	return *resp, nil
}
