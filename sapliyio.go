package sapliy

import (
	"net/http"
	"time"

	"github.com/sapliy/sapliy-sdk-go/generated"
)

const (
	DefaultBaseURL = "http://localhost:8080"
)

// Client is the main entry point for the Sapliy Go SDK.
// It provides access to all specialized services via high-level, idiomatic wrappers.
type Client struct {
	gen *generated.APIClient

	Auth          *AuthService
	Billing       *BillingService
	Ledger        *LedgerService
	Notifications *NotificationsService
	Payments      *PaymentsService
	Wallets       *WalletsService
	Flows         *FlowsService
	Zones         *ZonesService
	Events        *EventsService
	Playbooks     *PlaybooksService
}

// NewClient creates a new Financial Automation SDK client.
func NewClient(apiKey string, opts ...ClientOption) *Client {
	cfg := generated.NewConfiguration()
	cfg.Servers = generated.ServerConfigurations{
		{
			URL: DefaultBaseURL,
		},
	}
	cfg.AddDefaultHeader("Authorization", "Bearer "+apiKey)
	cfg.AddDefaultHeader("X-API-Key", apiKey) // Support both standard and custom headers

	for _, opt := range opts {
		opt(cfg)
	}

	gen := generated.NewAPIClient(cfg)

	c := &Client{
		gen: gen,
	}

	c.Auth = &AuthService{c: c}
	c.Billing = &BillingService{c: c}
	c.Ledger = &LedgerService{c: c}
	c.Notifications = &NotificationsService{c: c}
	c.Payments = &PaymentsService{c: c}
	c.Wallets = &WalletsService{c: c}
	c.Flows = &FlowsService{c: c}
	c.Zones = &ZonesService{c: c}
	c.Events = &EventsService{c: c}
	c.Playbooks = &PlaybooksService{c: c}

	return c
}

// ClientOption is a function that configures a Client.
type ClientOption func(*generated.Configuration)

// WithBaseURL sets the base URL for the client.
func WithBaseURL(url string) ClientOption {
	return func(cfg *generated.Configuration) {
		cfg.Servers = generated.ServerConfigurations{
			{
				URL: url,
			},
		}
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(hc *http.Client) ClientOption {
	return func(cfg *generated.Configuration) {
		cfg.HTTPClient = hc
	}
}

// WithTimeout sets a custom timeout for the default HTTP client.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(cfg *generated.Configuration) {
		if cfg.HTTPClient == nil {
			cfg.HTTPClient = &http.Client{}
		}
		cfg.HTTPClient.Timeout = timeout
	}
}
