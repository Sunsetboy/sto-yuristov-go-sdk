package stoyuristov

// QuestionUpdateRequest represents a request to update a question.
// All fields are optional - only non-nil fields will be included in the request.
type QuestionUpdateRequest struct {
	Categories  []int   `json:"categories,omitempty"`
	Directions  []int   `json:"directions,omitempty"`
	Title       *string `json:"title,omitempty"`
	SeoH1       *string `json:"seoH1,omitempty"`
	Description *string `json:"description,omitempty"`
}

// LeadUpdateRequest represents a request to update a lead.
type LeadUpdateRequest struct {
	Categories []int `json:"categories,omitempty"`
	Directions []int `json:"directions,omitempty"`
}

// SentimentUpdateRequest represents a request to update sentiment data.
type SentimentUpdateRequest struct {
	Type       string  `json:"type"`
	ID         int     `json:"id"`
	Positivity float64 `json:"positivity"`
}

// String returns a pointer to the given string.
// Useful for setting optional string fields in requests.
func String(s string) *string {
	return &s
}

// Statuses of a PushReport.
const (
	// PushReportProcessed means the pushes were sent; Results holds the outcomes.
	PushReportProcessed = "processed"
	// PushReportExpired means the job was received after its expiry time and
	// nothing was sent; Results is empty.
	PushReportExpired = "expired"
)

// Outcomes of sending a push to one subscription (PushDeliveryResult.Result).
const (
	// PushResultDelivered means the push service accepted the notification (2xx).
	PushResultDelivered = "delivered"
	// PushResultExpired means the subscription no longer exists (404/410).
	PushResultExpired = "expired"
	// PushResultFailed means any other error: another HTTP status, a timeout,
	// an encryption error.
	PushResultFailed = "failed"
)

// PushReport is the result of a web push delivery job.
type PushReport struct {
	MessageID string `json:"message_id"`
	// LogID is nil when the main application could not write the push log record.
	LogID *int `json:"log_id"`
	// Status is PushReportProcessed or PushReportExpired.
	Status  string               `json:"status"`
	Results []PushDeliveryResult `json:"results"`
}

// PushDeliveryResult is the outcome of sending a push to one subscription.
type PushDeliveryResult struct {
	SubscriptionID int    `json:"subscription_id"`
	EndpointHash   string `json:"endpoint_hash"`
	// Result is PushResultDelivered, PushResultExpired or PushResultFailed.
	Result string `json:"result"`
	// HTTPStatus is the push service's response status, nil for network and
	// encryption errors.
	HTTPStatus *int `json:"http_status"`
	// Reason describes a failure, at most 500 characters are stored.
	Reason string `json:"reason,omitempty"`
}
