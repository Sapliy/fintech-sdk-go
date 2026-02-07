package fintech

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// EmitOption configures an Emit call
type EmitOption func(*emitOptions)

type emitOptions struct {
	idempotencyKey string
	source         string
	env            string
}

// WithIdempotencyKey sets a custom idempotency key
func WithIdempotencyKey(key string) EmitOption {
	return func(o *emitOptions) {
		o.idempotencyKey = key
	}
}

// WithSource sets the source metadata
func WithSource(source string) EmitOption {
	return func(o *emitOptions) {
		o.source = source
	}
}

// WithEnv sets the environment metadata
func WithEnv(env string) EmitOption {
	return func(o *emitOptions) {
		o.env = env
	}
}

type EventEmitResponse struct {
	Status  string `json:"status"`
	EventID string `json:"event_id,omitempty"`
	Topic   string `json:"topic,omitempty"`
	Message string `json:"message,omitempty"`
}

// Emit sends an event to the Sapliy Gateway
func (c *Client) Emit(ctx context.Context, eventType string, data interface{}, opts ...EmitOption) (*EventEmitResponse, error) {
	options := &emitOptions{
		source: "sdk-go",
		env:    "development", // Default, should be overridden or detected
	}
	for _, opt := range opts {
		opt(options)
	}

	// Auto-generate idempotency key if not provided
	if options.idempotencyKey == "" {
		payloadBytes, _ := json.Marshal(data)
		hash := sha256.Sum256(payloadBytes)
		options.idempotencyKey = fmt.Sprintf("%s-%x-%d", eventType, hash[:8], time.Now().UnixNano())
	}

	payload := map[string]interface{}{
		"type": eventType,
		"data": data,
		"meta": map[string]string{
			"source": options.source,
			"env":    options.env,
		},
	}

	req, err := c.newRequest(ctx, http.MethodPost, "/v1/events/emit", payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Idempotency-Key", options.idempotencyKey)

	var res EventEmitResponse
	if err := c.doRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Helper methods to reuse client logic (assuming they are exported or I can access them)
// Since 'do' is private in sapliyio.go, I might need to duplicate some logic or export 'do'
// Let's check sapliyio.go again. 'do' is private.
// I can add this file to the same package 'fintech' so it can access 'do'.
// However, 'do' takes method, path, body, out. It doesn't allow setting custom headers easily inside 'do'
// because 'do' creates the request.
//
// I need to refactor 'do' or implement a variant that allows custom headers.
// 'do' in sapliyio.go:
// func (c *Client) do(...) { ... req, err := http.NewRequestWithContext(...) ... req.Header.Set("X-API-Key", ...) ... }
//
// I will add a 'doRequest' method to Client in 'sapliyio.go' that takes a *http.Request,
// and 'newRequest' that creates it.
//
// But first, let's write this file and then I will refactor sapliyio.go to support it.
