# @sapliyio/fintech-go

[![Go Reference](https://pkg.go.dev/badge/github.com/sapliy/sapliy-sdk-go.svg)](https://pkg.go.dev/github.com/sapliy/sapliy-sdk-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Official Go SDK for the Sapliy Ecosystem. Build robust financial applications with type-safe, idiomatic Go.

## Features

- **Payments** — Create charges, handle refunds, manage payment lifecycle
- **Wallets** — User balances and internal accounting
- **Ledger** — Double-entry bookkeeping for high-integrity transactions
- **Billing** — Subscriptions and recurring billing
- **Connect** — Multi-tenant support and managed accounts
- **Webhooks** — Event handling with signature verification

## Installation

```bash
go get github.com/sapliy/sapliy-sdk-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    fintech "github.com/sapliy/sapliy-sdk-go"
)

func main() {
    client := sapliy.NewClient("sk_test_...")

    // Create a payment
    payment, err := client.Payments.Create(context.Background(), &sapliy.CreateChargeRequest{
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
client := sapliy.NewClient("sk_test_...",
    sapliy.WithBaseURL("https://api.yourdomain.com"),
)

// Custom HTTP client
client := sapliy.NewClient("sk_test_...",
    sapliy.WithHTTPClient(&http.Client{
        Timeout: 30 * time.Second,
    }),
)
```

## API Reference

### Payments

```go
// Create a charge
payment, err := client.Payments.Create(ctx, &sapliy.CreateChargeRequest{
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
wallet, err := client.Wallets.Create(ctx, &sapliy.CreateWalletRequest{
    Name:     "User Wallet",
    Currency: "USD",
})

// Get wallet balance
wallet, err := client.Wallets.Get(ctx, "wal_123")

// Credit (add funds)
wallet, err := client.Wallets.Credit(ctx, "wal_123", &sapliy.WalletCreditRequest{
    Amount:      1000,
    Description: "Deposit",
})

// Debit (withdraw funds)
wallet, err := client.Wallets.Debit(ctx, "wal_123", &sapliy.WalletDebitRequest{
    Amount:      500,
    Description: "Purchase",
})
```

### Ledger

```go
// Record a transaction
resp, err := client.Ledger.RecordTransaction(ctx, &sapliy.RecordTransactionRequest{
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
sub, err := client.Billing.CreateSubscription(ctx, &sapliy.CreateSubscriptionRequest{
    CustomerID: "cust_123",
    PlanID:     "plan_monthly",
})

// Get subscription
sub, err := client.Billing.GetSubscription(ctx, "sub_123")

// Cancel subscription
err := client.Billing.CancelSubscription(ctx, "sub_123")
```

### Events (Automation)

```go
// Emit an event to trigger flows
err := client.TriggerEvent(ctx, "checkout.completed", "zone_123", map[string]interface{}{
    "cartId":     "cart_123",
    "total":      5000,
    "customerId": "cust_456",
})
```

## Flow Builder Template Variables

When creating flows in the Sapliy Flow Builder, you can use **Handlebars template syntax** to dynamically reference event data in your automation logic. This is particularly useful for approval messages, webhook payloads, and conditional logic.

### Available Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `{{event.type}}` | The event type that triggered the flow | `payment.completed` |
| `{{event.id}}` | Unique event identifier | `evt_abc123` |
| `{{event.payload.*}}` | Access any field from the event payload | `{{event.payload.amount}}` |
| `{{event.createdAt}}` | Timestamp when event was created | `2024-01-15T10:30:00Z` |

### Usage Examples

#### Approval Node Message

```handlebars
Approval required for payment of ${{event.payload.amount}} {{event.payload.currency}}
Customer: {{event.payload.customerId}}
```

When you emit an event like:

```go
client.TriggerEvent(ctx, "payment.high_value", "zone_123", map[string]interface{}{
    "amount":     10000,
    "currency":   "USD",
    "customerId": "cust_456",
})
```

The approval message will render as:

```
Approval required for payment of $10000 USD
Customer: cust_456
```

#### Webhook Payload Template

```json
{
  "orderId": "{{event.payload.orderId}}",
  "status": "approved",
  "approvedAt": "{{event.createdAt}}",
  "amount": {{event.payload.amount}}
}
```

### Best Practices

1. **Always validate data exists**: Use the Flow Builder's test mode to ensure your template variables resolve correctly
2. **Type safety**: Numeric fields don't need quotes in JSON templates: `{{amount}}` not `"{{amount}}"`
3. **Nested objects**: Access nested data with dot notation: `{{event.payload.customer.email}}`
4. **Debugging**: Use `sapliy listen` to see the actual event payloads and verify your template paths

## Webhook Handling

```go
package main

import (
    "net/http"

    fintech "github.com/sapliy/sapliy-sdk-go"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
    client := sapliy.NewClient("sk_test_...")
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
    // err contains: "Financial Automation API Error (404): Payment not found"
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

## Part of Sapliy Ecosystem

- [fintech-ecosystem](https://github.com/Sapliy/fintech-ecosystem) — Core backend
- [sapliy-sdk-node](https://github.com/Sapliy/sapliy-sdk-node) — Node.js SDK
- [sapliy-sdk-python](https://github.com/Sapliy/sapliy-sdk-python) — Python SDK

## License

MIT © [Sapliy](https://github.com/sapliy)
