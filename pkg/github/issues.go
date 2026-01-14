package github

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

//revive:disable:exported
const (
	IssuesCommentsCreateActivityName = "github.issues.comments.create"
	IssuesCommentsDeleteActivityName = "github.issues.comments.delete"
	IssuesCommentsUpdateActivityName = "github.issues.comments.update"
) //revive:enable:exported

// IssuesCommentsCreateRequest is based on:
// https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28#create-an-issue-comment
type IssuesCommentsCreateRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Owner       string `json:"owner"`
	Repo        string `json:"repo"`
	IssueNumber int    `json:"issue_number"`

	Body string `json:"body"`
}

// IssuesCommentsCreate is based on:
// https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28#create-an-issue-comment
func IssuesCommentsCreate(ctx workflow.Context, thrippyLinkID, owner, repo, body string, issue int) (*IssueComment, error) {
	req := IssuesCommentsCreateRequest{ThrippyLinkID: thrippyLinkID, Owner: owner, Repo: repo, IssueNumber: issue, Body: body}
	return internal.ExecuteTimpaniActivity[IssueComment](ctx, IssuesCommentsCreateActivityName, req)
}

// IssuesCommentsDeleteRequest is based on:
// https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28#delete-an-issue-comment
type IssuesCommentsDeleteRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Owner     string `json:"owner"`
	Repo      string `json:"repo"`
	CommentID int    `json:"comment_id"`
}

// IssuesCommentsDelete is based on:
// https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28#delete-an-issue-comment
func IssuesCommentsDelete(ctx workflow.Context, thrippyLinkID, owner, repo string, commentID int) error {
	req := IssuesCommentsDeleteRequest{ThrippyLinkID: thrippyLinkID, Owner: owner, Repo: repo, CommentID: commentID}
	return internal.ExecuteTimpaniActivityNoResp(ctx, IssuesCommentsDeleteActivityName, req)
}

// IssuesCommentsUpdateRequest is based on:
// https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28#update-an-issue-comment
type IssuesCommentsUpdateRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Owner     string `json:"owner"`
	Repo      string `json:"repo"`
	CommentID int    `json:"comment_id"`

	Body string `json:"body"`
}

// IssuesCommentsUpdate is based on:
// https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28#update-an-issue-comment
func IssuesCommentsUpdate(ctx workflow.Context, thrippyLinkID, owner, repo, body string, commentID int) (*IssueComment, error) {
	req := IssuesCommentsUpdateRequest{ThrippyLinkID: thrippyLinkID, Owner: owner, Repo: repo, CommentID: commentID, Body: body}
	return internal.ExecuteTimpaniActivity[IssueComment](ctx, IssuesCommentsUpdateActivityName, req)
}

// Issue is based on:
//   - https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28
//   - https://docs.github.com/en/webhooks/webhook-events-and-payloads#issue_comment
//   - https://docs.github.com/en/webhooks/webhook-events-and-payloads#issues
type Issue struct {
	ID      int    `json:"id"`
	NodeID  string `json:"node_id"`
	HTMLURL string `json:"html_url"`

	Number    int        `json:"number"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Reactions *Reactions `json:"reactions,omitempty"`

	State            string `json:"state"`                  // "open" or "closed".
	StateReason      string `json:"state_reason,omitempty"` // "completed", "not_planned".
	Draft            bool   `json:"draft,omitempty"`
	Locked           bool   `json:"locked,omitempty"`
	ActiveLockReason string `json:"active_lock_reason,omitempty"` // "off_topic", "too_heated", "resolved", "spam".
	// Labels        []Label    `json:"labels,omitempty"`
	// Milestone     *Milestone `json:"milestone,omitempty"`
	// Type          *IssueType `json:"type,omitempty"`.

	Comments int `json:"comments,omitempty"`

	PullRequest *PullRequestLinks `json:"pull_request,omitempty"`

	User      User   `json:"user"`
	Assignee  *User  `json:"assignee,omitempty"`
	Assignees []User `json:"assignees,omitempty"`
	ClosedBy  *User  `json:"closed_by,omitempty"`

	CreatedAt time.Time `json:"created_at,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
	ClosedAt  time.Time `json:"closed_at,omitzero"`
}

// IssueComment is based on:
//   - https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28
//   - https://docs.github.com/en/webhooks/webhook-events-and-payloads#issue_comment
type IssueComment struct {
	ID      int    `json:"id"`
	NodeID  string `json:"node_id"`
	HTMLURL string `json:"html_url"`

	Body      string     `json:"body"`
	Reactions *Reactions `json:"reactions,omitempty"`
	User      User       `json:"user"`

	CreatedAt time.Time `json:"created_at,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
}

// PullRequestLinks appears within an [Issue] when the issue is actually a [PullRequest].
// For example: https://docs.github.com/en/webhooks/webhook-events-and-payloads#issue_comment
type PullRequestLinks struct {
	URL      string `json:"url"`
	HTMLURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`

	MergedAt time.Time `json:"merged_at,omitzero"`
}
