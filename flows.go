package fintech

import (
	"context"

	"github.com/sapliy/fintech-sdk-go/generated"
)

// FlowsService manages business automation flows and rule-based logic.
type FlowsService struct {
	c *Client
}

func (s *FlowsService) Create(ctx context.Context, flow generated.AutomationFlow) (*generated.AutomationFlow, error) {
	res, _, err := s.c.gen.FlowsAPI.CreateFlow(ctx).
		AutomationFlow(flow).
		Execute()
	return res, err
}

func (s *FlowsService) Get(ctx context.Context, id string) (*generated.AutomationFlow, error) {
	res, _, err := s.c.gen.FlowsAPI.GetFlow(ctx, id).
		Execute()
	return res, err
}

func (s *FlowsService) List(ctx context.Context, zoneID string) ([]generated.AutomationFlow, error) {
	res, _, err := s.c.gen.FlowsAPI.ListFlows(ctx, zoneID).
		Execute()
	if err != nil {
		return nil, err
	}
	return res.GetFlows(), nil
}

func (s *FlowsService) Update(ctx context.Context, id string, flow generated.AutomationFlow) (*generated.AutomationFlow, error) {
	res, _, err := s.c.gen.FlowsAPI.UpdateFlow(ctx, id).
		AutomationFlow(flow).
		Execute()
	return res, err
}

func (s *FlowsService) Delete(ctx context.Context, id string) error {
	_, err := s.c.gen.FlowsAPI.DeleteFlow(ctx, id).
		Execute()
	return err
}

// Executions

func (s *FlowsService) GetExecution(ctx context.Context, executionID string) (*generated.AutomationFlowExecution, error) {
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
