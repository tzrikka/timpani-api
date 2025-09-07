# Timpani API

[![Go Reference](https://pkg.go.dev/badge/github.com/tzrikka/timpani-api.svg)](https://pkg.go.dev/github.com/tzrikka/timpani-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/tzrikka/timpani-api)](https://goreportcard.com/report/github.com/tzrikka/timpani-api)

Thin, easy-to-use API wrappers for [Temporal](https://temporal.io/) workflows that use [Timpani](https://github.com/tzrikka/timpani).

This package is not required in order to use Timpani. However, it saves time, simplifies code, minimizes duplication, and prevents mistakes.

Timpani is a Temporal worker that sends API calls and receives asynchronous event notifications to/from various well-known third-party services.

## Usage

First, when your Temporal worker starts, initialize the Temporal [activity options](https://pkg.go.dev/go.temporal.io/sdk/internal#ActivityOptions) for all Timpani activities.

```go
import (
    "github.com/tzrikka/timpani-api/pkg/temporal"
    "go.temporal.io/sdk/workflow"
)

temporal.ActivityOptions = temporal.DefaultActivityOptions("my-task-queue-name")

// Or:

temporal.ActivityOptions = &workflow.ActivityOptions{ /* ... */ }
```

Without this, the package's defaults are `DefaultActivityOptions("timpani")`, i.e.:

```go
workflow.ActivityOptions{
    TaskQueue: "timpani",
    StartToCloseTimeout: 5 * time.Second,
    RetryPolicy: &temporal.RetryPolicy{
        MaximumAttempts: 5,
    },
}
```

Then you may call any `*Activity()` function from any [`timpani-api`](https://pkg.go.dev/github.com/tzrikka/timpani-api/pkg) subpackage.

You may also call Temporal's [`workflow.ExecuteActivity()`](https://pkg.go.dev/go.temporal.io/sdk/workflow#ExecuteActivity) function directly, and just use the following from any [`timpani-api`](https://pkg.go.dev/github.com/tzrikka/timpani-api/pkg) subpackage:

- `*ActivityName` string as the `activity` parameter
- `*Request` struct as the `args` parameter
- Reference to an empty `*Response` struct as the `valuePtr` parameter in [`workflow.Future.Get()`](https://pkg.go.dev/go.temporal.io/sdk/internal#Future) calls
