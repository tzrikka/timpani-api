package slack

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	BotsInfoActivityName = "slack.bots.info"
)

// BotsInfoRequest is based on:
// https://docs.slack.dev/reference/methods/bots.info/
type BotsInfoRequest struct {
	Bot string `json:"bot"`

	TeamID string `json:"team_id,omitempty"`
}

// BotsInfoResponse is based on:
// https://docs.slack.dev/reference/methods/bots.info/
type BotsInfoResponse struct {
	Response

	Bot *Bot `json:"bot,omitempty"`
}

// Bot is based on:
// https://docs.slack.dev/reference/methods/bots.info/
type Bot struct {
	ID      string `json:"id"`
	TeamID  string `json:"team_id"`
	Name    string `json:"name"`
	AppID   string `json:"app_id"`
	UserID  string `json:"user_id"`
	Deleted bool   `json:"deleted"`
	Updated int64  `json:"updated"`
}

// BotsInfo is based on:
// https://docs.slack.dev/reference/methods/bots.info/
func BotsInfo(ctx workflow.Context, botID string) (*Bot, error) {
	req := BotsInfoRequest{Bot: botID}
	resp, err := internal.ExecuteTimpaniActivity[BotsInfoResponse](ctx, BotsInfoActivityName, req)
	if err != nil {
		return nil, err
	}
	return resp.Bot, nil
}
