package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sapliy/fintech-sdk-go"
	"github.com/sapliy/fintech-sdk-go/generated"
)

func main() {
	client := fintech.NewSapliyClient("sk_test_123", "http://localhost:8080")
	ctx := context.Background()

	fmt.Println("--- Payments Example ---")
	paymentReq := generated.PaymentsCreatePaymentIntentRequest{
		Amount:      generated.PtrString("5000"), // $50.00
		Currency:    generated.PtrString("USD"),
		Description: generated.PtrString("Example Payment for Go SDK"),
	}

	payment, _, err := client.Payments.PaymentServiceCreatePaymentIntent(ctx).Body(paymentReq).Execute()
	if err != nil {
		log.Fatalf("Error creating payment: %v", err)
	}
	fmt.Printf("Created Payment Intent: %s\n", *payment.Id)

	confirmReq := generated.PaymentServiceConfirmPaymentIntentBody{
		PaymentMethodId: generated.PtrString("pm_card_visa"),
	}
	_, _, err = client.Payments.PaymentServiceConfirmPaymentIntent(ctx, *payment.Id).Body(confirmReq).Execute()
	if err != nil {
		log.Fatalf("Error confirming payment: %v", err)
	}
	fmt.Println("Confirmed Payment Intent!")

	fmt.Println("\n--- Ledger Example ---")
	balance, _, err := client.Ledger.LedgerServiceGetAccount(ctx, "acc_123").Execute()
	if err != nil {
		log.Fatalf("Error getting balance: %v", err)
	}
	fmt.Printf("Account Balance: %s %s\n", *balance.Balance, *balance.Currency)
}
