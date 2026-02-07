package fintech

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// --- Flows Service ---

type FlowsService struct {
	client *Client
}

type Flow struct {
	ID          string    `json:"id"`
	ZoneID      string    `json:"zone_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Enabled     bool      `json:"enabled"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateFlowRequest struct {
	ZoneID      string `json:"zone_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Nodes       []Node `json:"nodes"`
	Edges       []Edge `json:"edges"`
}

type Node struct {
	ID   string          `json:"id"`
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

func (s *FlowsService) Create(ctx context.Context, req *CreateFlowRequest) (*Flow, error) {
	var res Flow
	err := s.client.do(ctx, http.MethodPost, "/v1/flows", req, &res)
	return &res, err
}

func (s *FlowsService) Get(ctx context.Context, id string) (*Flow, error) {
	var res Flow
	err := s.client.do(ctx, http.MethodGet, fmt.Sprintf("/v1/flows/%s", id), nil, &res)
	return &res, err
}

func (s *FlowsService) List(ctx context.Context, zoneID string) ([]Flow, error) {
	var res struct {
		Flows []Flow `json:"flows"`
	}
	err := s.client.do(ctx, http.MethodGet, fmt.Sprintf("/v1/zones/%s/flows", zoneID), nil, &res)
	return res.Flows, err
}

func (s *FlowsService) Update(ctx context.Context, id string, req *CreateFlowRequest) (*Flow, error) {
	var res Flow
	err := s.client.do(ctx, http.MethodPut, fmt.Sprintf("/v1/flows/%s", id), req, &res)
	return &res, err
}

func (s *FlowsService) Delete(ctx context.Context, id string) error {
	return s.client.do(ctx, http.MethodDelete, fmt.Sprintf("/v1/flows/%s", id), nil, nil)
}

func (s *FlowsService) Enable(ctx context.Context, id string) error {
	return s.client.do(ctx, http.MethodPost, fmt.Sprintf("/v1/flows/%s/enable", id), nil, nil)
}

func (s *FlowsService) Disable(ctx context.Context, id string) error {
	return s.client.do(ctx, http.MethodPost, fmt.Sprintf("/v1/flows/%s/disable", id), nil, nil)
}
