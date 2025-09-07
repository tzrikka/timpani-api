// Package temporal provides common, modifiable activity options
// for all Timpani activities, related to timeouts and retries.
package temporal

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// ActivityOptions are used when executing all Timpani activities.
// Temporal workers that use Timpani should override this global
// variable in order to set the Temporal task queue name.
var ActivityOptions = DefaultActivityOptions("timpani")

// DefaultActivityOptions adds a few reasonable values
// to the [workflow.ActivityOptions] defaults.
//   - Maximum number of attempts = 5
//   - Maximum runtime for each attempt = 5 seconds
func DefaultActivityOptions(taskQueue string) *workflow.ActivityOptions {
	return &workflow.ActivityOptions{
		TaskQueue:           taskQueue,
		StartToCloseTimeout: 5 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 5,
		},
	}
}
