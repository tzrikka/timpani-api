package slack

import (
	"strings"

	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	ConversationsArchiveActivityName    = "slack.conversations.archive"
	ConversationsCloseActivityName      = "slack.conversations.close"
	ConversationsCreateActivityName     = "slack.conversations.create"
	ConversationsHistoryActivityName    = "slack.conversations.history"
	ConversationsInfoActivityName       = "slack.conversations.info"
	ConversationsInviteActivityName     = "slack.conversations.invite"
	ConversationsJoinActivityName       = "slack.conversations.join"
	ConversationsKickActivityName       = "slack.conversations.kick"
	ConversationsLeaveActivityName      = "slack.conversations.leave"
	ConversationsListActivityName       = "slack.conversations.list"
	ConversationsMembersActivityName    = "slack.conversations.members"
	ConversationsOpenActivityName       = "slack.conversations.open"
	ConversationsRenameActivityName     = "slack.conversations.rename"
	ConversationsRepliesActivityName    = "slack.conversations.replies"
	ConversationsSetPurposeActivityName = "slack.conversations.setPurpose"
	ConversationsSetTopicActivityName   = "slack.conversations.setTopic"
)

// ConversationsArchiveRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.archive/
type ConversationsArchiveRequest struct {
	Channel string `json:"channel"`
}

// ConversationsArchiveResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.archive/
type ConversationsArchiveResponse Response

// ConversationsArchive is based on:
// https://docs.slack.dev/reference/methods/conversations.archive/
func ConversationsArchive(ctx workflow.Context, channelID string) error {
	req := ConversationsArchiveRequest{Channel: channelID}
	return internal.ExecuteTimpaniActivityNoResp(ctx, ConversationsArchiveActivityName, req)
}

// ConversationsCloseRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.close/
type ConversationsCloseRequest struct {
	Channel string `json:"channel"`
}

// ConversationsCloseResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.close/
type ConversationsCloseResponse struct {
	Response

	NoOp          bool `json:"no_op,omitempty"`
	AlreadyClosed bool `json:"already_closed,omitempty"`
}

// ConversationsCreateRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.create/
type ConversationsCreateRequest struct {
	Name string `json:"name"`

	IsPrivate bool   `json:"is_private,omitempty"`
	TeamID    string `json:"team_id,omitempty"`
}

// ConversationsCreateResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.create/
type ConversationsCreateResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"`
}

// ConversationsCreate is based on:
// https://docs.slack.dev/reference/methods/conversations.create/
func ConversationsCreate(ctx workflow.Context, name string, private bool) (string, error) {
	req := ConversationsCreateRequest{Name: name, IsPrivate: private}
	resp, err := internal.ExecuteTimpaniActivity[ConversationsCreateResponse](ctx, ConversationsCreateActivityName, req)
	if err != nil {
		return "", err
	}

	// The key is guaranteed to exist, and the value is guaranteed to be a string.
	return resp.Channel["id"].(string), nil
}

// ConversationsHistoryRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.history/
type ConversationsHistoryRequest struct {
	Channel string `json:"channel"`

	IncludeAllMetadata bool   `json:"include_all_metadata,omitempty"`
	Inclusive          bool   `json:"inclusive,omitempty"`
	Latest             string `json:"latest,omitempty"`
	Oldest             string `json:"oldest,omitempty"`

	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

// ConversationsHistoryResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.history/
type ConversationsHistoryResponse struct {
	Response

	Messages  []map[string]any `json:"messages,omitempty"`
	Latest    string           `json:"latest,omitempty"`
	HasMore   bool             `json:"has_more,omitempty"`
	IsLimited bool             `json:"is_limited,omitempty"`
	PinCount  int              `json:"pin_count,omitempty"`

	// Undocumented: "channel_actions_ts" and "channel_actions_count".
}

// ConversationsInfoRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.info/
type ConversationsInfoRequest struct {
	Channel string `json:"channel"`

	IncludeLocale     bool `json:"include_locale,omitempty"`
	IncludeNumMembers bool `json:"include_num_members,omitempty"`
}

// ConversationsInfoResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.info/
type ConversationsInfoResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"`
}

// ConversationsInfo is based on:
// https://docs.slack.dev/reference/methods/conversations.info/
func ConversationsInfo(ctx workflow.Context, channelID string, locale, numMembers bool) (map[string]any, error) {
	req := ConversationsInfoRequest{Channel: channelID, IncludeLocale: locale, IncludeNumMembers: numMembers}
	resp, err := internal.ExecuteTimpaniActivity[ConversationsInfoResponse](ctx, ConversationsInfoActivityName, req)
	if err != nil {
		return nil, err
	}
	return resp.Channel, nil
}

// ConversationsInviteRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.invite/
type ConversationsInviteRequest struct {
	Channel string `json:"channel"`
	Users   string `json:"users"`

	Force bool `json:"force,omitempty"`
}

// ConversationsInviteResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.invite/
type ConversationsInviteResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any   `json:"channel,omitempty"`
	Errors  []map[string]any `json:"errors,omitempty"`
}

// ConversationsInvite is based on:
// https://docs.slack.dev/reference/methods/conversations.invite/
func ConversationsInvite(ctx workflow.Context, channelID string, users []string, force bool) error {
	req := ConversationsInviteRequest{Channel: channelID, Users: strings.Join(users, ","), Force: force}
	return internal.ExecuteTimpaniActivityNoResp(ctx, ConversationsInviteActivityName, req)
}

// ConversationsJoinRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.join/
type ConversationsJoinRequest struct {
	Channel string `json:"channel"`
}

// ConversationsJoinResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.join/
type ConversationsJoinResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"`
}

// ConversationsKickRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.kick/
type ConversationsKickRequest struct {
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// ConversationsKickResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.kick/
type ConversationsKickResponse struct {
	Response
}

// ConversationsKick is based on:
// https://docs.slack.dev/reference/methods/conversations.kick/
func ConversationsKick(ctx workflow.Context, channelID, userID string) error {
	req := ConversationsKickRequest{Channel: channelID, User: userID}
	return internal.ExecuteTimpaniActivityNoResp(ctx, ConversationsKickActivityName, req)
}

// ConversationsLeaveRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.leave/
type ConversationsLeaveRequest struct {
	Channel string `json:"channel"`
}

// ConversationsLeaveResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.leave/
type ConversationsLeaveResponse struct {
	Response

	NotInChannel bool `json:"not_in_channel,omitempty"`
}

// ConversationsListRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.list/
type ConversationsListRequest struct {
	Types           string `json:"types,omitempty"`
	ExcludeArchived bool   `json:"exclude_archived,omitempty"`

	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`

	TeamID string `json:"team_id,omitempty"`
}

// ConversationsListResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.list/
type ConversationsListResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channels []map[string]any `json:"channels,omitempty"`
}

// ConversationsMembersRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.members/
type ConversationsMembersRequest struct {
	Channel string `json:"channel"`

	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

// ConversationsMembersResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.members/
type ConversationsMembersResponse struct {
	Response

	Members []string `json:"members,omitempty"`
}

// ConversationsOpenRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.open/
type ConversationsOpenRequest struct {
	Channel         string `json:"channel,omitempty"`
	ReturnIM        bool   `json:"return_im,omitempty"`
	Users           string `json:"users,omitempty"`
	PreventCreation bool   `json:"prevent_creation,omitempty"`
}

// ConversationsOpenResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.open/
type ConversationsOpenResponse struct {
	Response

	NoOp        bool `json:"no_op,omitempty"`
	AlreadyOpen bool `json:"already_open,omitempty"`
	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"`
}

// ConversationsRenameRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.rename/
type ConversationsRenameRequest struct {
	Channel string `json:"channel"`
	Name    string `json:"name"`
}

// ConversationsRenameResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.rename/
type ConversationsRenameResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"`
}

// ConversationsRename is based on:
// https://docs.slack.dev/reference/methods/conversations.rename/
func ConversationsRename(ctx workflow.Context, channelID, name string) error {
	req := ConversationsRenameRequest{Channel: channelID, Name: name}
	return internal.ExecuteTimpaniActivityNoResp(ctx, ConversationsRenameActivityName, req)
}

// ConversationsRepliesRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.replies/
type ConversationsRepliesRequest struct {
	Channel string `json:"channel"`
	TS      string `json:"ts"`

	IncludeAllMetadata bool   `json:"include_all_metadata,omitempty"`
	Inclusive          bool   `json:"inclusive,omitempty"`
	Latest             string `json:"latest,omitempty"`
	Oldest             string `json:"oldest,omitempty"`

	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

// ConversationsRepliesResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.replies/
type ConversationsRepliesResponse struct {
	Response

	Messages []map[string]any `json:"messages,omitempty"`
	HasMore  bool             `json:"has_more,omitempty"`
}

// ConversationsSetPurposeRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.setPurpose/
type ConversationsSetPurposeRequest struct {
	Channel string `json:"channel"`
	Purpose string `json:"purpose"`
}

// ConversationsSetPurposeResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.setPurpose/
type ConversationsSetPurposeResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"` // Empirically different from the documentation.
}

// ConversationsSetPurpose is based on:
// https://docs.slack.dev/reference/methods/conversations.setPurpose/
func ConversationsSetPurpose(ctx workflow.Context, channelID, purpose string) error {
	req := ConversationsSetPurposeRequest{Channel: channelID, Purpose: purpose}
	return internal.ExecuteTimpaniActivityNoResp(ctx, ConversationsSetPurposeActivityName, req)
}

// ConversationsSetTopicRequest is based on:
// https://docs.slack.dev/reference/methods/conversations.setTopic/
type ConversationsSetTopicRequest struct {
	Channel string `json:"channel"`
	Topic   string `json:"topic"`
}

// ConversationsSetTopicResponse is based on:
// https://docs.slack.dev/reference/methods/conversations.setTopic/
type ConversationsSetTopicResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"`
}

// ConversationsSetTopic is based on:
// https://docs.slack.dev/reference/methods/conversations.setTopic/
func ConversationsSetTopic(ctx workflow.Context, channelID, topic string) error {
	req := ConversationsSetTopicRequest{Channel: channelID, Topic: topic}
	return internal.ExecuteTimpaniActivityNoResp(ctx, ConversationsSetTopicActivityName, req)
}
