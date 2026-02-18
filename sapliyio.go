package fintech

import (
	"context"
	"net/http"

	"github.com/sapliy/fintech-sdk-go/generated"
)

const (
	DefaultBaseURL = "http://localhost:8080"
)

// Client is the main entry point for the Fintech SDK.
type Client struct {
	gen *generated.APIClient

	Ledger   *LedgerService
	Auth     *AuthService
	Payments *PaymentsService
	Wallets  *WalletsService
	Billing  *BillingService
	Events   *EventsService
	Zones    *ZonesService
}

// ClientOption is a function that configures a Client.
type ClientOption func(*generated.Configuration)

// NewClient creates a new Fintech SDK client.
func NewClient(apiKey string, opts ...ClientOption) *Client {
	cfg := generated.NewConfiguration()
	cfg.Servers = generated.ServerConfigurations{
		{
			URL: DefaultBaseURL,
		},
	}
	cfg.AddDefaultHeader("Authorization", "Bearer "+apiKey)

	for _, opt := range opts {
		opt(cfg)
	}

	gen := generated.NewAPIClient(cfg)

	c := &Client{
		gen: gen,
	}

	c.Ledger = &LedgerService{c: c}
	c.Auth = &AuthService{c: c}
	c.Payments = &PaymentsService{c: c}
	c.Wallets = &WalletsService{c: c}
	c.Billing = &BillingService{c: c}
	c.Events = &EventsService{c: c}
	c.Zones = &ZonesService{c: c}

	return c
}

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

// --- Payments Service ---

type PaymentsService struct {
	c *Client
}

type CreateChargeRequest struct {
	Amount      int64             `json:"amount"`
	Currency    string            `json:"currency"`
	Description string            `json:"description"`
	Metadata    map[string]string `json:"metadata"`
	ZoneID      string            `json:"zone_id"`
}

func (s *PaymentsService) Create(ctx context.Context, req *CreateChargeRequest) (*generated.PaymentIntent, error) {
	genReq := s.c.gen.PaymentsAPI.CreatePaymentIntent(ctx).
		CreatePaymentIntentRequest(generated.CreatePaymentIntentRequest{
			Amount:      req.Amount,
			Currency:    req.Currency,
			Description: &req.Description,
			Metadata:    req.Metadata,
		}).
		XZoneID(req.ZoneID)

	intent, _, err := genReq.Execute()
	return intent, err
}

func (s *PaymentsService) Get(ctx context.Context, id string, zoneID string) (*generated.PaymentIntent, error) {
	intent, _, err := s.c.gen.PaymentsAPI.GetPaymentIntent(ctx, id).
		XZoneID(zoneID).
		Execute()
	return intent, err
}

func (s *PaymentsService) Confirm(ctx context.Context, id string, zoneID string, paymentMethodID string) (*generated.PaymentIntent, error) {
	intent, _, err := s.c.gen.PaymentsAPI.ConfirmPaymentIntent(ctx, id).
		XZoneID(zoneID).
		ConfirmPaymentIntentRequest(generated.ConfirmPaymentIntentRequest{
			PaymentMethodId: &paymentMethodID,
		}).
		Execute()
	return intent, err
}

// --- Wallets Service ---

type WalletsService struct {
	c *Client
}

func (s *WalletsService) Get(ctx context.Context, userID string, zoneID string) (*generated.Wallet, error) {
	wallet, _, err := s.c.gen.WalletsAPI.GetWallet(ctx, userID).
		XZoneID(zoneID).
		Execute()
	return wallet, err
}

func (s *WalletsService) Topup(ctx context.Context, zoneID string, amount int64, currency string, referenceID string) (string, error) {
	res, _, err := s.c.gen.WalletsAPI.V1WalletsTopupPost(ctx).
		XZoneID(zoneID).
		V1WalletsTopupPostRequest(generated.V1WalletsTopupPostRequest{
			Amount:      amount,
			Currency:    currency,
			ReferenceId: referenceID,
		}).
		Execute()
	if err != nil {
		return "", err
	}
	return res.GetTransactionId(), nil
}

func (s *WalletsService) Transfer(ctx context.Context, zoneID string, toUserID string, amount int64, currency string, referenceID string) (string, error) {
	res, _, err := s.c.gen.WalletsAPI.V1WalletsTransferPost(ctx).
		XZoneID(zoneID).
		V1WalletsTransferPostRequest(generated.V1WalletsTransferPostRequest{
			ToUserId:    toUserID,
			Amount:      amount,
			Currency:    currency,
			ReferenceId: referenceID,
		}).
		Execute()
	if err != nil {
		return "", err
	}
	return res.GetTransactionId(), nil
}

// --- Ledger Service ---

type LedgerService struct {
	c *Client
}

func (s *LedgerService) CreateAccount(ctx context.Context, zoneID string, name string, accountType string, currency string) (*generated.LedgerAccount, error) {
	acc, _, err := s.c.gen.LedgerAPI.V1LedgerAccountsPost(ctx).
		XZoneID(zoneID).
		V1LedgerAccountsPostRequest(generated.V1LedgerAccountsPostRequest{
			Name:     name,
			Type:     accountType,
			Currency: &currency,
		}).
		Execute()
	return acc, err
}

func (s *LedgerService) RecordTransaction(ctx context.Context, zoneID string, referenceID string, description string, entries []generated.LedgerEntry) (string, error) {
	res, _, err := s.c.gen.LedgerAPI.V1LedgerTransactionsPost(ctx).
		XZoneID(zoneID).
		V1LedgerTransactionsPostRequest(generated.V1LedgerTransactionsPostRequest{
			ReferenceId: referenceID,
			Description: &description,
			Entries:     entries,
		}).
		Execute()
	if err != nil {
		return "", err
	}
	return res.GetStatus(), nil
}

// --- Auth Service ---

type AuthService struct {
	c *Client
}

func (s *AuthService) ValidateKey(ctx context.Context, key string) (bool, error) {
	res, _, err := s.c.gen.AuthAPI.ValidateKey(ctx).
		ValidateKeyRequest(generated.ValidateKeyRequest{
			Key: key,
		}).
		Execute()
	if err != nil {
		return false, err
	}
	return res.GetValid(), nil
}

// --- Zones Service ---

type ZonesService struct {
	c *Client
}

func (s *ZonesService) Create(ctx context.Context, orgID string, name string, mode string) (string, error) {
	res, _, err := s.c.gen.ZonesAPI.CreateZone(ctx).
		CreateZoneRequest(generated.CreateZoneRequest{
			OrgId: orgID,
			Name:  name,
			Mode:  mode,
		}).
		Execute()
	if err != nil {
		return "", err
	}
	return res.GetId(), nil
}

// --- Events Service ---

type EventsService struct {
	c *Client
}

func (s *EventsService) Emit(ctx context.Context, eventType string, data map[string]interface{}) (string, error) {
	res, _, err := s.c.gen.EventsAPI.EmitEvent(ctx).
		EmitEventRequest(generated.EmitEventRequest{
			Type: eventType,
			Data: data,
		}).
		Execute()
	if err != nil {
		return "", err
	}
	return res.GetEventId(), nil
}
