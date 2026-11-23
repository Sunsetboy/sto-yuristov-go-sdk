package stoyuristov

import "context"

// ReportPush sends the result of a web push delivery job to the main
// application: POST /intapi/push/report/.
//
// The application applies the results to the user's subscriptions and to the
// push log. A repeated report for the same log record is accepted and ignored,
// so the call is safe to retry; use IsRetryable to decide whether to.
//
// Example:
//
//	err := client.ReportPush(ctx, &stoyuristov.PushReport{
//	    MessageID: job.MessageID,
//	    LogID:     job.LogID,
//	    Status:    stoyuristov.PushReportProcessed,
//	    Results:   results,
//	})
func (c *Client) ReportPush(ctx context.Context, report *PushReport) error {
	body := *report
	if body.Results == nil {
		body.Results = []PushDeliveryResult{} // the API expects [] and not null
	}
	return c.do(ctx, "POST", "/intapi/push/report/", &body)
}
