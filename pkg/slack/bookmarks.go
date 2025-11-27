package slack

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	BookmarksAddActivityName    = "slack.bookmarks.add"
	BookmarksEditActivityName   = "slack.bookmarks.edit"
	BookmarksListActivityName   = "slack.bookmarks.list"
	BookmarksRemoveActivityName = "slack.bookmarks.remove"
)

// BookmarksAddRequest is based on:
// https://docs.slack.dev/reference/methods/bookmarks.add/
type BookmarksAddRequest struct {
	ChannelID string `json:"channel_id"`
	Title     string `json:"title"`
	Type      string `json:"type"`

	Link        string `json:"link,omitempty"`
	Emoji       string `json:"emoji,omitempty"`
	EntityID    string `json:"entity_id,omitempty"`
	AccessLevel string `json:"access_level,omitempty"`
	ParentID    string `json:"parent_id,omitempty"`
}

// BookmarksAddResponse is based on:
// https://docs.slack.dev/reference/methods/bookmarks.add/
type BookmarksAddResponse struct {
	Response

	Bookmark *Bookmark `json:"bookmark,omitempty"`
}

// BookmarksAdd is based on:
// https://docs.slack.dev/reference/methods/bookmarks.add/
func BookmarksAdd(ctx workflow.Context, channelID, title, url, emoji string) error {
	req := BookmarksAddRequest{ChannelID: channelID, Title: title, Type: "link", Link: url, Emoji: emoji}
	return internal.ExecuteTimpaniActivityNoResp(ctx, BookmarksAddActivityName, req)
}

// BookmarksEditRequest is based on:
// https://docs.slack.dev/reference/methods/bookmarks.edit/
type BookmarksEditRequest struct {
	ChannelID  string `json:"channel_id"`
	BookmarkID string `json:"bookmark_id"`

	Title string `json:"title,omitempty"`
	Link  string `json:"link,omitempty"`
	Emoji string `json:"emoji,omitempty"`
}

// BookmarksEditResponse is based on:
// https://docs.slack.dev/reference/methods/bookmarks.edit/
type BookmarksEditResponse struct {
	Response

	Bookmark *Bookmark `json:"bookmark,omitempty"`
}

// BookmarksEditTitle is based on:
// https://docs.slack.dev/reference/methods/bookmarks.edit/
func BookmarksEditTitle(ctx workflow.Context, channelID, bookmarkID, title string) error {
	req := BookmarksEditRequest{ChannelID: channelID, BookmarkID: bookmarkID, Title: title}
	return internal.ExecuteTimpaniActivityNoResp(ctx, BookmarksEditActivityName, req)
}

// BookmarksListRequest is based on:
// https://docs.slack.dev/reference/methods/bookmarks.list/
type BookmarksListRequest struct {
	ChannelID string `json:"channel_id"`
}

// BookmarksListResponse is based on:
// https://docs.slack.dev/reference/methods/bookmarks.list/
type BookmarksListResponse struct {
	Response

	Bookmarks []Bookmark `json:"bookmarks,omitempty"`
}

// BookmarksList is based on:
// https://docs.slack.dev/reference/methods/bookmarks.list/
func BookmarksList(ctx workflow.Context, channelID string) ([]Bookmark, error) {
	req := BookmarksListRequest{ChannelID: channelID}
	resp, err := internal.ExecuteTimpaniActivity[BookmarksListResponse](ctx, BookmarksListActivityName, req)
	if err != nil {
		return nil, err
	}
	return resp.Bookmarks, nil
}

// BookmarksRemoveRequest is based on:
// https://docs.slack.dev/reference/methods/bookmarks.remove/
type BookmarksRemoveRequest struct {
	ChannelID  string `json:"channel_id"`
	BookmarkID string `json:"bookmark_id"`

	QuipSectionID string `json:"quip_section_id,omitempty"`
}

// BookmarksRemoveResponse is based on:
// https://docs.slack.dev/reference/methods/bookmarks.remove/
type BookmarksRemoveResponse Response

// BookmarksRemove is based on:
// https://docs.slack.dev/reference/methods/bookmarks.remove/
func BookmarksRemove(ctx workflow.Context, channelID, bookmarkID string) error {
	req := BookmarksRemoveRequest{ChannelID: channelID, BookmarkID: bookmarkID}
	return internal.ExecuteTimpaniActivityNoResp(ctx, BookmarksRemoveActivityName, req)
}

type Bookmark struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
	Title     string `json:"title"`
	Type      string `json:"type"`

	Link     *string `json:"link,omitempty"`
	Emoji    *string `json:"emoji,omitempty"`
	IconURL  *string `json:"icon_url,omitempty"`
	EntityID *string `json:"entity_id,omitempty"`

	DateCreated         int64  `json:"date_created"`
	DateUpdated         int64  `json:"date_updated"`
	Rank                string `json:"rank"`
	LastUpdatedByUserID string `json:"last_updated_by_user_id"`
	LastUpdatedByTeamID string `json:"last_updated_by_team_id"`

	ShortcutID  string `json:"shortcut_id"`
	AppID       string `json:"app_id"`
	AppActionID string `json:"app_action_id"`
}
