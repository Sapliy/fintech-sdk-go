# @sapliyio/fintech-go

[![Go Reference](https://pkg.go.dev/badge/github.com/sapliy/fintech-sdk-go.svg)](https://pkg.go.dev/github.com/sapliy/fintech-sdk-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Official Go SDK for the Sapliy Fintech Ecosystem. Build robust financial applications with type-safe, idiomatic Go.

## Features

- **Payments** — Create charges, handle refunds, manage payment lifecycle
- **Wallets** — User balances and internal accounting
- **Ledger** — Double-entry bookkeeping for high-integrity transactions
- **Billing** — Subscriptions and recurring billing
- **Connect** — Multi-tenant support and managed accounts
- **Webhooks** — Event handling with signature verification

## Installation

```bash
go get github.com/sapliy/fintech-sdk-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    fintech "github.com/sapliy/fintech-sdk-go"
)

func main() {
    client := fintech.NewClient("sk_test_...")

    // Create a payment
    payment, err := client.Payments.Create(context.Background(), &fintech.CreateChargeRequest{
        Amount:      2000, // $20.00
        Currency:    "USD",
        SourceID:    "src_123",
        Description: "Order #1234",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Payment created: %s\n", payment.ID)
}
```

## Configuration

```go
// Custom base URL (for self-hosted)
client := fintech.NewClient("sk_test_...",
    fintech.WithBaseURL("https://api.yourdomain.com"),
)

// Custom HTTP client
client := fintech.NewClient("sk_test_...",
    fintech.WithHTTPClient(&http.Client{
        Timeout: 30 * time.Second,
    }),
)
```

## API Reference

### Payments

```go
// Create a charge
payment, err := client.Payments.Create(ctx, &fintech.CreateChargeRequest{
    Amount:      1000,
    Currency:    "USD",
    SourceID:    "src_123",
    Description: "Coffee",
})

// Get payment details
payment, err := client.Payments.Get(ctx, "pay_123")

// Refund a payment
payment, err := client.Payments.Refund(ctx, "pay_123", 500) // partial refund
```

### Wallets

```go
// Create a wallet
wallet, err := client.Wallets.Create(ctx, &fintech.CreateWalletRequest{
    Name:     "User Wallet",
    Currency: "USD",
})

// Get wallet balance
wallet, err := client.Wallets.Get(ctx, "wal_123")

// Credit (add funds)
wallet, err := client.Wallets.Credit(ctx, "wal_123", &fintech.WalletCreditRequest{
    Amount:      1000,
    Description: "Deposit",
})

// Debit (withdraw funds)
wallet, err := client.Wallets.Debit(ctx, "wal_123", &fintech.WalletDebitRequest{
    Amount:      500,
    Description: "Purchase",
})
```

### Ledger

```go
// Record a transaction
resp, err := client.Ledger.RecordTransaction(ctx, &fintech.RecordTransactionRequest{
    AccountID:   "acc_123",
    Amount:      1000,
    Currency:    "USD",
    Description: "Payment received",
    ReferenceID: "ref_456",
})

// Get account details
account, err := client.Ledger.GetAccount(ctx, "acc_123")
```

### Billing

```go
// Create a subscription
sub, err := client.Billing.CreateSubscription(ctx, &fintech.CreateSubscriptionRequest{
    CustomerID: "cust_123",
    PlanID:     "plan_monthly",
})

// Get subscription
sub, err := client.Billing.GetSubscription(ctx, "sub_123")

// Cancel subscription
err := client.Billing.CancelSubscription(ctx, "sub_123")
```

## Webhook Handling

```go
package main

import (
    "net/http"

    fintech "github.com/sapliy/fintech-sdk-go"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
    client := fintech.NewClient("sk_test_...")
    secret := "whsec_..."

    payload, _ := io.ReadAll(r.Body)
    signature := r.Header.Get("X-Sapliy-Signature")

    event, err := client.Webhooks.ConstructEvent(payload, signature, secret)
    if err != nil {
        http.Error(w, "Invalid signature", http.StatusBadRequest)
        return
    }

    switch event.Type {
    case "payment.succeeded":
        // Handle successful payment
    case "payment.failed":
        // Handle failed payment
    }

    w.WriteHeader(http.StatusOK)
}
```

## Error Handling

```go
payment, err := client.Payments.Get(ctx, "invalid_id")
if err != nil {
    // err contains: "Fintech API Error (404): Payment not found"
    log.Printf("Error: %v", err)
}
```

## Context Support

All methods accept a `context.Context` for cancellation and timeouts:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

payment, err := client.Payments.Create(ctx, req)
```

## Part of Sapliy Fintech Ecosystem

- [fintech-ecosystem](https://github.com/Sapliy/fintech-ecosystem) — Core backend
- [fintech-sdk-node](https://github.com/Sapliy/fintech-sdk-node) — Node.js SDK
- [fintech-sdk-python](https://github.com/Sapliy/fintech-sdk-python) — Python SDK

## License

MIT © [Sapliy](https://github.com/sapliy)
