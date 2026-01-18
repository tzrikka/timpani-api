package github

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

//revive:disable:exported
const (
	PullRequestsGetActivityName         = "github.pulls.get"
	PullRequestsListCommitsActivityName = "github.pulls.listCommits"
	PullRequestsListFilesActivityName   = "github.pulls.listFiles"
	PullRequestsMergeActivityName       = "github.pulls.merge"
	PullRequestsUpdateActivityName      = "github.pulls.update"

	PullRequestsCommentsCreateActivityName      = "github.pulls.reviewComments.create"
	PullRequestsCommentsCreateReplyActivityName = "github.pulls.reviewComments.createReply"
	PullRequestsCommentsDeleteActivityName      = "github.pulls.reviewComments.delete"
	PullRequestsCommentsUpdateActivityName      = "github.pulls.reviewComments.update"

	PullRequestsReviewsCreateActivityName  = "github.pulls.reviews.create"
	PullRequestsReviewsDeleteActivityName  = "github.pulls.reviews.deletePending"
	PullRequestsReviewsDismissActivityName = "github.pulls.reviews.dismiss"
	PullRequestsReviewsSubmitActivityName  = "github.pulls.reviews.submitPending"
	PullRequestsReviewsUpdateActivityName  = "github.pulls.reviews.update"
) //revive:enable:exported

// PullRequestsRequest contains common fields for PR-related requests.
type PullRequestsRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Owner      string `json:"owner,omitempty"`
	Repo       string `json:"repo,omitempty"`
	PullNumber int    `json:"pull_number,omitempty"`
}

// PullRequestsCommentsRequest contains common fields for PR comment-related requests.
type PullRequestsCommentsRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Owner     string `json:"owner,omitempty"`
	Repo      string `json:"repo,omitempty"`
	CommentID int    `json:"comment_id,omitempty"`
}

// PullRequestsReviewsRequest contains common fields for PR review-related requests.
type PullRequestsReviewsRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Owner      string `json:"owner,omitempty"`
	Repo       string `json:"repo,omitempty"`
	PullNumber int    `json:"pull_number,omitempty"`
	ReviewID   int    `json:"review_id,omitempty"`
}

// PullRequestsGetRequest is based on:
// https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#get-a-pull-request
type PullRequestsGetRequest = PullRequestsRequest

// PullRequestsGet is based on:
// https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#get-a-pull-request
func PullRequestsGet(ctx workflow.Context, thrippyLinkID, owner, repo string, prID int) (*PullRequest, error) {
	req := PullRequestsGetRequest{ThrippyLinkID: thrippyLinkID, Owner: owner, Repo: repo, PullNumber: prID}
	return internal.ExecuteTimpaniActivity[PullRequest](ctx, PullRequestsGetActivityName, req)
}

// PullRequestsListCommitsRequest is based on:
// https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#list-commits-on-a-pull-request
type PullRequestsListCommitsRequest struct {
	PullRequestsRequest

	// https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api
	PerPage int `json:"per_page,omitempty"`
	Page    int `json:"page,omitempty"`
}

// PullRequestsListFilesRequest is based on:
// https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#list-pull-requests-files
type PullRequestsListFilesRequest struct {
	PullRequestsRequest

	// https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api
	PerPage int `json:"per_page,omitempty"`
	Page    int `json:"page,omitempty"`
}

// PullRequestsMergeRequest is based on:
// https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#merge-a-pull-request
type PullRequestsMergeRequest struct {
	PullRequestsRequest

	CommitTitle   string `json:"commit_title,omitempty"`
	CommitMessage string `json:"commit_message,omitempty"`
	SHA           string `json:"sha,omitempty"`
	MergeMethod   string `json:"merge_method,omitempty"` // "merge", "squash", "rebase".
}

// PullRequestsMergeResponse is based on:
// https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#merge-a-pull-request
type PullRequestsMergeResponse struct {
	Merged  bool   `json:"merged"`
	Message string `json:"message"`
	SHA     string `json:"sha"`
}

// PullRequestsMerge is based on:
// https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#merge-a-pull-request
func PullRequestsMerge(ctx workflow.Context, req PullRequestsMergeRequest) (*PullRequestsMergeResponse, error) {
	return internal.ExecuteTimpaniActivity[PullRequestsMergeResponse](ctx, PullRequestsMergeActivityName, req)
}

// PullRequestsUpdateRequest is based on:
// https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#update-a-pull-request
type PullRequestsUpdateRequest struct {
	PullRequestsRequest

	Title               string `json:"title,omitempty"`
	Body                string `json:"body,omitempty"`
	State               string `json:"state,omitempty"` // "open", "closed".
	Base                string `json:"base,omitempty"`
	MaintainerCanModify bool   `json:"maintainer_can_modify,omitempty"`
}

// PullRequestsCommentsCreateRequest is based on:
// https://docs.github.com/en/rest/pulls/comments?apiVersion=2022-11-28#create-a-review-comment-for-a-pull-request
type PullRequestsCommentsCreateRequest struct {
	PullRequestsRequest

	Body string `json:"body"`

	CommitID string `json:"commit_id"`
	Path     string `json:"path"`

	InReplyTo   int    `json:"in_reply_to,omitempty"`
	SubjectType string `json:"subject_type,omitempty"` // "line", "file".
	StartSide   string `json:"start_side,omitempty"`   // "LEFT", "RIGHT".
	StartLine   int    `json:"start_line,omitempty"`
	Side        string `json:"side,omitempty"` // "LEFT", "RIGHT".
	Line        int    `json:"line,omitempty"`
}

// PullRequestsCommentsCreate is based on:
// https://docs.github.com/en/rest/pulls/comments?apiVersion=2022-11-28#create-a-review-comment-for-a-pull-request
func PullRequestsCommentsCreate(ctx workflow.Context, req PullRequestsCommentsCreateRequest) (*PullComment, error) {
	return internal.ExecuteTimpaniActivity[PullComment](ctx, PullRequestsCommentsCreateActivityName, req)
}

// PullRequestsCommentsCreateReplyRequest is based on:
// https://docs.github.com/en/rest/pulls/comments?apiVersion=2022-11-28#create-a-reply-for-a-review-comment
type PullRequestsCommentsCreateReplyRequest struct {
	PullRequestsRequest

	CommentID int    `json:"comment_id,omitempty"`
	Body      string `json:"body"`
}

// PullRequestsCommentsCreateReply is based on:
// https://docs.github.com/en/rest/pulls/comments?apiVersion=2022-11-28#create-a-reply-for-a-review-comment
func PullRequestsCommentsCreateReply(ctx workflow.Context, req PullRequestsCommentsCreateReplyRequest) (*PullComment, error) {
	return internal.ExecuteTimpaniActivity[PullComment](ctx, PullRequestsCommentsCreateReplyActivityName, req)
}

// PullRequestsCommentsDeleteRequest is based on:
// https://docs.github.com/en/rest/pulls/comments?apiVersion=2022-11-28#delete-a-review-comment-for-a-pull-request
type PullRequestsCommentsDeleteRequest = PullRequestsCommentsRequest

// PullRequestsCommentsDelete is based on:
// https://docs.github.com/en/rest/pulls/comments?apiVersion=2022-11-28#delete-a-review-comment-for-a-pull-request
func PullRequestsCommentsDelete(ctx workflow.Context, thrippyLinkID, owner, repo string, commentID int) error {
	req := PullRequestsCommentsDeleteRequest{ThrippyLinkID: thrippyLinkID, Owner: owner, Repo: repo, CommentID: commentID}
	return internal.ExecuteTimpaniActivityNoResp(ctx, PullRequestsCommentsDeleteActivityName, req)
}

// PullRequestsCommentsUpdateRequest is based on:
// https://docs.github.com/en/rest/pulls/comments?apiVersion=2022-11-28#update-a-review-comment-for-a-pull-request
type PullRequestsCommentsUpdateRequest struct {
	PullRequestsCommentsRequest

	Body string `json:"body"`
}

// PullRequestsCommentsUpdate is based on:
// https://docs.github.com/en/rest/pulls/comments?apiVersion=2022-11-28#update-a-review-comment-for-a-pull-request
func PullRequestsCommentsUpdate(ctx workflow.Context, thrippyLinkID, owner, repo string, commentID int, body string) (*PullComment, error) {
	comment := PullRequestsCommentsRequest{ThrippyLinkID: thrippyLinkID, Owner: owner, Repo: repo, CommentID: commentID}
	req := PullRequestsCommentsUpdateRequest{PullRequestsCommentsRequest: comment, Body: body}
	return internal.ExecuteTimpaniActivity[PullComment](ctx, PullRequestsCommentsUpdateActivityName, req)
}

// PullRequestsReviewsCreateRequest is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#create-a-review-for-a-pull-request
type PullRequestsReviewsCreateRequest struct {
	PullRequestsRequest

	CommitID string           `json:"commit_id,omitempty"`
	Body     string           `json:"body,omitempty"`
	Event    string           `json:"event,omitempty"` // "APPROVE", "REQUEST_CHANGES", "COMMENT", "" = "PENDING".
	Comments []map[string]any `json:"comments,omitempty"`
}

// PullRequestsReviewsCreate is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#create-a-review-for-a-pull-request
func PullRequestsReviewsCreate(ctx workflow.Context, req PullRequestsReviewsCreateRequest) (*Review, error) {
	return internal.ExecuteTimpaniActivity[Review](ctx, PullRequestsReviewsCreateActivityName, req)
}

// PullRequestsReviewsDeleteRequest is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#delete-a-pending-review-for-a-pull-request
type PullRequestsReviewsDeleteRequest = PullRequestsReviewsRequest

// PullRequestsReviewsDelete is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#delete-a-pending-review-for-a-pull-request
func PullRequestsReviewsDelete(ctx workflow.Context, thrippyLinkID, owner, repo string, prID, reviewID int) error {
	req := PullRequestsReviewsDeleteRequest{ThrippyLinkID: thrippyLinkID, Owner: owner, Repo: repo, PullNumber: prID, ReviewID: reviewID}
	return internal.ExecuteTimpaniActivityNoResp(ctx, PullRequestsReviewsDeleteActivityName, req)
}

// PullRequestsReviewsDismissRequest is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#dismiss-a-review-for-a-pull-request
type PullRequestsReviewsDismissRequest struct {
	PullRequestsReviewsRequest

	Message string `json:"message"`
	Event   string `json:"event,omitempty"` // "DISMISS", "".
}

// PullRequestsReviewsDismiss is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#dismiss-a-review-for-a-pull-request
func PullRequestsReviewsDismiss(ctx workflow.Context, req PullRequestsReviewsDismissRequest) (*Review, error) {
	return internal.ExecuteTimpaniActivity[Review](ctx, PullRequestsReviewsDismissActivityName, req)
}

// PullRequestsReviewsSubmitRequest is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#submit-a-review-for-a-pull-request
type PullRequestsReviewsSubmitRequest struct {
	PullRequestsReviewsRequest

	Body  string `json:"body,omitempty"`
	Event string `json:"event,omitempty"` // "APPROVE", "REQUEST_CHANGES", "COMMENT".
}

// PullRequestsReviewsSubmit is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#submit-a-review-for-a-pull-request
func PullRequestsReviewsSubmit(ctx workflow.Context, req PullRequestsReviewsSubmitRequest) (*Review, error) {
	return internal.ExecuteTimpaniActivity[Review](ctx, PullRequestsReviewsSubmitActivityName, req)
}

// PullRequestsReviewsUpdateRequest is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#update-a-review-for-a-pull-request
type PullRequestsReviewsUpdateRequest struct {
	PullRequestsReviewsRequest

	Body string `json:"body"`
}

// PullRequestsReviewsUpdate is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#update-a-review-for-a-pull-request
func PullRequestsReviewsUpdate(ctx workflow.Context, req PullRequestsReviewsUpdateRequest) (*Review, error) {
	return internal.ExecuteTimpaniActivity[Review](ctx, PullRequestsReviewsUpdateActivityName, req)
}

// AutoMerge is used in [PullRequest].
type AutoMerge struct {
	EnabledBy     User   `json:"enabled_by"`
	MergeMethod   string `json:"merge_method"`
	CommitTitle   string `json:"commit_title"`
	CommitMessage string `json:"commit_message"`
}

// Branch is used in [PullRequest].
type Branch struct {
	Label string     `json:"label"`
	Ref   string     `json:"ref"`
	SHA   string     `json:"sha"`
	Repo  Repository `json:"repo"`
	User  User       `json:"user"`
}

// PullRequest is based on:
//   - https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28
//   - https://docs.github.com/en/webhooks/webhook-events-and-payloads#pull_request
type PullRequest struct {
	ID       int64  `json:"id"`
	NodeID   string `json:"node_id"`
	HTMLURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`

	Number int    `json:"number"`
	State  string `json:"state"` // "open" or "closed".
	Title  string `json:"title"`
	Body   string `json:"body"`

	Draft            bool   `json:"draft,omitempty"`
	Locked           bool   `json:"locked,omitempty"`
	ActiveLockReason string `json:"active_lock_reason,omitempty"` // "off_topic", "too_heated", "resolved", "spam".
	// Labels        []Label    `json:"labels,omitempty"`
	// Milestone     *Milestone `json:"milestone,omitempty"`.

	Comments       int `json:"comments,omitempty"`
	ReviewComments int `json:"review_comments,omitempty"`
	Commits        int `json:"commits,omitempty"`
	Additions      int `json:"additions,omitempty"`
	Deletions      int `json:"deletions,omitempty"`
	ChangedFiles   int `json:"changed_files,omitempty"`

	Head Branch `json:"head"`
	Base Branch `json:"base"`

	User               User   `json:"user"`
	Assignee           *User  `json:"assignee,omitempty"`
	Assignees          []User `json:"assignees,omitempty"`
	RequestedReviewers []User `json:"requested_reviewers,omitempty"`
	RequestedTeams     []Team `json:"requested_teams,omitempty"`

	AutoMerge      *AutoMerge `json:"auto_merge,omitempty"`
	Mergeable      bool       `json:"mergeable,omitempty"`
	MergeableState string     `json:"mergeable_state,omitempty"` // "clean", "dirty", etc.
	Rebaseable     bool       `json:"rebaseable,omitempty"`

	CreatedAt time.Time `json:"created_at,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
	ClosedAt  time.Time `json:"closed_at,omitzero"`
	MergedAt  time.Time `json:"merged_at,omitzero"`

	Merged         bool   `json:"merged,omitempty"`
	MergedBy       *User  `json:"merged_by,omitempty"`
	MergeCommitSHA string `json:"merge_commit_sha,omitempty"`
}

// PullComment is based on:
//   - https://docs.github.com/en/rest/pulls/comments?apiVersion=2022-11-28
//   - https://docs.github.com/en/webhooks/webhook-events-and-payloads#pull_request_review_comment
type PullComment struct {
	ID      int    `json:"id"`
	NodeID  string `json:"node_id"`
	HTMLURL string `json:"html_url"`

	PullRequestReviewID int    `json:"pull_request_review_id"`
	InReplyTo           *int   `json:"in_reply_to_id,omitempty"`
	CommitID            string `json:"commit_id"`
	OriginalCommitID    string `json:"original_commit_id"`

	Path        string `json:"path"`
	SubjectType string `json:"subject_type,omitempty"`
	DiffHunk    string `json:"diff_hunk,omitempty"`

	User      User       `json:"user"`
	Body      string     `json:"body"`
	Reactions *Reactions `json:"reactions,omitempty"`

	StartLine         int    `json:"start_line,omitzero"`
	OriginalStartLine int    `json:"original_start_line,omitzero"`
	StartSide         string `json:"start_side,omitzero"`
	Line              int    `json:"line,omitzero"`
	OriginalLine      int    `json:"original_line,omitzero"`
	Side              string `json:"side,omitzero"`

	CreatedAt time.Time `json:"created_at,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
}

// Repository is based on:
//   - https://docs.github.com/en/rest/repos/repos?apiVersion=2022-11-28
//   - https://docs.github.com/en/webhooks/webhook-events-and-payloads#repository
type Repository struct {
	ID       int64  `json:"id"`
	NodeID   string `json:"node_id"`
	HTMLURL  string `json:"html_url"`
	CloneURL string `json:"clone_url,omitempty"`
	GitURL   string `json:"git_url,omitempty"`
	SSHURL   string `json:"ssh_url,omitempty"`

	Name        string  `json:"name"`
	FullName    string  `json:"full_name"`
	Description *string `json:"description,omitempty"`
	Owner       *User   `json:"owner,omitempty"`

	DefaultBranch string `json:"default_branch,omitempty"`
	Private       bool   `json:"private"`
	Fork          bool   `json:"fork"`

	PushedAt time.Time `json:"pushed_at,omitzero"`
}

// Review is based on:
//   - https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28
//   - https://docs.github.com/en/webhooks/webhook-events-and-payloads#pull_request_review
type Review struct {
	ID      int    `json:"id"`
	NodeID  string `json:"node_id"`
	HTMLURL string `json:"html_url"`

	User     User   `json:"user"`
	State    string `json:"state"`
	Body     string `json:"body"`
	CommitID string `json:"commit_id"`

	SubmittedAt time.Time `json:"submitted_at,omitzero"`
	UpdatedAt   time.Time `json:"updated_at,omitzero"`
}
