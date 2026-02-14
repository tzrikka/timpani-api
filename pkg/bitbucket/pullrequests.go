package bitbucket

import (
	"time"

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
	PullRequestsGetActivityName             = "bitbucket.pullrequests.get"
	PullRequestsGetCommentActivityName      = "bitbucket.pullrequests.getComment"
	PullRequestsListActivityLogActivityName = "bitbucket.pullrequests.listActivityLog"
	PullRequestsListCommitsActivityName     = "bitbucket.pullrequests.listCommits"
	PullRequestsListForCommitActivityName   = "bitbucket.pullrequests.listForCommit"
	PullRequestsListTasksActivityName       = "bitbucket.pullrequests.listTasks"
	PullRequestsMergeActivityName           = "bitbucket.pullrequests.merge"
	PullRequestsUnapproveActivityName       = "bitbucket.pullrequests.unapprove"
	PullRequestsUpdateActivityName          = "bitbucket.pullrequests.update"
	PullRequestsUpdateCommentActivityName   = "bitbucket.pullrequests.updateComment"
) //revive:enable:exported

// PullRequestsRequest contains common fields for PR-related requests.
type PullRequestsRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace     string `json:"workspace"`
	RepoSlug      string `json:"repo_slug"`
	PullRequestID string `json:"pull_request_id"`
}

// PullRequestsApproveRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-approve-post
type PullRequestsApproveRequest = PullRequestsRequest

// PullRequestsApprove is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-approve-post
func PullRequestsApprove(ctx workflow.Context, thrippyLinkID, workspace, repo, prID string) error {
	req := PullRequestsApproveRequest{ThrippyLinkID: thrippyLinkID, Workspace: workspace, RepoSlug: repo, PullRequestID: prID}
	return internal.ExecuteTimpaniActivityNoResp(ctx, PullRequestsApproveActivityName, req)
}

// PullRequestsCreateCommentRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
type PullRequestsCreateCommentRequest struct {
	PullRequestsRequest

	Markdown string `json:"text"`

	ParentID string `json:"parent_id,omitempty"`
}

// PullRequestsCreateCommentResponse is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
type PullRequestsCreateCommentResponse = Comment

// PullRequestsCreateComment is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-post
func PullRequestsCreateComment(ctx workflow.Context, req PullRequestsCreateCommentRequest) (*PullRequestsCreateCommentResponse, error) {
	return internal.ExecuteTimpaniActivity[PullRequestsCreateCommentResponse](ctx, PullRequestsCreateCommentActivityName, req)
}

// PullRequestsDeclineRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-decline-post
type PullRequestsDeclineRequest = PullRequestsRequest

// PullRequestsDecline is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-decline-post
func PullRequestsDecline(ctx workflow.Context, thrippyLinkID, workspace, repo, prID string) error {
	req := PullRequestsDeclineRequest{ThrippyLinkID: thrippyLinkID, Workspace: workspace, RepoSlug: repo, PullRequestID: prID}
	return internal.ExecuteTimpaniActivityNoResp(ctx, PullRequestsDeclineActivityName, req)
}

// PullRequestsDeleteCommentRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-delete
type PullRequestsDeleteCommentRequest struct {
	PullRequestsRequest

	CommentID string `json:"comment_id"`
}

// PullRequestsDeleteComment is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-delete
func PullRequestsDeleteComment(ctx workflow.Context, thrippyLinkID, workspace, repo, prID, commentID string) error {
	pr := PullRequestsRequest{ThrippyLinkID: thrippyLinkID, Workspace: workspace, RepoSlug: repo, PullRequestID: prID}
	req := PullRequestsDeleteCommentRequest{PullRequestsRequest: pr, CommentID: commentID}
	return internal.ExecuteTimpaniActivityNoResp(ctx, PullRequestsDeleteCommentActivityName, req)
}

// PullRequestsDiffstatRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-diffstat-get
type PullRequestsDiffstatRequest struct {
	PullRequestsRequest

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen string `json:"pagelen,omitempty"`
	Page    string `json:"page,omitempty"`

	Next string `json:"next,omitempty"` // Populated and used only in Timpani, for pagination.
}

// PullRequestsDiffstatResponse is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-diffstat-get
type PullRequestsDiffstatResponse = CommitsDiffstatResponse

// PullRequestsDiffstat is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-diffstat-get
//
// It retrieves the full list of diffstat entries by handling pagination internally.
func PullRequestsDiffstat(ctx workflow.Context, thrippyLinkID, workspace, repo, prID string) ([]Diffstat, error) {
	pr := PullRequestsRequest{ThrippyLinkID: thrippyLinkID, Workspace: workspace, RepoSlug: repo, PullRequestID: prID}
	req := PullRequestsDiffstatRequest{PullRequestsRequest: pr, Next: "start"}

	var ds []Diffstat
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

// PullRequestsGetRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-get
type PullRequestsGetRequest = PullRequestsRequest

// PullRequestsGet is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-get
func PullRequestsGet(ctx workflow.Context, thrippyLinkID, workspace, repo, prID string) (map[string]any, error) {
	req := PullRequestsGetRequest{ThrippyLinkID: thrippyLinkID, Workspace: workspace, RepoSlug: repo, PullRequestID: prID}
	resp, err := internal.ExecuteTimpaniActivity[map[string]any](ctx, PullRequestsGetActivityName, req)
	if err != nil {
		return nil, err
	}
	return *resp, nil
}

// PullRequestsGetCommentRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-get
type PullRequestsGetCommentRequest struct {
	PullRequestsRequest

	CommentID string `json:"comment_id"`
}

// PullRequestsGetCommentResponse is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-get
type PullRequestsGetCommentResponse = Comment

// PullRequestsGetComment is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-get
func PullRequestsGetComment(ctx workflow.Context, thrippyLinkID, workspace, repo, prID, commentID string) (*Comment, error) {
	pr := PullRequestsRequest{ThrippyLinkID: thrippyLinkID, Workspace: workspace, RepoSlug: repo, PullRequestID: prID}
	req := PullRequestsGetCommentRequest{PullRequestsRequest: pr, CommentID: commentID}
	return internal.ExecuteTimpaniActivity[Comment](ctx, PullRequestsGetCommentActivityName, req)
}

// PullRequestsListActivityLogRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-activity-get
type PullRequestsListActivityLogRequest struct {
	PullRequestsRequest

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen string `json:"pagelen,omitempty"`
	Page    string `json:"page,omitempty"`

	Next string `json:"next,omitempty"` // Populated and used only in Timpani, for pagination.
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
//
// It retrieves the full list of activity log entries by handling pagination internally.
func PullRequestsListActivityLog(ctx workflow.Context, thrippyLinkID, workspace, repo, prID string) ([]map[string]any, error) {
	pr := PullRequestsRequest{ThrippyLinkID: thrippyLinkID, Workspace: workspace, RepoSlug: repo, PullRequestID: prID}
	req := PullRequestsListActivityLogRequest{PullRequestsRequest: pr, Next: "start"}

	var activities []map[string]any
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
	PullRequestsRequest

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen string `json:"pagelen,omitempty"`
	Page    string `json:"page,omitempty"`

	Next string `json:"next,omitempty"` // Populated and used only in Timpani, for pagination.
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
//
// It retrieves the full list of commits by handling pagination internally.
func PullRequestsListCommits(ctx workflow.Context, thrippyLinkID, workspace, repo, prID string) ([]Commit, error) {
	pr := PullRequestsRequest{ThrippyLinkID: thrippyLinkID, Workspace: workspace, RepoSlug: repo, PullRequestID: prID}
	req := PullRequestsListCommitsRequest{PullRequestsRequest: pr, Next: "start"}

	var cs []Commit
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

	Next string `json:"next,omitempty"` // Populated and used only in Timpani, for pagination.
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
//
// It retrieves the full list of PRs by handling pagination internally.
func PullRequestsListForCommit(ctx workflow.Context, thrippyLinkID, workspace, repo, commit string) ([]map[string]any, error) {
	req := PullRequestsListForCommitRequest{ThrippyLinkID: thrippyLinkID, Workspace: workspace, RepoSlug: repo, Commit: commit, Next: "start"}

	var prs []map[string]any
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

// PullRequestsListTasksRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-tasks-get
type PullRequestsListTasksRequest struct {
	PullRequestsRequest

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen string `json:"pagelen,omitempty"`
	Page    string `json:"page,omitempty"`

	Next string `json:"next,omitempty"` // Populated and used only in Timpani, for pagination.
}

// PullRequestsListTasksResponse is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-tasks-get
type PullRequestsListTasksResponse struct {
	Values []Task `json:"values"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	Size    int    `json:"size,omitempty"`
	PageLen int    `json:"pagelen,omitempty"`
	Page    int    `json:"page,omitempty"`
	Next    string `json:"next,omitempty"`
}

// PullRequestsListTasks is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-tasks-get
//
// It retrieves the full list of tasks by handling pagination internally.
func PullRequestsListTasks(ctx workflow.Context, thrippyLinkID, workspace, repo, prID string) ([]Task, error) {
	pr := PullRequestsRequest{ThrippyLinkID: thrippyLinkID, Workspace: workspace, RepoSlug: repo, PullRequestID: prID}
	req := PullRequestsListTasksRequest{PullRequestsRequest: pr, Next: "start"}

	var tasks []Task
	for req.Next != "" {
		if req.Next == "start" {
			req.Next = ""
		}

		resp, err := internal.ExecuteTimpaniActivity[PullRequestsListTasksResponse](ctx, PullRequestsListTasksActivityName, req)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, resp.Values...)
		req.Next = resp.Next
	}

	return tasks, nil
}

// PullRequestsMergeRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-merge-post
type PullRequestsMergeRequest struct {
	PullRequestsRequest

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
type PullRequestsUnapproveRequest = PullRequestsRequest

// PullRequestsUnapprove is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-approve-delete
func PullRequestsUnapprove(ctx workflow.Context, thrippyLinkID, workspace, repo, prID string) error {
	req := PullRequestsUnapproveRequest{ThrippyLinkID: thrippyLinkID, Workspace: workspace, RepoSlug: repo, PullRequestID: prID}
	return internal.ExecuteTimpaniActivityNoResp(ctx, PullRequestsUnapproveActivityName, req)
}

// PullRequestsUpdateRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-put
type PullRequestsUpdateRequest struct {
	PullRequestsRequest

	PullRequest map[string]any `json:"pullrequest"`
}

// PullRequestsUpdate is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-put
func PullRequestsUpdate(ctx workflow.Context, thrippyLinkID, workspace, repo, prID string, update map[string]any) (map[string]any, error) {
	pr := PullRequestsRequest{ThrippyLinkID: thrippyLinkID, Workspace: workspace, RepoSlug: repo, PullRequestID: prID}
	req := PullRequestsUpdateRequest{PullRequestsRequest: pr, PullRequest: update}
	resp, err := internal.ExecuteTimpaniActivity[map[string]any](ctx, PullRequestsUpdateActivityName, req)
	if err != nil {
		return nil, err
	}
	return *resp, nil
}

// PullRequestsUpdateCommentRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-put
type PullRequestsUpdateCommentRequest struct {
	PullRequestsRequest

	CommentID string `json:"comment_id"`
	Markdown  string `json:"text"`
}

// PullRequestsUpdateComment is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-comments-comment-id-put
func PullRequestsUpdateComment(ctx workflow.Context, thrippyLinkID, workspace, repo, prID, commentID, markdown string) error {
	pr := PullRequestsRequest{ThrippyLinkID: thrippyLinkID, Workspace: workspace, RepoSlug: repo, PullRequestID: prID}
	req := PullRequestsUpdateCommentRequest{PullRequestsRequest: pr, CommentID: commentID, Markdown: markdown}
	return internal.ExecuteTimpaniActivityNoResp(ctx, PullRequestsUpdateCommentActivityName, req)
}

// Comment is based on:
// https://support.atlassian.com/bitbucket-cloud/docs/event-payloads/#Comment
type Comment struct {
	// Type string `json:"type"` // Always "pullrequest_comment".

	ID      int      `json:"id"`
	Content Rendered `json:"content"`
	Inline  *Inline  `json:"inline,omitempty"`
	User    User     `json:"user"`

	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on,omitzero"`
	Deleted   bool      `json:"deleted,omitempty"`
	Pending   bool      `json:"pending,omitempty"`

	Parent *Parent `json:"parent,omitempty"`

	// PullRequest `json:"pullrequest"` // Unnecessary.

	Links map[string]Link `json:"links"`
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

// Inline is based on:
// https://support.atlassian.com/bitbucket-cloud/docs/event-payloads/#Comment
type Inline struct {
	Path string `json:"path"`

	StartFrom *int `json:"start_from,omitempty"`
	StartTo   *int `json:"start_to,omitempty"`
	From      *int `json:"from,omitempty"`
	To        *int `json:"to,omitempty"`

	ContextLines string `json:"context_lines"`
	Outdated     bool   `json:"outdated,omitempty"`

	SrcRev  string `json:"src_rev"`
	DestRev string `json:"dest_rev"`

	// What is "base_rev"? It is useful?
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

//revive:disable:exported
type Rendered struct {
	// Type string `json:"type"` // Always "rendered".

	Raw    string `json:"raw"`
	Markup string `json:"markup"`
	HTML   string `json:"html"`
} //revive:enable:exported

// Task is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-pullrequests/#api-repositories-workspace-repo-slug-pullrequests-pull-request-id-tasks-get
type Task struct {
	ID      int      `json:"id"`
	State   string   `json:"state"`
	Content Rendered `json:"content"`
	Creator User     `json:"creator"`

	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on,omitzero"`
	Pending   bool      `json:"pending,omitempty"`

	ResolvedBy *User     `json:"resolved_by,omitempty"`
	ResolvedOn time.Time `json:"resolved_on,omitzero"`

	// Comment                              // Unnecessary.
	// Links map[string]Link `json:"links"` // Unnecessary.
}
