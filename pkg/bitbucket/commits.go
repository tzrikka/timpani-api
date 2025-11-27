package bitbucket

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	CommitsDiffActivityName     = "bitbucket.commits.diff"
	CommitsDiffstatActivityName = "bitbucket.commits.diffstat"
)

// CommitsDiffRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diff-spec-get
type CommitsDiffRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace string `json:"workspace"`
	RepoSlug  string `json:"repo_slug"`
	Spec      string `json:"spec"`

	Path string `json:"path,omitempty"`
}

// CommitsDiff is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diff-spec-get
func CommitsDiff(ctx workflow.Context, req CommitsDiffRequest) (string, error) {
	resp, err := internal.ExecuteTimpaniActivity[string](ctx, CommitsDiffActivityName, req)
	if err != nil {
		return "", err
	}
	return *resp, nil
}

// CommitsDiffstatRequest is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diffstat-spec-get
type CommitsDiffstatRequest struct {
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
	PageLen string `json:"pagelen,omitempty"`
	Page    string `json:"page,omitempty"`
	Next    string `json:"next,omitempty"`
}

// CommitsDiffstatResponse is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diffstat-spec-get
type CommitsDiffstatResponse struct {
	Values []Diffstat `json:"values"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	Size    int `json:"size,omitempty"`
	PageLen int `json:"pagelen,omitempty"`
	Page    int `json:"page,omitempty"`

	Next string `json:"next,omitempty"`
}

// CommitsDiffstat is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diffstat-spec-get
func CommitsDiffstat(ctx workflow.Context, req CommitsDiffstatRequest) ([]Diffstat, error) {
	var ds []Diffstat
	req.Next = "start"

	for req.Next != "" {
		req.Next = ""
		resp, err := internal.ExecuteTimpaniActivity[CommitsDiffstatResponse](ctx, CommitsDiffstatActivityName, req)
		if err != nil {
			return nil, err
		}

		ds = append(ds, resp.Values...)
		req.Next = resp.Next
	}

	return ds, nil
}

// Diffstat is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diffstat-spec-get
type Diffstat struct {
	// Type string `json:"type"` // Always "diffstat".

	Status string `json:"status"`

	LinesAdded   int `json:"lines_added"`
	LinesRemoved int `json:"lines_removed"`

	Old *CommitFile `json:"old,omitempty"`
	New *CommitFile `json:"new,omitempty"`
}

// CommitFile is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-commits/#api-repositories-workspace-repo-slug-diffstat-spec-get
type CommitFile struct {
	// Type string `json:"type"` // Always "commit_file".

	Path        string `json:"path"`
	EscapedPath string `json:"escaped_path"`

	Links map[string]Link `json:"links"`
}
