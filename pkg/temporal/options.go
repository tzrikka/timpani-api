// Package temporal provides common, modifiable activity options
// for all Timpani activities, related to timeouts and retries.
package temporal

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// ActivityOptions are used when executing all Timpani activities.
// Temporal workers that use Timpani may override this global variable
// to set a custom Temporal task queue name for the Timpani worker.
var ActivityOptions = DefaultActivityOptions("timpani")

// DefaultActivityOptions adds a few reasonable values
// to the [workflow.ActivityOptions] defaults:
//   - Maximum number of attempts = 5,
//   - Maximum runtime for each attempt = 5 seconds.
func DefaultActivityOptions(taskQueue string) *workflow.ActivityOptions {
	return &workflow.ActivityOptions{
		TaskQueue:           taskQueue,
		StartToCloseTimeout: 5 * time.Second,
		RetryPolicy: &RetryPolicy{
			MaximumAttempts: 5,
		},
	}
}

type RetryPolicy = temporal.RetryPolicy
