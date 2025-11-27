// Package internal requests the Timpani worker to execute
// Temporal activities with preconfigured Temporal options.
package internal

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/pkg/temporal"
)

// ExecuteTimpaniActivity requests the [Timpani worker] to execute one
// of its [activities] on behalf of the calling Temporal workflow, with
// preconfigured [temporal.ActivityOptions] related to timeouts and retries.
//
// [Timpani worker]: https://pkg.go.dev/github.com/tzrikka/timpani
// [activities]: https://pkg.go.dev/github.com/tzrikka/timpani/pkg/api
func ExecuteTimpaniActivity[T any](ctx workflow.Context, name string, req any) (*T, error) {
	opts := temporal.ActivityOptions
	if opts == nil {
		opts = temporal.DefaultActivityOptions("timpani")
	}

	resp := new(T)
	err := workflow.ExecuteActivity(workflow.WithActivityOptions(ctx, *opts), name, req).Get(ctx, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ExecuteTimpaniActivityNoResp is a convenience wrapper around
// [ExecuteTimpaniActivity] for activities that do not return a response.
func ExecuteTimpaniActivityNoResp(ctx workflow.Context, name string, req any) error {
	_, err := ExecuteTimpaniActivity[any](ctx, name, req)
	return err
}
