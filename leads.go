package stoyuristov

import (
	"context"
	"fmt"
)

// UpdateLead updates a lead with the given ID.
//
// Example:
//
//	err := client.UpdateLead(ctx, 456, &LeadUpdateRequest{
//	    Categories: []int{1, 2},
//	    Directions: []int{10, 20},
//	})
func (c *Client) UpdateLead(ctx context.Context, id int, req *LeadUpdateRequest) error {
	path := fmt.Sprintf("/intapi/lead/update/id/%d", id)
	return c.do(ctx, "POST", path, req)
}
