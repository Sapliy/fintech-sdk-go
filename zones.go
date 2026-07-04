package sapliy

import (
	"context"

	"github.com/sapliy/sapliy-sdk-go/generated"
)

// ZonesService handles multi-tenancy scoping and lifecycle.
type ZonesService struct {
	c *Client
}

func (s *ZonesService) Create(ctx context.Context, orgID string, name string, mode string, templateName string) (string, error) {
	req := generated.CreateZoneRequest{
		OrgId: orgID,
		Name:  name,
		Mode:  mode,
	}
	if templateName != "" {
		req.TemplateName = &templateName
	}

	res, _, err := s.c.gen.ZonesAPI.CreateZone(ctx).
		CreateZoneRequest(req).
		Execute()
	if err != nil {
		return "", err
	}
	return res.GetId(), nil
}

func (s *ZonesService) List(ctx context.Context, orgID string) ([]generated.ListZones200ResponseInner, error) {
	res, _, err := s.c.gen.ZonesAPI.ListZones(ctx).
		OrgId(orgID).
		Execute()
	return res, err
}

func (s *ZonesService) Get(ctx context.Context, id string) (*generated.ListZones200ResponseInner, error) {
	res, _, err := s.c.gen.ZonesAPI.ListZones(ctx).
		Id(id).
		Execute()
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return &res[0], nil
}
