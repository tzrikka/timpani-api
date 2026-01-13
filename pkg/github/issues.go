package github

import (
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

// IssuesCommentsCreateResponse is based on:
// https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28#create-an-issue-comment
type IssuesCommentsCreateResponse = IssueComment

// IssuesCommentsCreate is based on:
// https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28#create-an-issue-comment
func IssuesCommentsCreate(ctx workflow.Context, thrippyLinkID, owner, repo, body string, issue int) (*IssueComment, error) {
	req := IssuesCommentsCreateRequest{ThrippyLinkID: thrippyLinkID, Owner: owner, Repo: repo, IssueNumber: issue, Body: body}
	return internal.ExecuteTimpaniActivity[IssuesCommentsCreateResponse](ctx, IssuesCommentsCreateActivityName, req)
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

// IssuesCommentsUpdateResponse is based on:
// https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28#update-an-issue-comment
type IssuesCommentsUpdateResponse = IssueComment

// IssuesCommentsUpdate is based on:
// https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28#update-an-issue-comment
func IssuesCommentsUpdate(ctx workflow.Context, thrippyLinkID, owner, repo, body string, commentID int) (*IssueComment, error) {
	req := IssuesCommentsUpdateRequest{ThrippyLinkID: thrippyLinkID, Owner: owner, Repo: repo, CommentID: commentID, Body: body}
	return internal.ExecuteTimpaniActivity[IssuesCommentsUpdateResponse](ctx, IssuesCommentsUpdateActivityName, req)
}

// IssueComment is based on:
// https://docs.github.com/en/rest/issues/comments?apiVersion=2022-11-28#get-an-issue-comment
type IssueComment struct {
	ID      int    `json:"id"`
	NodeID  string `json:"node_id"`
	HTMLURL string `json:"html_url"`

	Body      string     `json:"body"`
	Reactions *Reactions `json:"reactions,omitempty"`
	User      User       `json:"user"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
