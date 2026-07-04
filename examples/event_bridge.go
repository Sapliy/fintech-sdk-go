package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sapliy/sapliy-sdk-go"
)

/**
 * Real-world Example: Micro-service Event Bridge
 *
 * This example shows how a Go service can react to internal events to:
 * 1. Provision a new Zone for a tenant
 * 2. Activate a bundle of automation Flows
 */
func main() {
	client := sapliy.NewClient("sk_live_bridge_991", sapliy.WithBaseURL("http://localhost:8080"))

	orgID := "org_tenant_445"
	ctx := context.Background()

	fmt.Println("--- Step 1: Provisioning Zone for Tenant ---")
	// Zones.Create(ctx, orgID, name, mode, templateName) returns (zoneID string, err error)
	zoneID, err := client.Zones.Create(ctx, orgID, "Production Environment", "live", "standard-retail")
	if err != nil {
		log.Fatalf("Failed to create zone: %v", err)
	}
	fmt.Printf("Zone created ID: %s\n", zoneID)

	fmt.Println("\n--- Step 2: Activating Automation Bundle ---")
	// Flows.List(ctx, zoneID) returns ([]AutomationFlow, error)
	flows, err := client.Flows.List(ctx, zoneID)
	if err != nil {
		log.Printf("Warning: Could not list flows: %v", err)
	} else if len(flows) > 0 {
		fmt.Printf("Found %d flows for the new zone.\n", len(flows))
		for _, f := range flows {
			fmt.Printf("  - Flow: %s (%s)\n", f.GetName(), f.GetId())
		}
	}

	fmt.Println("\nBridge operation complete.")
}
