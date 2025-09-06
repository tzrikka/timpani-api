package internal

import (
	"go.temporal.io/sdk/workflow"

	"github.com/tzrikka/timpani-api/pkg/temporal"
)

// ExecuteTimpaniActivity requests the execution of a [Timpani activity] in the
// context of a Temporal workflow, with preconfigured [temporal.ActivityOptions]
// related to timeouts and retries.
//
// [Timpani]: https://pkg.go.dev/github.com/tzrikka/timpani/pkg/api
func ExecuteTimpaniActivity(ctx workflow.Context, name string, req any) workflow.Future {
	if temporal.ActivityOptions != nil {
		ctx = workflow.WithActivityOptions(ctx, *temporal.ActivityOptions)
	}
	return workflow.ExecuteActivity(ctx, name, req)
}
