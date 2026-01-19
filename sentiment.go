package stoyuristov

import (
	"context"
	"fmt"
)

// UpdateSentiment updates sentiment data for an entity.
//
// entityType is typically "question" or "lead".
//
// Example:
//
//	err := client.UpdateSentiment(ctx, 123, "question", 0.85)
func (c *Client) UpdateSentiment(ctx context.Context, id int, entityType string, positivity float64) error {
	path := fmt.Sprintf("/intapi/sentiment/update/id/%d/", id)
	req := &SentimentUpdateRequest{
		Type:       entityType,
		ID:         id,
		Positivity: positivity,
	}
	return c.do(ctx, "POST", path, req)
}
