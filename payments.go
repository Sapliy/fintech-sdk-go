package fintech

import (
	"context"

	"github.com/sapliy/fintech-sdk-go/generated"
)

// PaymentsService provides high-level methods for creating and managing payments.
type PaymentsService struct {
	c *Client
}

func (s *PaymentsService) CreateIntent(ctx context.Context, zoneID string, amount int64, currency string, description string, metadata map[string]string) (*generated.PaymentIntent, error) {
	req := generated.CreatePaymentIntentRequest{
		Amount:   amount,
		Currency: currency,
	}
	if description != "" {
		req.Description = &description
	}
	if len(metadata) > 0 {
		req.Metadata = metadata
	}

	intent, _, err := s.c.gen.PaymentsAPI.CreatePaymentIntent(ctx).
		XZoneID(zoneID).
		CreatePaymentIntentRequest(req).
		Execute()
	return intent, err
}

func (s *PaymentsService) GetIntent(ctx context.Context, id string, zoneID string) (*generated.PaymentIntent, error) {
	intent, _, err := s.c.gen.PaymentsAPI.GetPaymentIntent(ctx, id).
		XZoneID(zoneID).
		Execute()
	return intent, err
}

func (s *PaymentsService) ConfirmIntent(ctx context.Context, id string, zoneID string, paymentMethodID string) (*generated.PaymentIntent, error) {
	intent, _, err := s.c.gen.PaymentsAPI.ConfirmPaymentIntent(ctx, id).
		XZoneID(zoneID).
		ConfirmPaymentIntentRequest(generated.ConfirmPaymentIntentRequest{
			PaymentMethodId: &paymentMethodID,
		}).
		Execute()
	return intent, err
}
