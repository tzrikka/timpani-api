package bitbucket

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

//revive:disable:exported
const (
	SourceGetFileActivityName = "bitbucket.source.getFile"
) //revive:enable:exported

// SourceGetRequest is based on:
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
	PageLen string `json:"pagelen,omitempty"`
	Page    string `json:"page,omitempty"`

	Next string `json:"next,omitempty"`
}

// SourceGetFile is based on:
// https://developer.atlassian.com/cloud/bitbucket/rest/api-group-source/#api-repositories-workspace-repo-slug-src-commit-path-get
func SourceGetFile(ctx workflow.Context, thrippyLinkID, workspace, repo, commit, path string) (string, error) {
	resp, err := internal.ExecuteTimpaniActivity[string](ctx, SourceGetFileActivityName, SourceGetRequest{
		ThrippyLinkID: thrippyLinkID,
		Workspace:     workspace,
		RepoSlug:      repo,
		Commit:        commit,
		Path:          path,
	})
	if err != nil {
		return "", err
	}
	return *resp, nil
}
