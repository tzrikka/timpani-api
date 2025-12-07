package slack

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

//revive:disable:exported
const (
	ChatDeleteActivityName        = "slack.chat.delete"
	ChatGetPermalinkActivityName  = "slack.chat.getPermalink"
	ChatPostEphemeralActivityName = "slack.chat.postEphemeral"
	ChatPostMessageActivityName   = "slack.chat.postMessage"
	ChatUpdateActivityName        = "slack.chat.update"

	TimpaniPostApprovalWorkflowName = "slack.timpani.postApproval"
) //revive:enable:exported

// ChatDeleteRequest is based on:
// https://docs.slack.dev/reference/methods/chat.delete/
type ChatDeleteRequest struct {
	Channel string `json:"channel"`
	TS      string `json:"ts"`

	AsUser bool `json:"as_user,omitempty"`
}

// ChatDeleteResponse is based on:
// https://docs.slack.dev/reference/methods/chat.delete/
type ChatDeleteResponse struct {
	Response

	Channel string `json:"channel,omitempty"`
	TS      string `json:"ts,omitempty"`
}

// ChatDelete is based on:
// https://docs.slack.dev/reference/methods/chat.delete/
func ChatDelete(ctx workflow.Context, channelID, timestamp string) error {
	req := ChatDeleteRequest{Channel: channelID, TS: timestamp}
	return internal.ExecuteTimpaniActivityNoResp(ctx, ChatDeleteActivityName, req)
}

// ChatGetPermalinkRequest is based on:
// https://docs.slack.dev/reference/methods/chat.getPermalink/
type ChatGetPermalinkRequest struct {
	Channel   string `json:"channel"`
	MessageTS string `json:"message_ts"`
}

// ChatGetPermalinkResponse is based on:
// https://docs.slack.dev/reference/methods/chat.getPermalink/
type ChatGetPermalinkResponse struct {
	Response

	Channel   string `json:"channel,omitempty"`
	Permalink string `json:"permalink,omitempty"`
}

// ChatGetPermalink is based on:
// https://docs.slack.dev/reference/methods/chat.getPermalink/
func ChatGetPermalink(ctx workflow.Context, channelID, timestamp string) (string, error) {
	req := ChatGetPermalinkRequest{Channel: channelID, MessageTS: timestamp}
	resp, err := internal.ExecuteTimpaniActivity[ChatGetPermalinkResponse](ctx, ChatGetPermalinkActivityName, req)
	if err != nil {
		return "", err
	}
	return resp.Permalink, nil
}

// ChatPostEphemeralRequest is based on:
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

// ChatPostEphemeralResponse is based on:
// https://docs.slack.dev/reference/methods/chat.postEphemeral/
type ChatPostEphemeralResponse struct {
	Response

	MessageTS string `json:"message_ts,omitempty"`
}

// ChatPostEphemeral is based on:
// https://docs.slack.dev/reference/methods/chat.postEphemeral/
func ChatPostEphemeral(ctx workflow.Context, req ChatPostEphemeralRequest) error {
	return internal.ExecuteTimpaniActivityNoResp(ctx, ChatPostEphemeralActivityName, req)
}

// ChatPostMessageRequest is based on:
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

// ChatPostMessageResponse is based on:
// https://docs.slack.dev/reference/methods/chat.postMessage/
type ChatPostMessageResponse struct {
	Response

	Channel string         `json:"channel,omitempty"`
	TS      string         `json:"ts,omitempty"`
	Message map[string]any `json:"message,omitempty"`
}

// ChatPostMessage is based on:
// https://docs.slack.dev/reference/methods/chat.postMessage/
func ChatPostMessage(ctx workflow.Context, req ChatPostMessageRequest) (*ChatPostMessageResponse, error) {
	return internal.ExecuteTimpaniActivity[ChatPostMessageResponse](ctx, ChatPostMessageActivityName, req)
}

// ChatUpdateRequest is based on:
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

// ChatUpdateResponse is based on:
// https://docs.slack.dev/reference/methods/chat.update/
type ChatUpdateResponse struct {
	Response

	Channel string         `json:"channel,omitempty"`
	TS      string         `json:"ts,omitempty"`
	Text    string         `json:"text,omitempty"`
	Message map[string]any `json:"message,omitempty"`
}

// ChatUpdate is based on:
// https://docs.slack.dev/reference/methods/chat.update/
func ChatUpdate(ctx workflow.Context, req ChatUpdateRequest) error {
	return internal.ExecuteTimpaniActivityNoResp(ctx, ChatUpdateActivityName, req)
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

// TimpaniPostApprovalResponse is based on:
// https://docs.slack.dev/reference/interaction-payloads/
type TimpaniPostApprovalResponse struct {
	Response

	InteractionEvent map[string]any `json:"interaction_event,omitempty"`
}

// TimpaniPostApprovalWorkflow is a convenience wrapper over [ChatPostMessage].
// It sends an interactive message to a user/group/channel with a short header,
// a markdown message, and 2 buttons. It then waits for (and returns) the user selection.
//
// For message formatting tips, see https://docs.slack.dev/messaging/formatting-message-text.
func TimpaniPostApprovalWorkflow(ctx workflow.Context, req TimpaniPostApprovalRequest) (map[string]any, error) {
	resp := new(TimpaniPostApprovalResponse)
	fut := workflow.ExecuteChildWorkflow(ctx, TimpaniPostApprovalWorkflowName, req)

	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp.InteractionEvent, nil
}
