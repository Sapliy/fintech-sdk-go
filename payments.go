package fintech

import (
	"context"

	"github.com/sapliy/fintech-sdk-go/generated"
)

// PaymentsService provides high-level methods for creating and managing payments.
type PaymentsService struct {
	c *Client
}

type CreatePaymentRequest struct {
	Amount      int64             `json:"amount"`
	Currency    string            `json:"currency"`
	Description string            `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	ZoneID      string            `json:"zone_id"`
}

func (s *PaymentsService) CreateIntent(ctx context.Context, req *CreatePaymentRequest) (*generated.PaymentIntent, error) {
	genReq := s.c.gen.PaymentsAPI.CreatePaymentIntent(ctx).
		CreatePaymentIntentRequest(generated.CreatePaymentIntentRequest{
			Amount:      req.Amount,
			Currency:    req.Currency,
			Description: &req.Description,
			Metadata:    req.Metadata,
		}).
		XZoneID(req.ZoneID)

	intent, _, err := genReq.Execute()
	return intent, err
}

// func (s *PaymentsService) GetIntent(ctx context.Context, id string, zoneID string) (*generated.PaymentIntent, error) {
// 	intent, _, err := s.c.gen.PaymentsAPI.GetPaymentIntent(ctx, id).
// 		XZoneID(zoneID).
// 		Execute()
// 	return intent, err
// }

func (s *PaymentsService) ConfirmIntent(ctx context.Context, id string, zoneID string, paymentMethodID string) (*generated.PaymentIntent, error) {
	intent, _, err := s.c.gen.PaymentsAPI.ConfirmPaymentIntent(ctx, id).
		XZoneID(zoneID).
		ConfirmPaymentIntentRequest(generated.ConfirmPaymentIntentRequest{
			PaymentMethodId: &paymentMethodID,
		}).
		Execute()
	return intent, err
}
