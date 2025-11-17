package bitbucket

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	SourceGetFileActivityName = "bitbucket.source.getFile"
)

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-source/#api-repositories-workspace-repo-slug-src-commit-path-get
type SourceGetRequest struct {
	ThrippyLinkID string `json:"thrippy_link_id,omitempty"`

	Workspace string `json:"workspace"`
	RepoSlug  string `json:"repo_slug"`
	Commit    string `json:"commit"`
	Path      string `json:"path"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#filtering
	Filter string `json:"q,omitempty"`
	Sort   string `json:"sort,omitempty"`

	// https://developer.atlassian.com/cloud/bitbucket/rest/intro/#pagination
	PageLen  string `json:"pagelen,omitempty"`
	Page     string `json:"page,omitempty"`
	AllPages bool   `json:"all_pages,omitempty"`
}

// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-source/#api-repositories-workspace-repo-slug-src-commit-path-get
func SourceGetFileActivity(ctx workflow.Context, req SourceGetRequest) (string, error) {
	fut := internal.ExecuteTimpaniActivity(ctx, SourceGetFileActivityName, req)

	resp := new(string)
	if err := fut.Get(ctx, resp); err != nil {
		return "", err
	}

	return *resp, nil
}
