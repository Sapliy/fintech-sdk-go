package sapliy

import (
	"context"

	"github.com/sapliy/sapliy-sdk-go/generated"
)

// EventsService provides high-level methods for emitting and replaying events.
type EventsService struct {
	c *Client
}

func (s *EventsService) Emit(ctx context.Context, eventType string, data map[string]interface{}, idempotencyKey string) (string, error) {
	req := generated.EmitEventRequest{
		Type: eventType,
		Data: data,
	}
	if idempotencyKey != "" {
		req.IdempotencyKey = &idempotencyKey
	}

	res, _, err := s.c.gen.EventsAPI.EmitEvent(ctx).
		EmitEventRequest(req).
		Execute()
	if err != nil {
		return "", err
	}

	return res.GetEventId(), nil
}

func (s *EventsService) Replay(ctx context.Context, eventID string, zoneID string) (string, error) {
	res, _, err := s.c.gen.EventsAPI.ReplayEvent(ctx, eventID).
		ReplayEventRequest(generated.ReplayEventRequest{
			ZoneId: zoneID,
		}).
		Execute()
	if err != nil {
		return "", err
	}
	return res.GetEventId(), nil
}

func (s *EventsService) ListPast(ctx context.Context, zoneID string, limit int32, offset int32) (*generated.GetPastEvents200Response, error) {
	res, _, err := s.c.gen.EventsAPI.GetPastEvents(ctx, zoneID).
		Limit(limit).
		Offset(offset).
		Execute()
	return res, err
}
