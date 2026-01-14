package github

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/internal"
)

//revive:disable:exported
const (
	UsersGetActivityName  = "github.users.get"
	UsersListActivityName = "github.users.list"
) //revive:enable:exported

// UsersGetRequest is based on:
// https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28
type UsersGetRequest struct {
	AccountID string `json:"account_id,omitempty"`
	Username  string `json:"username,omitempty"`
}

// UsersGetAuthenticated is based on:
// https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-the-authenticated-user
func UsersGetAuthenticated(ctx workflow.Context) (*User, error) {
	return internal.ExecuteTimpaniActivity[User](ctx, UsersGetActivityName, UsersGetRequest{})
}

// UsersGetByAccountID is based on:
// https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-a-user-using-their-id
func UsersGetByAccountID(ctx workflow.Context, accountID string) (*User, error) {
	req := UsersGetRequest{AccountID: accountID}
	return internal.ExecuteTimpaniActivity[User](ctx, UsersGetActivityName, req)
}

// UsersGetByUsername is based on:
// https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-a-user
func UsersGetByUsername(ctx workflow.Context, username string) (*User, error) {
	req := UsersGetRequest{Username: username}
	return internal.ExecuteTimpaniActivity[User](ctx, UsersGetActivityName, req)
}

// UsersListRequest is based on:
//   - https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#list-users
//   - https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api
type UsersListRequest struct {
	Since   int `json:"since,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

// UsersList is based on:
//   - https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#list-users
//   - https://docs.github.com/rest/using-the-rest-api/using-pagination-in-the-rest-api
func UsersList(ctx workflow.Context, since, perPage int) ([]User, error) {
	req := UsersListRequest{Since: since, PerPage: perPage}
	resp, err := internal.ExecuteTimpaniActivity[[]User](ctx, UsersListActivityName, req)
	if err != nil {
		return nil, err
	}
	return *resp, nil
}

// User is based on:
// https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28
type User struct {
	ID        int64  `json:"id"`
	NodeID    string `json:"node_id"`
	HTMLURL   string `json:"html_url"`
	AvatarURL string `json:"avatar_url,omitempty"`

	Login string `json:"login"`
	Type  string `json:"type"` // "User", "Bot", "Organization".
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`

	UserViewType string `json:"user_view_type,omitempty"` // "private", "public".

	CreatedAt time.Time `json:"created_at,omitzero"`
	UpdatedAt time.Time `json:"updated_at,omitzero"`
}
