package fintech

import (
	"context"

	"github.com/sapliy/fintech-sdk-go/generated"
)

// FlowsService manages business automation flows and rule-based logic.
type FlowsService struct {
	c *Client
}

func (s *FlowsService) Create(ctx context.Context, flow generated.Flow) (*generated.Flow, error) {
	res, _, err := s.c.gen.FlowsAPI.CreateFlow(ctx).
		Flow(flow).
		Execute()
	return res, err
}

func (s *FlowsService) Get(ctx context.Context, id string) (*generated.Flow, error) {
	res, _, err := s.c.gen.FlowsAPI.GetFlow(ctx, id).
		Execute()
	return res, err
}

func (s *FlowsService) List(ctx context.Context, zoneID string) ([]generated.Flow, error) {
	res, _, err := s.c.gen.FlowsAPI.ListFlows(ctx, zoneID).
		Execute()
	if err != nil {
		return nil, err
	}
	return res.GetFlows(), nil
}

func (s *FlowsService) Update(ctx context.Context, id string, flow generated.Flow) (*generated.Flow, error) {
	res, _, err := s.c.gen.FlowsAPI.UpdateFlow(ctx, id).
		Flow(flow).
		Execute()
	return res, err
}

func (s *FlowsService) Delete(ctx context.Context, id string) error {
	_, err := s.c.gen.FlowsAPI.DeleteFlow(ctx, id).
		Execute()
	return err
}

func (s *FlowsService) BulkUpdate(ctx context.Context, flowIDs []string, enabled bool) error {
	// Fallback to iterative update since bulk endpoint is missing in generated client
	for _, id := range flowIDs {
		// Fetch existing flow first to preserve other fields
		if id == "" {
			continue
		}

		f, err := s.Get(ctx, id)
		if err != nil {
			return err
		}

		if f == nil {
			continue
		}

		// Update enabled status
		f.Enabled = &enabled

		// Execute update
		_, _, err = s.c.gen.FlowsAPI.UpdateFlow(ctx, id).
			Flow(*f).
			Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

// Executions

func (s *FlowsService) GetExecution(ctx context.Context, executionID string) (*generated.FlowExecution, error) {
	res, _, err := s.c.gen.ExecutionsAPI.GetExecution(ctx, executionID).
		Execute()
	return res, err
}

func (s *FlowsService) ResumeExecution(ctx context.Context, executionID string, input map[string]interface{}) (string, error) {
	res, _, err := s.c.gen.ExecutionsAPI.ResumeExecution(ctx, executionID).
		RequestBody(input).
		Execute()
	if err != nil {
		return "", err
	}
	return res.GetMessage(), nil
}
