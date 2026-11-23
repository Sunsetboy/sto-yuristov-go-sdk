package stoyuristov

import (
	"context"
	"errors"
	"fmt"
)

var (
	// ErrUnauthorized is returned when the API returns a 401 status.
	ErrUnauthorized = errors.New("unauthorized: invalid or missing auth token")

	// ErrForbidden is returned when the API returns a 403 status.
	ErrForbidden = errors.New("forbidden: access denied")

	// ErrNotFound is returned when the API returns a 404 status.
	ErrNotFound = errors.New("not found")

	// ErrInternalServer is returned when the API returns a 5xx status.
	ErrInternalServer = errors.New("internal server error")

	// ErrRequestFailed wraps errors of sending the request or reading the
	// response: connection refused, DNS, timeouts.
	ErrRequestFailed = errors.New("failed to make request")
)

// IsRetryable reports whether repeating the same request may succeed:
// network errors, timeouts and 5xx responses. Other API errors (4xx, 3xx) are
// permanent, as is the cancellation of the caller's context.
func IsRetryable(err error) bool {
	if err == nil || errors.Is(err, context.Canceled) {
		return false
	}
	return errors.Is(err, ErrInternalServer) || errors.Is(err, ErrRequestFailed)
}

// APIError represents an error response from the API.
type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Body)
}

// newAPIError creates an appropriate error based on status code.
func newAPIError(statusCode int, body string) error {
	switch statusCode {
	case 401:
		return ErrUnauthorized
	case 403:
		return ErrForbidden
	case 404:
		return ErrNotFound
	}

	if statusCode >= 500 {
		return fmt.Errorf("%w: %s", ErrInternalServer, body)
	}

	return &APIError{
		StatusCode: statusCode,
		Body:       body,
	}
}
