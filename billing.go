package fintech

import (
	"context"

	"github.com/sapliy/fintech-sdk-go/generated"
)

// BillingService manages subscriptions and recurring payments.
type BillingService struct {
	c *Client
}

func (s *BillingService) CreateSubscription(ctx context.Context, planID string, customerID string) (*generated.Subscription, error) {
	sub, _, err := s.c.gen.BillingAPI.V1BillingSubscriptionsPost(ctx).
		V1BillingSubscriptionsPostRequest(generated.V1BillingSubscriptionsPostRequest{
			PlanId:     planID,
			CustomerId: &customerID,
		}).
		Execute()
	return sub, err
}

func (s *BillingService) GetSubscription(ctx context.Context, id string) (*generated.Subscription, error) {
	sub, _, err := s.c.gen.BillingAPI.GetSubscription(ctx, id).
		Execute()
	return sub, err
}

func (s *BillingService) CancelSubscription(ctx context.Context, id string) error {
	_, err := s.c.gen.BillingAPI.CancelSubscription(ctx, id).
		Execute()
	return err
}
