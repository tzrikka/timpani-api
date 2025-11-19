package slack

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	FilesGetUploadURLExternalActivityName   = "slack.files.getUploadURLExternal"
	FilesCompleteUploadExternalActivityName = "slack.files.completeUploadExternal"
)

// https://docs.slack.dev/reference/methods/files.getuploadurlexternal/
type FilesGetUploadURLExternalRequest struct {
	Length   int    `json:"length"`
	Filename string `json:"filename"`

	SnippetType string `json:"snippet_type,omitempty"`
	AltText     string `json:"alt_text,omitempty"`
}

// https://docs.slack.dev/reference/methods/files.getuploadurlexternal/
type FilesGetUploadURLExternalResponse struct {
	Response

	UploadURL string `json:"upload_url,omitempty"`
	FileID    string `json:"file_id,omitempty"`
}

// https://docs.slack.dev/reference/methods/files.getuploadurlexternal/
func FilesGetUploadURLExternalActivity(ctx workflow.Context, content []byte, name, snippetType, altText string) (string, error) {
	req := FilesGetUploadURLExternalRequest{Length: len(content), Filename: name, SnippetType: snippetType, AltText: altText}
	fut := internal.ExecuteTimpaniActivity(ctx, FilesGetUploadURLExternalActivityName, req)

	resp := new(FilesGetUploadURLExternalResponse)
	if err := fut.Get(ctx, resp); err != nil { // Timpani also uploads [content] to [resp.UploadURL].
		return "", err
	}

	return resp.FileID, nil
}

// https://docs.slack.dev/reference/methods/files.completeuploadexternal/
type FilesCompleteUploadExternalRequest struct {
	Files []File `json:"files"`

	ChannelID      string `json:"channel_id,omitempty"`
	ThreadTS       string `json:"thread_ts,omitempty"`
	Channels       string `json:"channels,omitempty"`
	InitialComment string `json:"initial_comment,omitempty"`
	Blocks         string `json:"blocks,omitempty"`
}

// https://docs.slack.dev/reference/methods/files.completeuploadexternal/
type FilesCompleteUploadExternalResponse struct {
	Response

	Files []File `json:"files,omitempty"`
}

// https://docs.slack.dev/reference/methods/files.completeuploadexternal/
func FilesCompleteUploadExternalActivity(ctx workflow.Context, fileID, title, channelID, threadTS string) error {
	req := FilesCompleteUploadExternalRequest{
		Files:     []File{{ID: fileID, Title: title}},
		ChannelID: channelID,
		ThreadTS:  threadTS,
	}
	fut := internal.ExecuteTimpaniActivity(ctx, FilesCompleteUploadExternalActivityName, req)

	resp := new(FilesCompleteUploadExternalResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return err
	}

	return nil
}

// https://docs.slack.dev/reference/methods/files.completeuploadexternal/
type File struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}
