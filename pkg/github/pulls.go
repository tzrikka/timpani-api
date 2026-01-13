package github

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

	Owner      string `json:"owner"`
	Repo       string `json:"repo"`
	PullNumber int    `json:"pull_number"`
}

// PullRequestsCommentsRequest contains common fields for PR comment-related requests.
type PullRequestsCommentsRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Owner     string `json:"owner"`
	Repo      string `json:"repo"`
	CommentID int    `json:"comment_id"`
}

// PullRequestsGetRequest is based on:
// https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#get-a-pull-request
type PullRequestsGetRequest = PullRequestsRequest

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

// PullRequestsCommentsCreateReplyRequest is based on:
// https://docs.github.com/en/rest/pulls/comments?apiVersion=2022-11-28#create-a-reply-for-a-review-comment
type PullRequestsCommentsCreateReplyRequest struct {
	PullRequestsRequest

	CommentID int    `json:"comment_id"`
	Body      string `json:"body,omitempty"`
}

// PullRequestsCommentsDeleteRequest is based on:
// https://docs.github.com/en/rest/pulls/comments?apiVersion=2022-11-28#delete-a-review-comment-for-a-pull-request
type PullRequestsCommentsDeleteRequest = PullRequestsCommentsRequest

// PullRequestsCommentsUpdateRequest is based on:
// https://docs.github.com/en/rest/pulls/comments?apiVersion=2022-11-28#update-a-review-comment-for-a-pull-request
type PullRequestsCommentsUpdateRequest struct {
	PullRequestsCommentsRequest

	Body string `json:"body"`
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

// PullRequestsReviewsDeleteRequest is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#delete-a-pending-review-for-a-pull-request
type PullRequestsReviewsDeleteRequest struct {
	PullRequestsRequest

	ReviewID int `json:"review_id"`
}

// PullRequestsReviewsDismissRequest is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#dismiss-a-review-for-a-pull-request
type PullRequestsReviewsDismissRequest struct {
	PullRequestsRequest

	ReviewID int    `json:"review_id"`
	Message  string `json:"message"`

	Event string `json:"event,omitempty"` // "DISMISS", "".
}

// PullRequestsReviewsSubmitRequest is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#submit-a-review-for-a-pull-request
type PullRequestsReviewsSubmitRequest = PullRequestsReviewsDismissRequest

// PullRequestsReviewsUpdateRequest is based on:
// https://docs.github.com/en/rest/pulls/reviews?apiVersion=2022-11-28#update-a-review-for-a-pull-request
type PullRequestsReviewsUpdateRequest struct {
	PullRequestsRequest

	ReviewID int    `json:"review_id"`
	Event    string `json:"event"` // "APPROVE", "REQUEST_CHANGES", "COMMENT".

	Body string `json:"body,omitempty"`
}
