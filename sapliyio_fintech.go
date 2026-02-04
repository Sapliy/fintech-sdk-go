package fintech

import (
	"strings"

	"github.com/sapliy/fintech-sdk-go/generated"
)

type SapliyClient struct {
	Auth          *generated.AuthServiceAPIService
	Billing       *generated.BillingServiceAPIService
	Ledger        *generated.LedgerServiceAPIService
	Notifications *generated.NotificationServiceAPIService
	Payments      *generated.PaymentServiceAPIService
	Wallets       *generated.WalletServiceAPIService
	Flows         *generated.FlowServiceAPIService
	Zones         *generated.ZoneServiceAPIService
	Config        *generated.Configuration
}

func NewSapliyClient(apiKey string, baseURL string) *SapliyClient {
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	cfg := generated.NewConfiguration()
	cfg.Host = strings.TrimPrefix(strings.TrimPrefix(baseURL, "http://"), "https://")
	cfg.Scheme = "http"
	if strings.HasPrefix(baseURL, "https") {
		cfg.Scheme = "https"
	}
	cfg.AddDefaultHeader("X-API-Key", apiKey)

	apiClient := generated.NewAPIClient(cfg)

	return &SapliyClient{
		Auth:          apiClient.AuthServiceAPI,
		Billing:       apiClient.BillingServiceAPI,
		Ledger:        apiClient.LedgerServiceAPI,
		Notifications: apiClient.NotificationServiceAPI,
		Payments:      apiClient.PaymentServiceAPI,
		Wallets:       apiClient.WalletServiceAPI,
		Flows:         apiClient.FlowServiceAPI,
		Zones:         apiClient.ZoneServiceAPI,
		Config:        cfg,
	}
}
