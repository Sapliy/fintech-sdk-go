package payouts

import (
	"context"
	"fmt"
	"log"

	"github.com/sapliy/fintech-sdk-go"
)

func main() {
	client := fintech.NewClient("sk_test_123", fintech.WithBaseURL("http://localhost:8080"))
	ctx := context.Background()

	fmt.Println("--- Payments Example ---")
	zoneID := "zone_test_123"
	payment, err := client.Payments.CreateIntent(ctx, &fintech.CreatePaymentRequest{
		Amount:      5000,
		Currency:    "USD",
		Description: "Example Payment for Go SDK",
		ZoneID:      zoneID,
	})
	if err != nil {
		log.Fatalf("Error creating payment: %v", err)
	}
	fmt.Printf("Created Payment Intent: %s\n", payment.GetId())

	// ConfirmIntent(ctx, id, zoneID, paymentMethodID)
	_, err = client.Payments.ConfirmIntent(ctx, payment.GetId(), zoneID, "pm_card_visa")
	if err != nil {
		log.Fatalf("Error confirming payment: %v", err)
	}
	fmt.Println("Confirmed Payment Intent!")

	fmt.Println("\n--- Ledger Example ---")
	// GetAccount requires zoneID + accountID
	balance, err := client.Ledger.GetAccount(ctx, zoneID, "acc_123")
	if err != nil {
		log.Fatalf("Error getting balance: %v", err)
	}
	fmt.Printf("Account Balance: %d %s\n", balance.GetBalance(), balance.GetCurrency())
}
