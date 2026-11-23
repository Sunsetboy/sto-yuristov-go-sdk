# sto-yuristov-go-sdk

Go SDK for the 100yuristov internal API.

## Installation

```bash
go get github.com/sunsetboy/sto-yuristov-go-sdk
```

## Usage

```go
package main

import (
    "context"
    "log"
    "time"

    stoyuristov "github.com/sunsetboy/sto-yuristov-go-sdk"
)

func main() {
    // Create a client
    client := stoyuristov.NewClient(
        "http://100yuristov:8080",
        "your-auth-token",
    )

    // With custom timeout
    client = stoyuristov.NewClient(
        "http://100yuristov:8080",
        "your-auth-token",
        stoyuristov.WithTimeout(60*time.Second),
    )

    ctx := context.Background()

    // Update question categories
    err := client.UpdateQuestion(ctx, 123, &stoyuristov.QuestionUpdateRequest{
        Categories: []int{1, 2, 3},
        Directions: []int{10, 20},
    })
    if err != nil {
        log.Fatal(err)
    }

    // Update question title
    err = client.UpdateQuestion(ctx, 123, &stoyuristov.QuestionUpdateRequest{
        Title: stoyuristov.String("New Title"),
        SeoH1: stoyuristov.String("SEO Heading"),
    })
    if err != nil {
        log.Fatal(err)
    }

    // Update lead
    err = client.UpdateLead(ctx, 456, &stoyuristov.LeadUpdateRequest{
        Categories: []int{1, 2},
    })
    if err != nil {
        log.Fatal(err)
    }

    // Update sentiment
    err = client.UpdateSentiment(ctx, 123, "question", 0.85)
    if err != nil {
        log.Fatal(err)
    }

    // Report the result of a web push delivery job
    logID, gone := 12345, 410
    err = client.ReportPush(ctx, &stoyuristov.PushReport{
        MessageID: "4f6c1a52-3b1e-4c8e-9d43-8a7b0c2e1f00",
        LogID:     &logID, // nil if the application did not write the log record
        Status:    stoyuristov.PushReportProcessed,
        Results: []stoyuristov.PushDeliveryResult{
            {SubscriptionID: 55, EndpointHash: "9b74…", Result: stoyuristov.PushResultDelivered},
            {SubscriptionID: 56, EndpointHash: "1c2d…", Result: stoyuristov.PushResultExpired, HTTPStatus: &gone},
        },
    })
    if err != nil {
        log.Fatal(err)
    }
}
```

The report format is described in the main application, `protected/docs/WEB_PUSH_ASYNC.md`.

## Error Handling

The SDK provides typed errors for common HTTP status codes:

```go
import "errors"

err := client.UpdateQuestion(ctx, id, req)
if errors.Is(err, stoyuristov.ErrNotFound) {
    // Handle 404
} else if errors.Is(err, stoyuristov.ErrUnauthorized) {
    // Handle 401
} else if err != nil {
    // Handle other errors
}
```

`IsRetryable(err)` tells whether repeating the request may succeed: network errors,
timeouts (`ErrRequestFailed`) and 5xx responses (`ErrInternalServer`). 4xx and 3xx
responses and a cancelled context are permanent.

The client does not follow redirects: all API methods are POST, and a redirected POST
would be replayed as GET without the body. A 3xx response is returned as `*APIError`.

## Configuration Options

- `WithTimeout(duration)` - Set HTTP client timeout (default: 30s)
- `WithHTTPClient(client)` - Use a custom `*http.Client`

## Versioning

This SDK follows [Semantic Versioning](https://semver.org/).

```bash
# Update to latest version
go get github.com/sunsetboy/sto-yuristov-go-sdk@latest

# Use specific version
go get github.com/sunsetboy/sto-yuristov-go-sdk@v1.0.0
```
