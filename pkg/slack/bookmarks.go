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

// https://docs.slack.dev/reference/methods/bookmarks.add/
type BookmarksAddResponse struct {
	Response

	Bookmark *Bookmark `json:"bookmark,omitempty"`
}

// https://docs.slack.dev/reference/methods/bookmarks.add/
func BookmarksAddActivity(ctx workflow.Context, channelID, title, url string) error {
	req := BookmarksAddRequest{ChannelID: channelID, Title: title, Type: "link", Link: url}
	return internal.ExecuteTimpaniActivity(ctx, BookmarksAddActivityName, req).Get(ctx, nil)
}

// https://docs.slack.dev/reference/methods/bookmarks.edit/
type BookmarksEditRequest struct {
	ChannelID  string `json:"channel_id"`
	BookmarkID string `json:"bookmark_id"`

	Title string `json:"title,omitempty"`
	Link  string `json:"link,omitempty"`
	Emoji string `json:"emoji,omitempty"`
}

// https://docs.slack.dev/reference/methods/bookmarks.edit/
type BookmarksEditResponse struct {
	Response

	Bookmark *Bookmark `json:"bookmark,omitempty"`
}

// https://docs.slack.dev/reference/methods/bookmarks.edit/
func BookmarksEditTitleActivity(ctx workflow.Context, channelID, bookmarkID, title string) error {
	req := BookmarksEditRequest{ChannelID: channelID, BookmarkID: bookmarkID, Title: title}
	return internal.ExecuteTimpaniActivity(ctx, BookmarksEditActivityName, req).Get(ctx, nil)
}

// https://docs.slack.dev/reference/methods/bookmarks.list/
type BookmarksListRequest struct {
	ChannelID string `json:"channel_id"`
}

// https://docs.slack.dev/reference/methods/bookmarks.list/
type BookmarksListResponse struct {
	Response

	Bookmarks []Bookmark `json:"bookmarks,omitempty"`
}

// https://docs.slack.dev/reference/methods/bookmarks.list/
func BookmarksListActivity(ctx workflow.Context, channelID string) ([]Bookmark, error) {
	req := BookmarksListRequest{ChannelID: channelID}
	fut := internal.ExecuteTimpaniActivity(ctx, BookmarksListActivityName, req)

	resp := new(BookmarksListResponse)
	if err := fut.Get(ctx, resp); err != nil {
		return nil, err
	}

	return resp.Bookmarks, nil
}

// https://docs.slack.dev/reference/methods/bookmarks.remove/
type BookmarksRemoveRequest struct {
	ChannelID  string `json:"channel_id"`
	BookmarkID string `json:"bookmark_id"`

	QuipSectionID string `json:"quip_section_id,omitempty"`
}

// https://docs.slack.dev/reference/methods/bookmarks.remove/
type BookmarksRemoveResponse Response

// https://docs.slack.dev/reference/methods/bookmarks.remove/
func BookmarksRemoveActivity(ctx workflow.Context, channelID, bookmarkID string) error {
	req := BookmarksRemoveRequest{ChannelID: channelID, BookmarkID: bookmarkID}
	return internal.ExecuteTimpaniActivity(ctx, BookmarksRemoveActivityName, req).Get(ctx, nil)
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
