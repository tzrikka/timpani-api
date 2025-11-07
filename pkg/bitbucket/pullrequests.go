package bitbucket

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	PullRequestsApproveActivityName         = "bitbucket.pullrequests.approve"
	PullRequestsCreateCommentActivityName   = "bitbucket.pullrequests.createComment"
	PullRequestsDeclineActivityName         = "bitbucket.pullrequests.decline"
	PullRequestsDeleteCommentActivityName   = "bitbucket.pullrequests.deleteComment"
	PullRequestsListActivityLogActivityName = "bitbucket.pullrequests.listActivityLog"
	PullRequestsListCommitsActivityName     = "bitbucket.pullrequests.listCommits"
	PullRequestsListForCommitActivityName   = "bitbucket.pullrequests.listForCommit"
	PullRequestsMergeActivityName           = "bitbucket.pullrequests.merge"
	PullRequestsUnapproveActivityName       = "bitbucket.pullrequests.unapprove"
	PullRequestsUpdateCommentActivityName   = "bitbucket.pullrequests.updateComment"
)

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-approve-post
type PullRequestsApproveRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-approve-post
func PullRequestsApproveActivity(ctx workflow.Context, req PullRequestsApproveRequest) error {
	return internal.ExecuteTimpaniActivity(ctx, PullRequestsApproveActivityName, req).Get(ctx, nil)
}

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

	// Content          `json:"content"`     // Unnecessary.
	// PullRequest      `json:"pullrequest"` // Unnecessary.
	// CreatedOn string `json:"created_on"`  // Unnecessary.
	// UpdatedOn string `json:"updated_on"`  // Unnecessary.
	// Deleted   bool   `json:"deleted"`     // Always false.
	// Pending   bool   `json:"pending"`     // Unnecessary.

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

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-decline-post
type PullRequestsDeclineRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-decline-post
func PullRequestsDeclineActivity(ctx workflow.Context, req PullRequestsDeclineRequest) error {
	return internal.ExecuteTimpaniActivity(ctx, PullRequestsDeclineActivityName, req).Get(ctx, nil)
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

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-activity-get
type PullRequestsListActivityLogRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen  string `json:"pagelen,omitempty"`
	Page     string `json:"page,omitempty"`
	AllPages bool   `json:"all_pages,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-activity-get
type PullRequestsListActivityLogResponse struct {
	Values []map[string]any `json:"values"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen int    `json:"pagelen,omitempty"`
	Next    string `json:"next,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-activity-get
func PullRequestsListActivityLogActivity(ctx workflow.Context, req PullRequestsListActivityLogRequest) (*PullRequestsListActivityLogResponse, error) {
	fut := internal.ExecuteTimpaniActivity(ctx, PullRequestsListActivityLogActivityName, req)

	resp := new(PullRequestsListActivityLogResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-commits-get
type PullRequestsListCommitsRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen  string `json:"pagelen,omitempty"`
	Page     string `json:"page,omitempty"`
	AllPages bool   `json:"all_pages,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-commits-get
type PullRequestsListCommitsResponse struct {
	Values []Commit `json:"values"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen int    `json:"pagelen,omitempty"`
	Page    int    `json:"page,omitempty"`
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

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-commit-commit-pullrequests-get
type PullRequestsListForCommitRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace string `json:"workspace"`
	RepoSlug  string `json:"repo_slug"`
	Commit    string `json:"commit"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen  string `json:"pagelen,omitempty"`
	Page     string `json:"page,omitempty"`
	AllPages bool   `json:"all_pages,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-commit-commit-pullrequests-get
type PullRequestsListForCommitResponse struct {
	Values []map[string]any `json:"values"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen int    `json:"pagelen,omitempty"`
	Page    int    `json:"page,omitempty"`
	Next    string `json:"next,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-commit-commit-pullrequests-get
func PullRequestsListForCommitActivity(ctx workflow.Context, req PullRequestsListForCommitRequest) (*PullRequestsListForCommitResponse, error) {
	fut := internal.ExecuteTimpaniActivity(ctx, PullRequestsListForCommitActivityName, req)

	resp := new(PullRequestsListForCommitResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-merge-post
type PullRequestsMergeRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`

	Type              string `json:"type,omitempty"`
	Message           string `json:"message,omitempty"`
	MergeStrategy     string `json:"merge_strategy,omitempty"`
	CloseSourceBranch bool   `json:"close_source_branch,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-merge-post
func PullRequestsMergeActivity(ctx workflow.Context, req PullRequestsMergeRequest) error {
	return internal.ExecuteTimpaniActivity(ctx, PullRequestsMergeActivityName, req).Get(ctx, nil)
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-approve-delete
type PullRequestsUnapproveRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-approve-delete
func PullRequestsUnapproveActivity(ctx workflow.Context, req PullRequestsUnapproveRequest) error {
	return internal.ExecuteTimpaniActivity(ctx, PullRequestsUnapproveActivityName, req).Get(ctx, nil)
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

	// Repository *Repository `json:"repository,omitempty"` // Unnecessary.

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
