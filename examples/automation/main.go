package automation

import (
	"context"
	"fmt"
	"log"

	"github.com/sapliy/fintech-sdk-go"
)

//nolint:deadcode
func main() {
	// Initialize SDK
	client := fintech.NewClient("sk_test_51...your_key", fintech.WithBaseURL("http://localhost:8080"))

	// 1. Create a Zone for this example
	// Zones.Create(ctx, orgID, name, mode, templateName) returns (zoneID string, err error)
	zoneID, err := client.Zones.Create(context.Background(), "org_demo", "Automation-Test", "test", "standard")
	if err != nil {
		log.Fatalf("Failed to create zone: %v", err)
	}
	fmt.Printf("Created Zone: %s\n", zoneID)

	// 2. Create a Payment Intent (this will trigger a Kafka event)
	// CreateIntent(ctx, zoneID, amount, currency, description, metadata)
	intent, err := client.Payments.CreateIntent(context.Background(), zoneID, 15000, "USD", "Automation demo", nil)
	if err != nil {
		log.Fatalf("Failed to create intent: %v", err)
	}
	fmt.Printf("Created Payment Intent: %s. This should trigger an automated ledger provision.\n", intent.GetId())

	// In a real scenario, the Flow Engine would now pick up the 'payment.succeeded' event
	// and execute nodes like 'Check Amount > 100' -> 'Provision Ledger'.

	fmt.Println("🚀 Automation example completed. Check Audit Logs in the Dashboard!")
}
