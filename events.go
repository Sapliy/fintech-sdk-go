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
