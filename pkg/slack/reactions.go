package slack

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	ReactionsAddActivityName    = "slack.reactions.add"
	ReactionsGetActivityName    = "slack.reactions.get"
	ReactionsListActivityName   = "slack.reactions.list"
	ReactionsRemoveActivityName = "slack.reactions.remove"
)

// https://docs.slack.dev/reference/methods/reactions.add/
type ReactionsAddRequest struct {
	Channel   string `json:"channel"`
	Timestamp string `json:"timestamp"`
	Name      string `json:"name"`
}

// https://docs.slack.dev/reference/methods/reactions.add/
type ReactionsAddResponse struct {
	Response
}

// https://docs.slack.dev/reference/methods/reactions.add/
func ReactionsAddActivity(ctx workflow.Context, channelID, timestamp, name string) error {
	req := ReactionsAddRequest{Channel: channelID, Timestamp: timestamp, Name: name}
	return internal.ExecuteTimpaniActivity(ctx, ReactionsAddActivityName, req).Get(ctx, nil)
}

// https://docs.slack.dev/reference/methods/reactions.get/
type ReactionsGetRequest struct {
	Channel     string `json:"channel,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	File        string `json:"file,omitempty"`
	FileComment string `json:"file_comment,omitempty"`
	Full        bool   `json:"full,omitempty"`
}

// https://docs.slack.dev/reference/methods/reactions.get/
type ReactionsGetResponse struct {
	Response

	Type        string         `json:"type,omitempty"`
	Message     map[string]any `json:"message,omitempty"`
	File        map[string]any `json:"file,omitempty"`
	FileComment map[string]any `json:"file_comment,omitempty"`
	Channel     string         `json:"channel,omitempty"`
}

// https://docs.slack.dev/reference/methods/reactions.get/
func ReactionsGetActivity(ctx workflow.Context, channelID, timestamp string) (map[string]any, error) {
	req := ReactionsGetRequest{Channel: channelID, Timestamp: timestamp, Full: true}
	fut := internal.ExecuteTimpaniActivity(ctx, ReactionsGetActivityName, req)

	resp := new(ReactionsGetResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp.Message, nil
}

// https://docs.slack.dev/reference/methods/reactions.list/
type ReactionsListRequest struct {
	User string `json:"user,omitempty"`

	Full   bool   `json:"full,omitempty"`
	Count  int    `json:"count,omitempty"`
	Page   int    `json:"page,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`

	TeamID string `json:"team_id,omitempty"`
}

// https://docs.slack.dev/reference/methods/reactions.list/
type ReactionsListResponse struct {
	Response

	Items []map[string]any `json:"items,omitempty"`
}

// https://docs.slack.dev/reference/methods/reactions.remove/
type ReactionsRemoveRequest struct {
	Name string `json:"name"`

	Channel     string `json:"channel,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	File        string `json:"file,omitempty"`
	FileComment string `json:"file_comment,omitempty"`
}

// https://docs.slack.dev/reference/methods/reactions.remove/
type ReactionsRemoveResponse Response

// https://docs.slack.dev/reference/methods/reactions.remove/
func ReactionsRemoveActivity(ctx workflow.Context, channelID, timestamp, name string) error {
	req := ReactionsRemoveRequest{Channel: channelID, Timestamp: timestamp, Name: name}
	return internal.ExecuteTimpaniActivity(ctx, ReactionsRemoveActivityName, req).Get(ctx, nil)
}
