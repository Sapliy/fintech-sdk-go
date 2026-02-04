package automation

import (
	"context"
	"fmt"
	"log"

	"github.com/sapliy/fintech-sdk-go"
)

func main() {
	// Initialize SDK
	client := fintech.NewClient("sk_test_51...your_key", fintech.WithBaseURL("http://localhost:8080"))

	// 1. Create a Zone for this example
	zone, err := client.Zones.Create(context.Background(), &fintech.CreateZoneRequest{
		Name: "Automation-Test",
		Mode: "test",
	})
	if err != nil {
		log.Fatalf("Failed to create zone: %v", err)
	}
	fmt.Printf("Created Zone: %s\n", zone.ID)

	// 2. Create a Payment Intent (this will trigger a Kafka event)
	intent, err := client.Payments.CreateIntent(context.Background(), &fintech.PaymentIntentRequest{
		Amount:   15000, // $150.00
		Currency: "USD",
		ZoneID:   zone.ID, // Scope to our new zone
	})
	if err != nil {
		log.Fatalf("Failed to create intent: %v", err)
	}
	fmt.Printf("Created Payment Intent: %s. This should trigger an automated ledger provision.\n", intent.ID)

	// In a real scenario, the Flow Engine would now pick up the 'payment.succeeded' event
	// and execute nodes like 'Check Amount > 100' -> 'Provision Ledger'.

	fmt.Println("🚀 Automation example completed. Check Audit Logs in the Dashboard!")
}
