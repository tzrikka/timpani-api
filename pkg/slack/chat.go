package slack

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	ChatDeleteActivityName        = "slack.chat.delete"
	ChatGetPermalinkActivityName  = "slack.chat.getPermalink"
	ChatPostEphemeralActivityName = "slack.chat.postEphemeral"
	ChatPostMessageActivityName   = "slack.chat.postMessage"
	ChatUpdateActivityName        = "slack.chat.update"

	TimpaniPostApprovalWorkflowName = "slack.timpani.postApproval"
)

// https://docs.slack.dev/reference/methods/chat.delete/
type ChatDeleteRequest struct {
	Channel string `json:"channel"`
	TS      string `json:"ts"`

	AsUser bool `json:"as_user,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.delete/
type ChatDeleteResponse struct {
	Response

	Channel string `json:"channel,omitempty"`
	TS      string `json:"ts,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.delete/
func ChatDeleteActivity(ctx workflow.Context, channelID, timestamp string) error {
	req := ChatDeleteRequest{Channel: channelID, TS: timestamp}
	return internal.ExecuteTimpaniActivity(ctx, ChatDeleteActivityName, req).Get(ctx, nil)
}

// https://docs.slack.dev/reference/methods/chat.getPermalink/
type ChatGetPermalinkRequest struct {
	Channel   string `json:"channel"`
	MessageTS string `json:"message_ts"`
}

// https://docs.slack.dev/reference/methods/chat.getPermalink/
type ChatGetPermalinkResponse struct {
	Response

	Channel   string `json:"channel,omitempty"`
	Permalink string `json:"permalink,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.getPermalink/
func ChatGetPermalinkActivity(ctx workflow.Context, channelID, timestamp string) (string, error) {
	req := ChatGetPermalinkRequest{Channel: channelID, MessageTS: timestamp}
	fut := internal.ExecuteTimpaniActivity(ctx, ChatGetPermalinkActivityName, req)

	resp := new(ChatGetPermalinkResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return "", err
	}

	return resp.Permalink, nil
}

// https://docs.slack.dev/reference/methods/chat.postEphemeral/
type ChatPostEphemeralRequest struct {
	Channel string `json:"channel"`
	User    string `json:"user"`

	Blocks       []map[string]any `json:"blocks,omitempty"`
	Attachments  []map[string]any `json:"attachments,omitempty"`
	MarkdownText string           `json:"markdown_text,omitempty"`
	Text         string           `json:"text,omitempty"`

	ThreadTS string `json:"thread_ts,omitempty"`

	IconEmoji string `json:"icon_emoji,omitempty"`
	IconURL   string `json:"icon_url,omitempty"`

	LinkNames bool   `json:"link_names,omitempty"`
	Parse     string `json:"parse,omitempty"`
	Username  string `json:"username,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.postEphemeral/
type ChatPostEphemeralResponse struct {
	Response

	MessageTS string `json:"message_ts,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.postEphemeral/
func ChatPostEphemeralActivity(ctx workflow.Context, req ChatPostEphemeralRequest) error {
	return internal.ExecuteTimpaniActivity(ctx, ChatPostEphemeralActivityName, req).Get(ctx, nil)
}

// https://docs.slack.dev/reference/methods/chat.postMessage/
type ChatPostMessageRequest struct {
	Channel string `json:"channel"`

	Blocks       []map[string]any `json:"blocks,omitempty"`
	Attachments  []map[string]any `json:"attachments,omitempty"`
	MarkdownText string           `json:"markdown_text,omitempty"`
	Text         string           `json:"text,omitempty"`

	ThreadTS       string `json:"thread_ts,omitempty"`
	ReplyBroadcast bool   `json:"reply_broadcast,omitempty"`

	IconEmoji string `json:"icon_emoji,omitempty"`
	IconURL   string `json:"icon_url,omitempty"`
	Username  string `json:"username,omitempty"`

	Metadata map[string]any `json:"metadata,omitempty"`

	LinkNames bool `json:"link_names,omitempty"`
	// Ignoring "mrkdwn" for now, because it has an unusual default value (true).
	Parse       string `json:"parse,omitempty"`
	UnfurlLinks bool   `json:"unfurl_links,omitempty"`
	UnfurlMedia bool   `json:"unfurl_media,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.postMessage/
type ChatPostMessageResponse struct {
	Response

	Channel string         `json:"channel,omitempty"`
	TS      string         `json:"ts,omitempty"`
	Message map[string]any `json:"message,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.postMessage/
func ChatPostMessageActivity(ctx workflow.Context, req ChatPostMessageRequest) (*ChatPostMessageResponse, error) {
	fut := internal.ExecuteTimpaniActivity(ctx, ChatPostMessageActivityName, req)

	resp := new(ChatPostMessageResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// https://docs.slack.dev/reference/methods/chat.update/
type ChatUpdateRequest struct {
	Channel string `json:"channel"`
	TS      string `json:"ts"`

	Blocks       []map[string]any `json:"blocks,omitempty"`
	Attachments  []map[string]any `json:"attachments,omitempty"`
	MarkdownText string           `json:"markdown_text,omitempty"`
	Text         string           `json:"text,omitempty"`

	FileIDs        []string       `json:"file_ids,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	ReplyBroadcast bool           `json:"reply_broadcast,omitempty"`

	AsUser    bool   `json:"as_user,omitempty"`
	LinkNames bool   `json:"link_names,omitempty"`
	Parse     string `json:"parse,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.update/
type ChatUpdateResponse struct {
	Response

	Channel string         `json:"channel,omitempty"`
	TS      string         `json:"ts,omitempty"`
	Text    string         `json:"text,omitempty"`
	Message map[string]any `json:"message,omitempty"`
}

// https://docs.slack.dev/reference/methods/chat.update/
func ChatUpdateActivity(ctx workflow.Context, req ChatUpdateRequest) error {
	return internal.ExecuteTimpaniActivity(ctx, ChatUpdateActivityName, req).Get(ctx, nil)
}

// TimpaniPostApprovalRequest is similar to [ChatPostMessageRequest]. If button
// labels are not specified here, their default values are "Approve" and "Deny".
type TimpaniPostApprovalRequest struct {
	Channel string `json:"channel"`

	Header      string `json:"header"`
	Message     string `json:"message"`
	GreenButton string `json:"green_button,omitempty"`
	RedButton   string `json:"red_button,omitempty"`

	ThreadTS       string         `json:"thread_ts,omitempty"`
	ReplyBroadcast bool           `json:"reply_broadcast,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`

	Timeout string `json:"timeout,omitempty"`
}

type TimpaniPostApprovalResponse struct {
	Response

	InteractionEvent map[string]any `json:"interaction_event,omitempty"`
}

// TimpaniPostApprovalWorkflow is a convenience wrapper over
// [ChatPostMessageActivity]. It sends an interactive message to a
// user/group/channel with a short header, a markdown message, and
// 2 buttons. It then waits for (and returns) the user selection.
//
// For message formatting tips, see
// https://docs.slack.dev/messaging/formatting-message-text.
func TimpaniPostApprovalWorkflow(ctx workflow.Context, req TimpaniPostApprovalRequest) (map[string]any, error) {
	resp := new(TimpaniPostApprovalResponse)
	fut := workflow.ExecuteChildWorkflow(ctx, TimpaniPostApprovalWorkflowName, req)

	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp.InteractionEvent, nil
}
