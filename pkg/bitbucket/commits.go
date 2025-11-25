package bitbucket

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	CommitsDiffActivityName     = "bitbucket.commits.diff"
	CommitsDiffStatActivityName = "bitbucket.commits.diffstat"
)

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diff-spec-get
type CommitsDiffRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace string `json:"workspace"`
	RepoSlug  string `json:"repo_slug"`
	Spec      string `json:"spec"`

	Path string `json:"path,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diff-spec-get
func CommitsDiffActivity(ctx workflow.Context, req CommitsDiffRequest) (string, error) {
	fut := internal.ExecuteTimpaniActivity(ctx, CommitsDiffActivityName, req)

	resp := new(string)
	if err := fut.Get(ctx, resp); err != nil {
		return "", err
	}

	return *resp, nil
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diffstat-spec-get
type CommitsDiffStatRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace string `json:"workspace"`
	RepoSlug  string `json:"repo_slug"`
	Spec      string `json:"spec"`

	IgnoreWhitespace bool   `json:"ignore_whitespace,omitempty"`
	Merge            bool   `json:"merge,omitempty"`
	Renames          bool   `json:"renames,omitempty"`
	Topic            bool   `json:"topic,omitempty"`
	Path             string `json:"path,omitempty"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen  string `json:"pagelen,omitempty"`
	Page     string `json:"page,omitempty"`
	AllPages bool   `json:"all_pages,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diffstat-spec-get
type CommitsDiffStatResponse struct {
	Values []DiffStat `json:"values"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	Size    int    `json:"size,omitempty"`
	PageLen int    `json:"pagelen,omitempty"`
	Page    int    `json:"page,omitempty"`
	Next    string `json:"next,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diffstat-spec-get
func CommitsDiffStatActivity(ctx workflow.Context, req CommitsDiffStatRequest) (*CommitsDiffStatResponse, error) {
	fut := internal.ExecuteTimpaniActivity(ctx, CommitsDiffStatActivityName, req)

	resp := new(CommitsDiffStatResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diffstat-spec-get
type DiffStat struct {
	// Type string `json:"type"` // Always "diffstat".

	Status string `json:"status"`

	LinesAdded   int `json:"lines_added"`
	LinesRemoved int `json:"lines_removed"`

	Old *CommitFile `json:"old,omitempty"`
	New *CommitFile `json:"new,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diffstat-spec-get
type CommitFile struct {
	// Type string `json:"type"` // Always "commit_file".

	Path        string `json:"path"`
	EscapedPath string `json:"escaped_path"`

	Links map[string]Link `json:"links"`
}
