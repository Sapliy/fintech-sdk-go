package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sapliy/fintech-sdk-go"
	"github.com/sapliy/fintech-sdk-go/generated"
)

/**
 * Real-world Example: Micro-service Event Bridge
 *
 * This example shows how a Go service can react to internal events to:
 * 1. Provision a new Zone for a tenant
 * 2. Activate a bundle of automation Flows
 */
func main() {
	client := fintech.NewSapliyClient("sk_live_bridge_991", "http://localhost:8080")

	orgID := "org_tenant_445"
	ctx := context.Background()

	fmt.Println("--- Step 1: Provisioning Zone for Tenant ---")
	zoneReq := generated.ZoneCreateZoneRequest{
		OrgId:        &orgID,
		Name:         generated.PtrString("Production Environment"),
		Mode:         generated.PtrString("live"),
		TemplateName: generated.PtrString("standard-retail"),
	}

	zone, _, err := client.Zones.ZoneServiceCreateZone(ctx).Body(zoneReq).Execute()
	if err != nil {
		log.Fatalf("Failed to create zone: %v", err)
	}
	fmt.Printf("Zone created: %s (ID: %s)\n", *zone.Name, *zone.Id)

	fmt.Println("\n--- Step 2: Activating Automation Bundle ---")
	// We might have a set of flows that need to be enabled for this new zone
	// For this example, we'll fetch flows and enable them in bulk

	flows, _, err := client.Flows.FlowServiceListFlows(ctx).ZoneId(*zone.Id).Execute()
	if err != nil {
		log.Printf("Warning: Could not list flows: %v", err)
	} else if len(flows.Flows) > 0 {
		var flowIDs []string
		for _, f := range flows.Flows {
			flowIDs = append(flowIDs, *f.Id)
		}

		bulkReq := generated.FlowBulkUpdateFlowsRequest{
			FlowIds: flowIDs,
			Enabled: generated.PtrBool(true),
		}

		_, _, err = client.Flows.FlowServiceBulkUpdateFlows(ctx).Body(bulkReq).Execute()
		if err != nil {
			log.Printf("Failed to activate flows: %v", err)
		} else {
			fmt.Printf("Activated %d flows for the new zone.\n", len(flowIDs))
		}
	}

	fmt.Println("\nBridge operation complete.")
}
