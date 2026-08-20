package sapliy

import (
	"fmt"
)

// Playbook is an operational playbook in the Sapliy MVP catalog.
type Playbook struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

// PlaybookCatalog is the static catalog of MVP operational playbooks.
var PlaybookCatalog = []Playbook{
	{
		Type:        "revenue-recovery",
		Description: "Recover failed subscription payments with automated dunning and smart retries",
	},
	{
		Type:        "refund-approval",
		Description: "Route refunds and invoice adjustments through the policy engine for approval",
	},
	{
		Type:        "invoice-reminders",
		Description: "Send automated reminders for overdue invoices",
	},
}

// BootstrapConfig is the scaffold configuration returned by PlaybooksService.Bootstrap.
type BootstrapConfig struct {
	Type        string                 `json:"type"`
	Description string                 `json:"description"`
	Config      map[string]interface{} `json:"config"`
}

// PlaybooksService provides access to the operational playbook catalog.
type PlaybooksService struct {
	c *Client
}

// List returns all playbooks in the MVP catalog.
func (s *PlaybooksService) List() []Playbook {
	return PlaybookCatalog
}

// Bootstrap returns a scaffold configuration for the given playbook kind.
func (s *PlaybooksService) Bootstrap(kind string) (*BootstrapConfig, error) {
	for _, p := range PlaybookCatalog {
		if p.Type == kind {
			return &BootstrapConfig{
				Type:        p.Type,
				Description: p.Description,
				Config:      map[string]interface{}{},
			}, nil
		}
	}
	return nil, fmt.Errorf("playbook %q not found", kind)
}