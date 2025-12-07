package slack

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

//revive:disable:exported
const (
	ReactionsAddActivityName    = "slack.reactions.add"
	ReactionsGetActivityName    = "slack.reactions.get"
	ReactionsListActivityName   = "slack.reactions.list"
	ReactionsRemoveActivityName = "slack.reactions.remove"
) //revive:enable:exported

// ReactionsAddRequest is based on:
// https://docs.slack.dev/reference/methods/reactions.add/
type ReactionsAddRequest struct {
	Channel   string `json:"channel"`
	Timestamp string `json:"timestamp"`
	Name      string `json:"name"`
}

// ReactionsAddResponse is based on:
// https://docs.slack.dev/reference/methods/reactions.add/
type ReactionsAddResponse Response

// ReactionsAdd is based on:
// https://docs.slack.dev/reference/methods/reactions.add/
func ReactionsAdd(ctx workflow.Context, channelID, timestamp, name string) error {
	req := ReactionsAddRequest{Channel: channelID, Timestamp: timestamp, Name: name}
	return internal.ExecuteTimpaniActivityNoResp(ctx, ReactionsAddActivityName, req)
}

// ReactionsGetRequest is based on:
// https://docs.slack.dev/reference/methods/reactions.get/
type ReactionsGetRequest struct {
	Channel     string `json:"channel,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	File        string `json:"file,omitempty"`
	FileComment string `json:"file_comment,omitempty"`
	Full        bool   `json:"full,omitempty"`
}

// ReactionsGetResponse is based on:
// https://docs.slack.dev/reference/methods/reactions.get/
type ReactionsGetResponse struct {
	Response

	Type        string         `json:"type,omitempty"`
	Message     map[string]any `json:"message,omitempty"`
	File        map[string]any `json:"file,omitempty"`
	FileComment map[string]any `json:"file_comment,omitempty"`
	Channel     string         `json:"channel,omitempty"`
}

// ReactionsGet is based on:
// https://docs.slack.dev/reference/methods/reactions.get/
func ReactionsGet(ctx workflow.Context, channelID, timestamp string) (map[string]any, error) {
	req := ReactionsGetRequest{Channel: channelID, Timestamp: timestamp, Full: true}
	resp, err := internal.ExecuteTimpaniActivity[ReactionsGetResponse](ctx, ReactionsGetActivityName, req)
	if err != nil {
		return nil, err
	}
	return resp.Message, nil
}

// ReactionsListRequest is based on:
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

// ReactionsListResponse is based on:
// https://docs.slack.dev/reference/methods/reactions.list/
type ReactionsListResponse struct {
	Response

	Items []map[string]any `json:"items,omitempty"`
}

// ReactionsRemoveRequest is based on:
// https://docs.slack.dev/reference/methods/reactions.remove/
type ReactionsRemoveRequest struct {
	Name string `json:"name"`

	Channel     string `json:"channel,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	File        string `json:"file,omitempty"`
	FileComment string `json:"file_comment,omitempty"`
}

// ReactionsRemoveResponse is based on:
// https://docs.slack.dev/reference/methods/reactions.remove/
type ReactionsRemoveResponse Response

// ReactionsRemove is based on:
// https://docs.slack.dev/reference/methods/reactions.remove/
func ReactionsRemove(ctx workflow.Context, channelID, timestamp, name string) error {
	req := ReactionsRemoveRequest{Channel: channelID, Timestamp: timestamp, Name: name}
	return internal.ExecuteTimpaniActivityNoResp(ctx, ReactionsRemoveActivityName, req)
}
