package slack

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

const (
	AuthTestActivityName = "slack.auth.test"
)

// https://docs.slack.dev/reference/methods/auth.test/
type AuthTestResponse struct {
	Response

	URL                 string `json:"url,omitempty"`
	Team                string `json:"team,omitempty"`
	User                string `json:"user,omitempty"`
	TeamID              string `json:"team_id,omitempty"`
	UserID              string `json:"user_id,omitempty"`
	BotID               string `json:"bot_id,omitempty"`
	EnterpriseID        string `json:"enterprise_id,omitempty"`
	IsEnterpriseInstall bool   `json:"is_enterprise_install,omitempty"`
}

// https://docs.slack.dev/reference/methods/auth.test/
func AuthTestActivity(ctx workflow.Context) (*AuthTestResponse, error) {
	resp := new(AuthTestResponse)
	err := internal.ExecuteTimpaniActivity(ctx, AuthTestActivityName, nil).Get(ctx, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
