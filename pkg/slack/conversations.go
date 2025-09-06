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

// https://docs.slack.dev/reference/methods/conversations.archive/
type ConversationsArchiveRequest struct {
	Channel string `json:"channel"`
}

// https://docs.slack.dev/reference/methods/conversations.archive/
type ConversationsArchiveResponse struct {
	Response
}

// https://docs.slack.dev/reference/methods/conversations.archive/
func ConversationsArchiveActivity(ctx workflow.Context, channelID string) error {
	req := ConversationsArchiveRequest{Channel: channelID}
	return internal.ExecuteTimpaniActivity(ctx, ConversationsArchiveActivityName, req).Get(ctx, nil)
}

// https://docs.slack.dev/reference/methods/conversations.close/
type ConversationsCloseRequest struct {
	Channel string `json:"channel"`
}

// https://docs.slack.dev/reference/methods/conversations.close/
type ConversationsCloseResponse struct {
	Response

	NoOp          bool `json:"no_op,omitempty"`
	AlreadyClosed bool `json:"already_closed,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.create/
type ConversationsCreateRequest struct {
	Name string `json:"name"`

	IsPrivate bool   `json:"is_private,omitempty"`
	TeamID    string `json:"team_id,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.create/
type ConversationsCreateResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.create/
func ConversationsCreateActivity(ctx workflow.Context, name string, private bool) (string, error) {
	req := ConversationsCreateRequest{Name: name, IsPrivate: private}
	fut := internal.ExecuteTimpaniActivity(ctx, ConversationsCreateActivityName, req)

	resp := &ConversationsCreateResponse{}
	if err := fut.Get(ctx, resp); err != nil {
		return "", err
	}

	// The key is guaranteed to exist, and the value is guaranteed to be a string.
	return resp.Channel["id"].(string), nil
}

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

// https://docs.slack.dev/reference/methods/conversations.info/
type ConversationsInfoRequest struct {
	Channel string `json:"channel"`

	IncludeLocale     bool `json:"include_locale,omitempty"`
	IncludeNumMembers bool `json:"include_num_members,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.info/
type ConversationsInfoResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.invite/
type ConversationsInviteRequest struct {
	Channel string `json:"channel"`
	Users   string `json:"users"`

	Force bool `json:"force,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.invite/
type ConversationsInviteResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any   `json:"channel,omitempty"`
	Errors  []map[string]any `json:"errors,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.invite/
func ConversationsInviteActivity(ctx workflow.Context, channelID string, users []string, force bool) error {
	req := ConversationsInviteRequest{Channel: channelID, Users: strings.Join(users, ","), Force: force}
	return internal.ExecuteTimpaniActivity(ctx, ConversationsInviteActivityName, req).Get(ctx, nil)
}

// https://docs.slack.dev/reference/methods/conversations.join/
type ConversationsJoinRequest struct {
	Channel string `json:"channel"`
}

// https://docs.slack.dev/reference/methods/conversations.join/
type ConversationsJoinResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.kick/
type ConversationsKickRequest struct {
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// https://docs.slack.dev/reference/methods/conversations.kick/
type ConversationsKickResponse struct {
	Response
}

// https://docs.slack.dev/reference/methods/conversations.kick/
func ConversationsKickActivity(ctx workflow.Context, channelID, userID string) error {
	req := ConversationsKickRequest{Channel: channelID, User: userID}
	return internal.ExecuteTimpaniActivity(ctx, ConversationsKickActivityName, req).Get(ctx, nil)
}

// https://docs.slack.dev/reference/methods/conversations.leave/
type ConversationsLeaveRequest struct {
	Channel string `json:"channel"`
}

// https://docs.slack.dev/reference/methods/conversations.leave/
type ConversationsLeaveResponse struct {
	Response

	NotInChannel bool `json:"not_in_channel,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.list/
type ConversationsListRequest struct {
	Types           string `json:"types,omitempty"`
	ExcludeArchived bool   `json:"exclude_archived,omitempty"`

	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`

	TeamID string `json:"team_id,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.list/
type ConversationsListResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channels []map[string]any `json:"channels,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.members/
type ConversationsMembersRequest struct {
	Channel string `json:"channel"`

	Limit  int    `json:"limit,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.members/
type ConversationsMembersResponse struct {
	Response

	Members []string `json:"members,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.open/
type ConversationsOpenRequest struct {
	Channel         string `json:"channel,omitempty"`
	ReturnIM        bool   `json:"return_im,omitempty"`
	Users           string `json:"users,omitempty"`
	PreventCreation bool   `json:"prevent_creation,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.open/
type ConversationsOpenResponse struct {
	Response

	NoOp        bool `json:"no_op,omitempty"`
	AlreadyOpen bool `json:"already_open,omitempty"`
	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.rename/
type ConversationsRenameRequest struct {
	Channel string `json:"channel"`
	Name    string `json:"name"`
}

// https://docs.slack.dev/reference/methods/conversations.rename/
type ConversationsRenameResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"`
}

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

// https://docs.slack.dev/reference/methods/conversations.replies/
type ConversationsRepliesResponse struct {
	Response

	Messages []map[string]any `json:"messages,omitempty"`
	HasMore  bool             `json:"has_more,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.setPurpose/
type ConversationsSetPurposeRequest struct {
	Channel string `json:"channel"`
	Purpose string `json:"purpose"`
}

// https://docs.slack.dev/reference/methods/conversations.setPurpose/
type ConversationsSetPurposeResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"` // Empirically different from the documentation.
}

// https://docs.slack.dev/reference/methods/conversations.setPurpose/
func ConversationsSetPurposeActivity(ctx workflow.Context, channelID, purpose string) error {
	req := ConversationsSetPurposeRequest{Channel: channelID, Purpose: purpose}
	return internal.ExecuteTimpaniActivity(ctx, ConversationsSetPurposeActivityName, req).Get(ctx, nil)
}

// https://docs.slack.dev/reference/methods/conversations.setTopic/
type ConversationsSetTopicRequest struct {
	Channel string `json:"channel"`
	Topic   string `json:"topic"`
}

// https://docs.slack.dev/reference/methods/conversations.setTopic/
type ConversationsSetTopicResponse struct {
	Response

	// https://docs.slack.dev/reference/objects/conversation-object/
	Channel map[string]any `json:"channel,omitempty"`
}

// https://docs.slack.dev/reference/methods/conversations.setTopic/
func ConversationsSetTopicActivity(ctx workflow.Context, channelID, topic string) error {
	req := ConversationsSetTopicRequest{Channel: channelID, Topic: topic}
	return internal.ExecuteTimpaniActivity(ctx, ConversationsSetTopicActivityName, req).Get(ctx, nil)
}
