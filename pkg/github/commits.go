package github

import "time"

// Commit is based on:
//   - https://docs.github.com/en/rest/commits/commits?apiVersion=2022-11-28
//   - https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#list-commits-on-a-pull-request
type Commit struct {
	SHA     string `json:"sha"`
	NodeID  string `json:"node_id"`
	HTMLURL string `json:"html_url"`

	Commit CommitDetails `json:"commit"`

	Author    User           `json:"author"`
	Committer User           `json:"committer"`
	Parents   []CommitParent `json:"parents"`
}

// CommitDetails is based on:
//   - https://docs.github.com/en/rest/commits/commits?apiVersion=2022-11-28
//   - https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#list-commits-on-a-pull-request
type CommitDetails struct {
	Message      string       `json:"message"`
	CommentCount int          `json:"comment_count"`
	Verification Verification `json:"verification"`

	// Not really necessary, duplicated in other fields.
	// Author    CommitUser `json:"author"`
	// Committer CommitUser `json:"committer"`
	// Tree      CommitTree `json:"tree"`.
}

// CommitUser is based on:
//   - https://docs.github.com/en/rest/commits/commits?apiVersion=2022-11-28
//   - https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#list-commits-on-a-pull-request
type CommitUser struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Username string    `json:"username,omitempty"`
	Date     time.Time `json:"date,omitzero"`
	When     time.Time `json:"when,omitzero"`
}

// CommitParent is based on:
//   - https://docs.github.com/en/rest/commits/commits?apiVersion=2022-11-28
//   - https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#list-commits-on-a-pull-request
type CommitParent struct {
	SHA     string `json:"sha"`
	HTMLURL string `json:"html_url"`
}

// Verification is based on:
//   - https://docs.github.com/en/rest/commits/commits?apiVersion=2022-11-28
//   - https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#list-commits-on-a-pull-request
type Verification struct {
	Verified   bool      `json:"verified"`
	Reason     string    `json:"reason"`
	Signature  string    `json:"signature,omitempty"`
	Payload    string    `json:"payload,omitempty"`
	VerifiedAt time.Time `json:"verified_at,omitzero"`
}
