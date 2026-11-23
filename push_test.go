package stoyuristov

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestReportPush_SendsReport(t *testing.T) {
	var method, path, token string
	var body map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method, path, token = r.Method, r.URL.Path, r.Header.Get("X-Auth-Token")
		raw, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(raw, &body); err != nil {
			t.Errorf("invalid JSON body %q: %v", raw, err)
		}
		w.Write([]byte(`{"message":"Report applied"}`))
	}))
	defer srv.Close()

	logID, status := 12345, 410
	err := NewClient(srv.URL+"/", "secret").ReportPush(context.Background(), &PushReport{
		MessageID: "4f6c1a52-3b1e-4c8e-9d43-8a7b0c2e1f00",
		LogID:     &logID,
		Status:    PushReportProcessed,
		Results: []PushDeliveryResult{
			{SubscriptionID: 55, EndpointHash: "9b74", Result: PushResultDelivered},
			{SubscriptionID: 56, EndpointHash: "1c2d", Result: PushResultExpired, HTTPStatus: &status, Reason: "Gone"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if method != http.MethodPost || path != "/intapi/push/report/" || token != "secret" {
		t.Errorf("unexpected request: %s %s token=%q", method, path, token)
	}
	results := body["results"].([]any)
	first, second := results[0].(map[string]any), results[1].(map[string]any)
	if body["log_id"] != float64(12345) || first["http_status"] != nil || second["http_status"] != float64(410) {
		t.Errorf("unexpected body: %v", body)
	}
	if _, ok := first["reason"]; ok {
		t.Errorf("empty reason must be omitted: %v", first)
	}
}

func TestReportPush_EncodesNullLogIDAndEmptyResults(t *testing.T) {
	var raw []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ = io.ReadAll(r.Body)
	}))
	defer srv.Close()

	report := &PushReport{MessageID: "m", Status: PushReportExpired}
	if err := NewClient(srv.URL, "k").ReportPush(context.Background(), report); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), `"log_id":null`) || !strings.Contains(string(raw), `"results":[]`) {
		t.Errorf("unexpected body: %s", raw)
	}
	if report.Results != nil {
		t.Error("the caller's report must not be modified")
	}
}

func TestClient_DoesNotFollowRedirects(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("redirect must not be followed, got %s %s", r.Method, r.URL.Path)
		}
		http.Redirect(w, r, "/elsewhere/", http.StatusMovedPermanently)
	}))
	defer srv.Close()

	err := NewClient(srv.URL, "k").ReportPush(context.Background(), &PushReport{MessageID: "m"})
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusMovedPermanently {
		t.Fatalf("want 301 APIError, got %v", err)
	}
	if IsRetryable(err) {
		t.Error("a redirect is not retryable")
	}
}

func TestIsRetryable(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Query().Get("status") {
		case "503":
			w.WriteHeader(http.StatusServiceUnavailable)
		case "400":
			w.WriteHeader(http.StatusBadRequest)
		case "403":
			w.WriteHeader(http.StatusForbidden)
		case "slow":
			time.Sleep(200 * time.Millisecond)
		}
	}))
	defer srv.Close()

	call := func(c *Client, status string) error {
		return c.do(context.Background(), "POST", "/?status="+status, nil)
	}
	client := NewClient(srv.URL, "k")

	if err := call(client, "503"); !IsRetryable(err) {
		t.Errorf("5xx must be retryable: %v", err)
	}
	if err := call(client, "400"); err == nil || IsRetryable(err) {
		t.Errorf("400 must be permanent: %v", err)
	}
	if err := call(client, "403"); !errors.Is(err, ErrForbidden) || IsRetryable(err) {
		t.Errorf("403 must be permanent: %v", err)
	}
	if err := call(NewClient(srv.URL, "k", WithTimeout(20*time.Millisecond)), "slow"); !IsRetryable(err) {
		t.Errorf("timeout must be retryable: %v", err)
	}
	if err := call(NewClient("http://127.0.0.1:1", "k"), ""); !IsRetryable(err) {
		t.Errorf("connection error must be retryable: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := client.ReportPush(ctx, &PushReport{MessageID: "m"}); err == nil || IsRetryable(err) {
		t.Errorf("cancelled context must not be retryable: %v", err)
	}
	if IsRetryable(nil) {
		t.Error("nil is not retryable")
	}
}
