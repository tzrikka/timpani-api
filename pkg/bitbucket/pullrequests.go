package bitbucket

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	PullRequestsCreateCommentActivityName = "bitbucket.pullrequests.createComment"
)

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
type PullRequestsCreateCommentRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
	Markdown      string `json:"text"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
type PullRequestsCreateCommentResponse = map[string]any

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
func PullRequestsCreateCommentActivity(ctx workflow.Context, workspace, repo, prID, markdown, linkID string) (*PullRequestsCreateCommentResponse, error) {
	req := PullRequestsCreateCommentRequest{ThrippyLinkID: linkID, Workspace: workspace, RepoSlug: repo, PullRequestID: prID, Markdown: markdown}
	fut := internal.ExecuteTimpaniActivity(ctx, PullRequestsCreateCommentActivityName, req)

	resp := new(PullRequestsCreateCommentResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
