package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sapliy/fintech-sdk-go"
)

/**
 * Real-world Example: Micro-service Event Bridge
 *
 * This example shows how a Go service can react to internal events to:
 * 1. Provision a new Zone for a tenant
 * 2. Activate a bundle of automation Flows
 */
func main() {
	client := fintech.NewClient("sk_live_bridge_991", fintech.WithBaseURL("http://localhost:8080"))

	orgID := "org_tenant_445"
	ctx := context.Background()

	fmt.Println("--- Step 1: Provisioning Zone for Tenant ---")
	// Use high-level SDK
	zoneID, err := client.Zones.Create(ctx, orgID, "Production Environment", "live", "standard-retail")
	if err != nil {
		log.Fatalf("Failed to create zone: %v", err)
	}
	fmt.Printf("Zone created ID: %s\n", zoneID)

	fmt.Println("\n--- Step 2: Activating Automation Bundle ---")
	// Use high-level SDK
	flows, err := client.Flows.List(ctx, zoneID)
	if err != nil {
		log.Printf("Warning: Could not list flows: %v", err)
	} else if len(flows) > 0 {
		var flowIDs []string
		for _, f := range flows {
			if f.Id != "" {
				flowIDs = append(flowIDs, f.Id)
			}
		}

		err = client.Flows.BulkUpdate(ctx, flowIDs, true)
		if err != nil {
			log.Printf("Failed to activate flows: %v", err)
		} else {
			fmt.Printf("Activated %d flows for the new zone.\n", len(flowIDs))
		}
	}

	fmt.Println("\nBridge operation complete.")
}
