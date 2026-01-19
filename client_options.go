package stoyuristov

import (
	"net/http"
	"time"
)

// Option configures a Client.
type Option func(*Client)

// WithTimeout sets the HTTP client timeout.
// Default is 30 seconds.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithHTTPClient sets a custom HTTP client.
// Use this for advanced configurations like custom transport, proxies, etc.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
