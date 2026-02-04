package fintech

import (
	"context"
)

// MockClient is a mock implementation of the Fintech Client for testing.
type MockClient struct {
	*Client
	OnCreateIntent func(req *PaymentIntentRequest) (*PaymentResponse, error)
	OnTrigger      func(eventType string, zoneID string, data interface{}) error
}

func NewMockClient() *MockClient {
	return &MockClient{
		Client: &Client{
			Payments: &PaymentsService{},
			Zones:    &ZonesService{},
		},
	}
}

// SimulatePaymentSuccess fires a mock payment.succeeded event.
func (m *MockClient) SimulatePaymentSuccess(ctx context.Context, zoneID string, amount int64) error {
	return m.TriggerEvent(ctx, "payment.succeeded", zoneID, map[string]interface{}{
		"amount": amount,
		"status": "succeeded",
	})
}

// SimulatePaymentFailure fires a mock payment.failed event.
func (m *MockClient) SimulatePaymentFailure(ctx context.Context, zoneID string, reason string) error {
	return m.TriggerEvent(ctx, "payment.failed", zoneID, map[string]interface{}{
		"reason": reason,
		"status": "failed",
	})
}
