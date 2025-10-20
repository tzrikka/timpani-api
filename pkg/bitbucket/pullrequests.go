package bitbucket

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	PullRequestsCreateCommentActivityName = "bitbucket.pullrequests.createComment"
	PullRequestsDeleteCommentActivityName = "bitbucket.pullrequests.deleteComment"
	PullRequestsListCommitsActivityName   = "bitbucket.pullrequests.listCommits"
	PullRequestsUpdateCommentActivityName = "bitbucket.pullrequests.updateComment"
)

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
type PullRequestsCreateCommentRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
	Markdown      string `json:"text"`

	ParentID string `json:"parent_id,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
type PullRequestsCreateCommentResponse struct {
	// Type string `json:"type"` // Always "pullrequest_comment".

	ID int `json:"id"`

	// Content          `json:"content"`     // Inconsequential.
	// PullRequest      `json:"pullrequest"` // Inconsequential.
	// CreatedOn string `json:"created_on"`  // Inconsequential.
	// UpdatedOn string `json:"updated_on"`  // Inconsequential.
	// Deleted   bool   `json:"deleted"`     // Always false.
	// Pending   bool   `json:"pending"`     // Inconsequential.

	User   User    `json:"user"`
	Parent *Parent `json:"parent,omitempty"`

	Links map[string]Link `json:"links"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
func PullRequestsCreateCommentActivity(ctx workflow.Context, req PullRequestsCreateCommentRequest) (*PullRequestsCreateCommentResponse, error) {
	fut := internal.ExecuteTimpaniActivity(ctx, PullRequestsCreateCommentActivityName, req)

	resp := new(PullRequestsCreateCommentResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-delete
type PullRequestsDeleteCommentRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
	CommentID     string `json:"comment_id"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-delete
func PullRequestsDeleteCommentActivity(ctx workflow.Context, req PullRequestsDeleteCommentRequest) error {
	return internal.ExecuteTimpaniActivity(ctx, PullRequestsDeleteCommentActivityName, req).Get(ctx, nil)
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-commits-get
type PullRequestsListCommitsRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	Page int `json:"page,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-commits-get
type PullRequestsListCommitsResponse struct {
	Values []Commit `json:"values"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	Pagelen int    `json:"pagelen"`
	Next    string `json:"next,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-commits-get
func PullRequestsListCommitsActivity(ctx workflow.Context, req PullRequestsListCommitsRequest) (*PullRequestsListCommitsResponse, error) {
	fut := internal.ExecuteTimpaniActivity(ctx, PullRequestsListCommitsActivityName, req)

	resp := new(PullRequestsListCommitsResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-put
type PullRequestsUpdateCommentRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
	CommentID     string `json:"comment_id"`
	Markdown      string `json:"text"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-put
func PullRequestsUpdateCommentActivity(ctx workflow.Context, req PullRequestsUpdateCommentRequest) error {
	return internal.ExecuteTimpaniActivity(ctx, PullRequestsUpdateCommentActivityName, req).Get(ctx, nil)
}

type Commit struct {
	// Type string `json:"type"` // Always "commit".

	Hash    string    `json:"hash"`
	Date    string    `json:"date,omitempty"`
	Author  *User     `json:"author,omitempty"`
	Message string    `json:"message,omitempty"`
	Summary *Rendered `json:"summary,omitempty"`
	Parents []Commit  `json:"parents,omitempty"`

	// Repository *Repository `json:"repository,omitempty"` // Inconsequential.

	Links map[string]Link `json:"links"`
}

type Link struct {
	HRef string `json:"href"`
}

type Parent struct {
	ID    int             `json:"id"`
	Links map[string]Link `json:"links"`
}

type Rendered struct {
	// Type string `json:"type"` // Always "rendered".

	Raw    string `json:"raw"`
	Markup string `json:"markup"`
	HTML   string `json:"html"`
}
