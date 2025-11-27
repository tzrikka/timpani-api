package slack

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	FilesCompleteUploadExternalActivityName = "slack.files.completeUploadExternal"
	FilesGetUploadURLExternalActivityName   = "slack.files.getUploadURLExternal"

	TimpaniUploadExternalActivityName = "slack.timpani.uploadExternal"
)

// FilesGetUploadURLExternalRequest is based on:
// https://docs.slack.dev/reference/methods/files.getuploadurlexternal/
type FilesGetUploadURLExternalRequest struct {
	Length   int    `json:"length"`
	Filename string `json:"filename"`

	SnippetType string `json:"snippet_type,omitempty"`
	AltTxt      string `json:"alt_txt,omitempty"`
}

// FilesGetUploadURLExternalResponse is based on:
// https://docs.slack.dev/reference/methods/files.getuploadurlexternal/
type FilesGetUploadURLExternalResponse struct {
	Response

	UploadURL string `json:"upload_url,omitempty"`
	FileID    string `json:"file_id,omitempty"`
}

// FilesGetUploadURLExternal is based on:
// https://docs.slack.dev/reference/methods/files.getuploadurlexternal/
func FilesGetUploadURLExternal(ctx workflow.Context, length int, filename, snippetType, altTxt string) (string, string, error) {
	req := FilesGetUploadURLExternalRequest{Length: length, Filename: filename, SnippetType: snippetType, AltTxt: altTxt}
	resp, err := internal.ExecuteTimpaniActivity[FilesGetUploadURLExternalResponse](ctx, FilesGetUploadURLExternalActivityName, req)
	if err != nil {
		return "", "", err
	}
	return resp.UploadURL, resp.FileID, nil
}

// FilesCompleteUploadExternalRequest is based on:
// https://docs.slack.dev/reference/methods/files.completeuploadexternal/
type FilesCompleteUploadExternalRequest struct {
	Files []File `json:"files"`

	ChannelID      string `json:"channel_id,omitempty"`
	ThreadTS       string `json:"thread_ts,omitempty"`
	Channels       string `json:"channels,omitempty"`
	InitialComment string `json:"initial_comment,omitempty"`
	Blocks         string `json:"blocks,omitempty"`
}

// FilesCompleteUploadExternalResponse is based on:
// https://docs.slack.dev/reference/methods/files.completeuploadexternal/
type FilesCompleteUploadExternalResponse struct {
	Response

	Files []File `json:"files,omitempty"`
}

// FilesCompleteUploadExternal is based on:
// https://docs.slack.dev/reference/methods/files.completeuploadexternal/
func FilesCompleteUploadExternal(ctx workflow.Context, req FilesCompleteUploadExternalRequest) ([]File, error) {
	resp, err := internal.ExecuteTimpaniActivity[FilesCompleteUploadExternalResponse](ctx, FilesCompleteUploadExternalActivityName, req)
	if err != nil {
		return nil, err
	}
	return resp.Files, nil
}

// TimpaniUploadExternalRequest is based on:
// https://docs.slack.dev/messaging/working-with-files/
type TimpaniUploadExternalRequest struct {
	URL      string `json:"url"`
	MimeType string `json:"mime_type"`
	Content  []byte `json:"content"`
}

// TimpaniUploadExternal is based on:
// https://docs.slack.dev/messaging/working-with-files/
func TimpaniUploadExternal(ctx workflow.Context, url, mimeType string, content []byte) error {
	req := TimpaniUploadExternalRequest{URL: url, MimeType: mimeType, Content: content}
	return internal.ExecuteTimpaniActivityNoResp(ctx, TimpaniUploadExternalActivityName, req)
}

// File is based on:
// https://docs.slack.dev/reference/objects/file-object/#types
type File struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`

	Created      int    `json:"created,omitempty"`
	Updated      int    `json:"updated,omitempty"`
	Size         int    `json:"size,omitempty"`
	Name         string `json:"name,omitempty"`
	FileType     string `json:"filetype,omitempty"`
	MimeType     string `json:"mimetype,omitempty"`
	PrettyType   string `json:"pretty_type,omitempty"`
	ExternalType string `json:"external_type,omitempty"`

	UserTeam string `json:"user_team,omitempty"`
	User     string `json:"user,omitempty"`

	Channels []string `json:"channels,omitempty"`
	Groups   []string `json:"groups,omitempty"`
	IMs      []string `json:"ims,omitempty"`
	PinnedTo []string `json:"pinned_to,omitempty"`

	Mode       string `json:"mode,omitempty"`
	FileAccess string `json:"file_access,omitempty"`
	Lines      int    `json:"lines,omitempty"`
	LinesMore  int    `json:"lines_more,omitempty"`

	Editable           bool `json:"editable,omitempty"`
	HasMoreShares      bool `json:"has_more_shares,omitempty"`
	HasRichPreview     bool `json:"has_rich_preview,omitempty"`
	IsExternal         bool `json:"is_external,omitempty"`
	IsPublic           bool `json:"is_public,omitempty"`
	PreviewIsTruncated bool `json:"preview_is_truncated,omitempty"`
	PublicURLShared    bool `json:"public_url_shared,omitempty"`

	EditLink           string `json:"edit_link,omitempty"`
	Permalink          string `json:"permalink,omitempty"`
	PermalinkPublic    string `json:"permalink_public,omitempty"`
	URLPrivate         string `json:"url_private,omitempty"`
	URLPrivateDownload string `json:"url_private_download,omitempty"`

	Preview          string `json:"preview,omitempty"`
	PreviewHighlight string `json:"preview_highlight,omitempty"`

	Shares map[string]map[string][]FileShare `json:"shares,omitempty"` // "private"/"public" -> channel ID -> []FileShare.
}

// FileShare is based on:
// https://docs.slack.dev/reference/objects/file-object/#types
type FileShare struct {
	ChannelName string `json:"channel_name,omitempty"`
	ShareUserID string `json:"share_user_id,omitempty"`
	TeamID      string `json:"team_id,omitempty"`
	TS          string `json:"ts,omitempty"`

	IsSilentShare bool   `json:"is_silent_share,omitempty"`
	Source        string `json:"source,omitempty"`
}
