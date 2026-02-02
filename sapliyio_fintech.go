package fintech

import (
	"fmt"

	"github.com/sapliy/fintech-sdk-go/generated"
)

type SapliyClient struct {
	Auth          *generated.AuthServiceAPIService
	Billing       *generated.BillingServiceAPIService
	Ledger        *generated.LedgerServiceAPIService
	Notifications *generated.NotificationServiceAPIService
	Payments      *generated.PaymentServiceAPIService
	Wallets       *generated.WalletServiceAPIService
	Config        *generated.Configuration
}

func NewSapliyClient(apiKey string, baseURL string) *SapliyClient {
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	cfg := generated.NewConfiguration()
	cfg.Host = baseURL
	cfg.Scheme = "http"
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	apiClient := generated.NewAPIClient(cfg)

	return &SapliyClient{
		Auth:          apiClient.AuthServiceAPI,
		Billing:       apiClient.BillingServiceAPI,
		Ledger:        apiClient.LedgerServiceAPI,
		Notifications: apiClient.NotificationServiceAPI,
		Payments:      apiClient.PaymentServiceAPI,
		Wallets:       apiClient.WalletServiceAPI,
		Config:        cfg,
	}
}
