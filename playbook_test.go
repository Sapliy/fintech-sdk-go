package sapliy

import (
	"strings"
	"testing"
)

func TestPlaybooksList(t *testing.T) {
	c := NewClient("test-key")
	playbooks := c.Playbooks.List()

	if len(playbooks) != 3 {
		t.Fatalf("expected 3 playbooks, got %d", len(playbooks))
	}

	want := []string{"revenue-recovery", "refund-approval", "invoice-reminders"}
	for i, w := range want {
		if playbooks[i].Type != w {
			t.Errorf("playbook %d: expected type %q, got %q", i, w, playbooks[i].Type)
		}
		if strings.TrimSpace(playbooks[i].Description) == "" {
			t.Errorf("playbook %q: expected non-empty description", w)
		}
	}
}

func TestPlaybooksBootstrap(t *testing.T) {
	c := NewClient("test-key")

	cfg, err := c.Playbooks.Bootstrap("revenue-recovery")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Type != "revenue-recovery" {
		t.Errorf("expected type revenue-recovery, got %q", cfg.Type)
	}
	if cfg.Description == "" {
		t.Error("expected non-empty description")
	}
	if cfg.Config == nil {
		t.Error("expected non-nil config scaffold")
	}

	cfg, err = c.Playbooks.Bootstrap("refund-approval")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Type != "refund-approval" {
		t.Errorf("expected type refund-approval, got %q", cfg.Type)
	}

	cfg, err = c.Playbooks.Bootstrap("invoice-reminders")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Type != "invoice-reminders" {
		t.Errorf("expected type invoice-reminders, got %q", cfg.Type)
	}

	if _, err := c.Playbooks.Bootstrap("does-not-exist"); err == nil {
		t.Error("expected error for unknown playbook")
	}
}