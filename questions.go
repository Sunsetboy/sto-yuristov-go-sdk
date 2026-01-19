package stoyuristov

import (
	"context"
	"fmt"
)

// UpdateQuestion updates a question with the given ID.
// Only non-nil/non-empty fields in the request will be sent to the API.
//
// Example - update categories:
//
//	err := client.UpdateQuestion(ctx, 123, &QuestionUpdateRequest{
//	    Categories: []int{1, 2, 3},
//	})
//
// Example - update title:
//
//	err := client.UpdateQuestion(ctx, 123, &QuestionUpdateRequest{
//	    Title: stoyuristov.String("New Title"),
//	    SeoH1: stoyuristov.String("SEO Heading"),
//	})
func (c *Client) UpdateQuestion(ctx context.Context, id int, req *QuestionUpdateRequest) error {
	path := fmt.Sprintf("/intapi/question/update/id/%d/", id)
	return c.do(ctx, "POST", path, req)
}
