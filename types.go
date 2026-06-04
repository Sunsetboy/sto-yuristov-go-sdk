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
