package bitbucket

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

//revive:disable:exported
const (
	PullRequestsApproveActivityName         = "bitbucket.pullrequests.approve"
	PullRequestsCreateCommentActivityName   = "bitbucket.pullrequests.createComment"
	PullRequestsDeclineActivityName         = "bitbucket.pullrequests.decline"
	PullRequestsDeleteCommentActivityName   = "bitbucket.pullrequests.deleteComment"
	PullRequestsDiffstatActivityName        = "bitbucket.pullrequests.diffstat"
	PullRequestsListActivityLogActivityName = "bitbucket.pullrequests.listActivityLog"
	PullRequestsListCommitsActivityName     = "bitbucket.pullrequests.listCommits"
	PullRequestsListForCommitActivityName   = "bitbucket.pullrequests.listForCommit"
	PullRequestsMergeActivityName           = "bitbucket.pullrequests.merge"
	PullRequestsUnapproveActivityName       = "bitbucket.pullrequests.unapprove"
	PullRequestsUpdateCommentActivityName   = "bitbucket.pullrequests.updateComment"
) //revive:enable:exported

// PullRequestsApproveRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-approve-post
type PullRequestsApproveRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
}

// PullRequestsApprove is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-approve-post
func PullRequestsApprove(ctx workflow.Context, req PullRequestsApproveRequest) error {
	return internal.ExecuteTimpaniActivityNoResp(ctx, PullRequestsApproveActivityName, req)
}

// PullRequestsCreateCommentRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
type PullRequestsCreateCommentRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
	Markdown      string `json:"text"`

	ParentID string `json:"parent_id,omitempty"`
}

// PullRequestsCreateCommentResponse is based on:
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

// PullRequestsCreateComment is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
func PullRequestsCreateComment(ctx workflow.Context, req PullRequestsCreateCommentRequest) (*PullRequestsCreateCommentResponse, error) {
	return internal.ExecuteTimpaniActivity[PullRequestsCreateCommentResponse](ctx, PullRequestsCreateCommentActivityName, req)
}

// PullRequestsDeclineRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-decline-post
type PullRequestsDeclineRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
}

// PullRequestsDecline is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-decline-post
func PullRequestsDecline(ctx workflow.Context, req PullRequestsDeclineRequest) error {
	return internal.ExecuteTimpaniActivityNoResp(ctx, PullRequestsDeclineActivityName, req)
}

// PullRequestsDeleteCommentRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-delete
type PullRequestsDeleteCommentRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
	CommentID     string `json:"comment_id"`
}

// PullRequestsDeleteComment is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-delete
func PullRequestsDeleteComment(ctx workflow.Context, req PullRequestsDeleteCommentRequest) error {
	return internal.ExecuteTimpaniActivityNoResp(ctx, PullRequestsDeleteCommentActivityName, req)
}

// PullRequestsDiffstatRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-diffstat-get
type PullRequestsDiffstatRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen string `json:"pagelen,omitempty"`
	Page    string `json:"page,omitempty"`

	Next string `json:"next,omitempty"`
}

// PullRequestsDiffstatResponse is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-diffstat-get
type PullRequestsDiffstatResponse = CommitsDiffstatResponse

// PullRequestsDiffstat is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-diffstat-get
func PullRequestsDiffstat(ctx workflow.Context, req PullRequestsDiffstatRequest) ([]Diffstat, error) {
	var ds []Diffstat
	req.Next = "start"

	for req.Next != "" {
		if req.Next == "start" {
			req.Next = ""
		}

		resp, err := internal.ExecuteTimpaniActivity[PullRequestsDiffstatResponse](ctx, PullRequestsDiffstatActivityName, req)
		if err != nil {
			return nil, err
		}

		ds = append(ds, resp.Values...)
		req.Next = resp.Next
	}

	return ds, nil
}

// PullRequestsListActivityLogRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-activity-get
type PullRequestsListActivityLogRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen string `json:"pagelen,omitempty"`
	Page    string `json:"page,omitempty"`

	Next string `json:"next,omitempty"`
}

// PullRequestsListActivityLogResponse is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-activity-get
type PullRequestsListActivityLogResponse struct {
	Values []map[string]any `json:"values"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen int    `json:"pagelen,omitempty"`
	Next    string `json:"next,omitempty"`
}

// PullRequestsListActivityLog is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-activity-get
func PullRequestsListActivityLog(ctx workflow.Context, req PullRequestsListActivityLogRequest) ([]map[string]any, error) {
	var activities []map[string]any
	req.Next = "start"

	for req.Next != "" {
		if req.Next == "start" {
			req.Next = ""
		}

		resp, err := internal.ExecuteTimpaniActivity[PullRequestsListActivityLogResponse](ctx, PullRequestsListActivityLogActivityName, req)
		if err != nil {
			return nil, err
		}

		activities = append(activities, resp.Values...)
		req.Next = resp.Next
	}

	return activities, nil
}

// PullRequestsListCommitsRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-commits-get
type PullRequestsListCommitsRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen string `json:"pagelen,omitempty"`
	Page    string `json:"page,omitempty"`

	Next string `json:"next,omitempty"`
}

// PullRequestsListCommitsResponse is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-commits-get
type PullRequestsListCommitsResponse struct {
	Values []Commit `json:"values"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen int    `json:"pagelen,omitempty"`
	Page    int    `json:"page,omitempty"`
	Next    string `json:"next,omitempty"`
}

// PullRequestsListCommits is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-commits-get
func PullRequestsListCommits(ctx workflow.Context, req PullRequestsListCommitsRequest) ([]Commit, error) {
	var cs []Commit
	req.Next = "start"

	for req.Next != "" {
		if req.Next == "start" {
			req.Next = ""
		}

		resp, err := internal.ExecuteTimpaniActivity[PullRequestsListCommitsResponse](ctx, PullRequestsListCommitsActivityName, req)
		if err != nil {
			return nil, err
		}

		cs = append(cs, resp.Values...)
		req.Next = resp.Next
	}

	return cs, nil
}

// PullRequestsListForCommitRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-commit-commit-pullrequests-get
type PullRequestsListForCommitRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace string `json:"workspace"`
	RepoSlug  string `json:"repo_slug"`
	Commit    string `json:"commit"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen string `json:"pagelen,omitempty"`
	Page    string `json:"page,omitempty"`

	Next string `json:"next,omitempty"`
}

// PullRequestsListForCommitResponse is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-commit-commit-pullrequests-get
type PullRequestsListForCommitResponse struct {
	Values []map[string]any `json:"values"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen int    `json:"pagelen,omitempty"`
	Page    int    `json:"page,omitempty"`
	Next    string `json:"next,omitempty"`
}

// PullRequestsListForCommit is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-commit-commit-pullrequests-get
func PullRequestsListForCommit(ctx workflow.Context, req PullRequestsListForCommitRequest) ([]map[string]any, error) {
	var prs []map[string]any
	req.Next = "start"

	for req.Next != "" {
		if req.Next == "start" {
			req.Next = ""
		}

		resp, err := internal.ExecuteTimpaniActivity[PullRequestsListForCommitResponse](ctx, PullRequestsListForCommitActivityName, req)
		if err != nil {
			return nil, err
		}

		prs = append(prs, resp.Values...)
		req.Next = resp.Next
	}

	return prs, nil
}

// PullRequestsMergeRequest is based on:
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

// PullRequestsMerge is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-merge-post
func PullRequestsMerge(ctx workflow.Context, req PullRequestsMergeRequest) error {
	return internal.ExecuteTimpaniActivityNoResp(ctx, PullRequestsMergeActivityName, req)
}

// PullRequestsUnapproveRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-approve-delete
type PullRequestsUnapproveRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
}

// PullRequestsUnapprove is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-approve-delete
func PullRequestsUnapprove(ctx workflow.Context, req PullRequestsUnapproveRequest) error {
	return internal.ExecuteTimpaniActivityNoResp(ctx, PullRequestsUnapproveActivityName, req)
}

// PullRequestsUpdateCommentRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-put
type PullRequestsUpdateCommentRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
	CommentID     string `json:"comment_id"`
	Markdown      string `json:"text"`
}

// PullRequestsUpdateComment is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-put
func PullRequestsUpdateComment(ctx workflow.Context, req PullRequestsUpdateCommentRequest) error {
	return internal.ExecuteTimpaniActivityNoResp(ctx, PullRequestsUpdateCommentActivityName, req)
}

// Commit is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-commit-commit-get
type Commit struct {
	// Type string `json:"type"` // Always "commit".

	Hash    string   `json:"hash"`
	Date    string   `json:"date,omitempty"`
	Author  *User    `json:"author,omitempty"`
	Message string   `json:"message,omitempty"`
	Parents []Commit `json:"parents,omitempty"`

	// Repository *Repository `json:"repository,omitempty"` // Unnecessary.
	// Summary *Rendered `json:"summary,omitempty"`         // Unnecessary.

	Links map[string]Link `json:"links"`
}

//revive:disable:exported
type Link struct {
	HRef string `json:"href"`
} //revive:enable:exported

// Parent is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
type Parent struct {
	ID    int             `json:"id"`
	Links map[string]Link `json:"links"`
}
