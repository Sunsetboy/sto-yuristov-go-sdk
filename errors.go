package stoyuristov

import (
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
)

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
