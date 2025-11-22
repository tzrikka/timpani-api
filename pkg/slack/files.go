package slack

import (
	"strconv"

	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	FilesGetUploadURLExternalActivityName   = "slack.files.getUploadURLExternal"
	FilesCompleteUploadExternalActivityName = "slack.files.completeUploadExternal"

	TimpaniUploadExternalActivityName = "slack.timpani.uploadExternal"
)

// https://docs.slack.dev/reference/methods/files.getuploadurlexternal/
type FilesGetUploadURLExternalRequest struct {
	Length   string `json:"length"` // Contrary to Slack's documentation, this has to be a string.
	Filename string `json:"filename"`

	SnippetType string `json:"snippet_type,omitempty"`
	AltTxt      string `json:"alt_txt,omitempty"`
}

// https://docs.slack.dev/reference/methods/files.getuploadurlexternal/
type FilesGetUploadURLExternalResponse struct {
	Response

	UploadURL string `json:"upload_url,omitempty"`
	FileID    string `json:"file_id,omitempty"`
}

// https://docs.slack.dev/reference/methods/files.getuploadurlexternal/
func FilesGetUploadURLExternalActivity(ctx workflow.Context, length int, filename, snippetType, altTxt string) (string, string, error) {
	fut := internal.ExecuteTimpaniActivity(ctx, FilesGetUploadURLExternalActivityName, FilesGetUploadURLExternalRequest{
		Length:      strconv.Itoa(length), // Contrary to Slack's documentation, this has to be a string.
		Filename:    filename,
		SnippetType: snippetType,
		AltTxt:      altTxt,
	})

	resp := new(FilesGetUploadURLExternalResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return "", "", err
	}

	return resp.UploadURL, resp.FileID, nil
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

	Files []map[string]any `json:"files,omitempty"`
}

// https://docs.slack.dev/reference/methods/files.completeuploadexternal/
func FilesCompleteUploadExternalActivity(ctx workflow.Context, fileID, title, channelID, threadTS string) error {
	req := FilesCompleteUploadExternalRequest{Files: []File{{ID: fileID, Title: title}}, ChannelID: channelID, ThreadTS: threadTS}
	return internal.ExecuteTimpaniActivity(ctx, FilesCompleteUploadExternalActivityName, req).Get(ctx, nil)
}

type TimpaniUploadExternalRequest struct {
	URL      string `json:"url"`
	MimeType string `json:"mime_type"`
	Content  []byte `json:"content"`
}

func TimpaniUploadExternalActivity(ctx workflow.Context, url, mimeType string, content []byte) error {
	req := TimpaniUploadExternalRequest{URL: url, MimeType: mimeType, Content: content}
	return internal.ExecuteTimpaniActivity(ctx, TimpaniUploadExternalActivityName, req).Get(ctx, nil)
}

// https://docs.slack.dev/reference/methods/files.completeuploadexternal/
type File struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}
