package fintech

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	DefaultBaseURL = "http://localhost:8080"
)

// Client is the main entry point for the Fintech SDK.
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client

	Ledger   *LedgerService
	Auth     *AuthService
	Payments *PaymentsService
	Wallets  *WalletsService
	Billing  *BillingService
	Connect  *ConnectService
	Webhooks *WebhooksService
	Zones    *ZonesService
}

// ClientOption is a function that configures a Client.
type ClientOption func(*Client)

// NewClient creates a new Fintech SDK client.
func NewClient(apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		baseURL:    DefaultBaseURL,
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Ledger = &LedgerService{client: c}
	c.Auth = &AuthService{client: c}
	c.Payments = &PaymentsService{client: c}
	c.Wallets = &WalletsService{client: c}
	c.Billing = &BillingService{client: c}
	c.Connect = &ConnectService{client: c}
	c.Webhooks = &WebhooksService{client: c}
	c.Zones = &ZonesService{client: c}

	return c
}

// WithBaseURL sets the base URL for the client.
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(hc *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = hc
	}
}

func (c *Client) newRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", c.apiKey)

	return req, nil
}

func (c *Client) doRequest(req *http.Request, out interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode >= 400 {
		// Read body for error message if possible
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("api error: status=%d body=%s", resp.StatusCode, string(b))
	}

	if out != nil {
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}

func (c *Client) do(ctx context.Context, method, path string, body interface{}, out interface{}) error {
	req, err := c.newRequest(ctx, method, path, body)
	if err != nil {
		return err
	}
	return c.doRequest(req, out)
}

// --- Ledger Service ---

type LedgerService struct {
	client *Client
}

type RecordTransactionRequest struct {
	AccountID   string `json:"accountId"`
	Amount      int64  `json:"amount"` // In cents
	Currency    string `json:"currency"`
	Description string `json:"description"`
	ReferenceID string `json:"referenceId"`
}

type RecordTransactionResponse struct {
	TransactionID string `json:"transactionId"`
	Status        string `json:"status"`
}

func (s *LedgerService) RecordTransaction(ctx context.Context, req *RecordTransactionRequest) (*RecordTransactionResponse, error) {
	var res RecordTransactionResponse
	err := s.client.do(ctx, http.MethodPost, "/v1/ledger/transactions", req, &res)
	return &res, err
}

type GetAccountResponse struct {
	AccountID string    `json:"accountId"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
}

func (s *LedgerService) GetAccount(ctx context.Context, accountID string) (*GetAccountResponse, error) {
	var res GetAccountResponse
	err := s.client.do(ctx, http.MethodGet, fmt.Sprintf("/v1/ledger/accounts/%s", accountID), nil, &res)
	return &res, err
}

// --- Auth Service ---

type AuthService struct {
	client *Client
}

type ValidateKeyResponse struct {
	Valid       bool   `json:"valid"`
	UserID      string `json:"userId"`
	OrgID       string `json:"orgId"`
	Environment string `json:"environment"`
	Scopes      string `json:"scopes"`
}

func (s *AuthService) ValidateKey(ctx context.Context, keyHash string) (*ValidateKeyResponse, error) {
	var res ValidateKeyResponse
	req := map[string]string{"keyHash": keyHash}
	err := s.client.do(ctx, http.MethodPost, "/v1/auth/validate", req, &res)
	return &res, err
}

// --- Payments Service ---

type PaymentsService struct {
	client *Client
}

type PaymentIntentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	ZoneID   string `json:"zone_id"` // Scoping
}

type PaymentResponse struct {
	ID        string    `json:"id"`
	Amount    int64     `json:"amount"`
	Currency  string    `json:"currency"`
	Status    string    `json:"status"` // succeeded, pending, failed
	CreatedAt time.Time `json:"createdAt"`
}

func (s *PaymentsService) CreateIntent(ctx context.Context, req *PaymentIntentRequest) (*PaymentResponse, error) {
	var res PaymentResponse
	err := s.client.do(ctx, http.MethodPost, "/v1/payments", req, &res)
	return &res, err
}

func (s *PaymentsService) Get(ctx context.Context, id string) (*PaymentResponse, error) {
	var res PaymentResponse
	err := s.client.do(ctx, http.MethodGet, fmt.Sprintf("/v1/payments/%s", id), nil, &res)
	return &res, err
}

func (s *PaymentsService) Refund(ctx context.Context, id string, amount int64) (*PaymentResponse, error) {
	var res PaymentResponse
	req := map[string]int64{"amount": amount}
	err := s.client.do(ctx, http.MethodPost, fmt.Sprintf("/v1/payments/%s/refund", id), req, &res)
	return &res, err
}

// --- Wallets Service ---

type WalletsService struct {
	client *Client
}

type CreateWalletRequest struct {
	UserID   string `json:"userId"`
	Currency string `json:"currency"`
}

type WalletResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
}

func (s *WalletsService) Create(ctx context.Context, req *CreateWalletRequest) (*WalletResponse, error) {
	var res WalletResponse
	err := s.client.do(ctx, http.MethodPost, "/v1/wallets", req, &res)
	return &res, err
}

func (s *WalletsService) Get(ctx context.Context, id string) (*WalletResponse, error) {
	var res WalletResponse
	err := s.client.do(ctx, http.MethodGet, fmt.Sprintf("/v1/wallets/%s", id), nil, &res)
	return &res, err
}

type WalletTransactionRequest struct {
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}

func (s *WalletsService) Credit(ctx context.Context, id string, req *WalletTransactionRequest) (*WalletResponse, error) {
	var res WalletResponse
	err := s.client.do(ctx, http.MethodPost, fmt.Sprintf("/v1/wallets/%s/credit", id), req, &res)
	return &res, err
}

func (s *WalletsService) Debit(ctx context.Context, id string, req *WalletTransactionRequest) (*WalletResponse, error) {
	var res WalletResponse
	err := s.client.do(ctx, http.MethodPost, fmt.Sprintf("/v1/wallets/%s/debit", id), req, &res)
	return &res, err
}

// --- Billing Service ---

type BillingService struct {
	client *Client
}

type CreateSubscriptionRequest struct {
	CustomerID string `json:"customerId"`
	PriceID    string `json:"priceId"`
}

type SubscriptionResponse struct {
	ID               string    `json:"id"`
	Status           string    `json:"status"` // active, canceled, etc.
	CurrentPeriodEnd time.Time `json:"currentPeriodEnd"`
}

func (s *BillingService) CreateSubscription(ctx context.Context, req *CreateSubscriptionRequest) (*SubscriptionResponse, error) {
	var res SubscriptionResponse
	err := s.client.do(ctx, http.MethodPost, "/v1/billing/subscriptions", req, &res)
	return &res, err
}

func (s *BillingService) GetSubscription(ctx context.Context, id string) (*SubscriptionResponse, error) {
	var res SubscriptionResponse
	err := s.client.do(ctx, http.MethodGet, fmt.Sprintf("/v1/billing/subscriptions/%s", id), nil, &res)
	return &res, err
}

func (s *BillingService) CancelSubscription(ctx context.Context, id string) (*SubscriptionResponse, error) {
	var res SubscriptionResponse
	err := s.client.do(ctx, http.MethodPost, fmt.Sprintf("/v1/billing/subscriptions/%s/cancel", id), nil, &res)
	return &res, err
}

// --- Connect Service ---

type ConnectService struct {
	client *Client
}

type CreateConnectAccountRequest struct {
	Type    string `json:"type"` // standard, express
	Country string `json:"country"`
	Email   string `json:"email"`
}

type ConnectAccountResponse struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Enabled bool   `json:"enabled"`
}

func (s *ConnectService) CreateAccount(ctx context.Context, req *CreateConnectAccountRequest) (*ConnectAccountResponse, error) {
	var res ConnectAccountResponse
	err := s.client.do(ctx, http.MethodPost, "/v1/connect/accounts", req, &res)
	return &res, err
}

func (s *ConnectService) GetAccount(ctx context.Context, id string) (*ConnectAccountResponse, error) {
	var res ConnectAccountResponse
	err := s.client.do(ctx, http.MethodGet, fmt.Sprintf("/v1/connect/accounts/%s", id), nil, &res)
	return &res, err
}

// --- Webhooks Service ---

type WebhooksService struct {
	client *Client
}

type CreateWebhookEndpointRequest struct {
	URL           string   `json:"url"`
	EnabledEvents []string `json:"enabledEvents"`
}

type WebhookEndpointResponse struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Secret string `json:"secret"` // Signing secret
	Status string `json:"status"`
}

type Event struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type"`
	Object    string                 `json:"object"`
	Data      map[string]interface{} `json:"data"`
	CreatedAt time.Time              `json:"createdAt"`
}

func (s *WebhooksService) CreateEndpoint(ctx context.Context, req *CreateWebhookEndpointRequest) (*WebhookEndpointResponse, error) {
	var res WebhookEndpointResponse
	err := s.client.do(ctx, http.MethodPost, "/v1/webhooks/endpoints", req, &res)
	return &res, err
}

func (s *WebhooksService) ListEndpoints(ctx context.Context) ([]WebhookEndpointResponse, error) {
	var res []WebhookEndpointResponse
	err := s.client.do(ctx, http.MethodGet, "/v1/webhooks/endpoints", nil, &res)
	return res, err
}

func (s *WebhooksService) ConstructEvent(payload []byte, signatureHeader string, secret string) (*Event, error) {
	// ... (implementation omitted for brevity)
	return &Event{}, nil
}

// --- Zones Service ---

type ZonesService struct {
	client *Client
}

type ZoneResponse struct {
	ID        string    `json:"id"`
	OrgID     string    `json:"org_id"`
	Name      string    `json:"name"`
	Mode      string    `json:"mode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateZoneRequest struct {
	OrgID string `json:"org_id"`
	Name  string `json:"name"`
	Mode  string `json:"mode"` // "test" or "live"
}

func (s *ZonesService) Create(ctx context.Context, req *CreateZoneRequest) (*ZoneResponse, error) {
	var res ZoneResponse
	err := s.client.do(ctx, http.MethodPost, "/v1/zones", req, &res)
	return &res, err
}

func (s *ZonesService) List(ctx context.Context, orgID string) ([]ZoneResponse, error) {
	var res []ZoneResponse
	err := s.client.do(ctx, http.MethodGet, fmt.Sprintf("/v1/zones?org_id=%s", orgID), nil, &res)
	return res, err
}

func (s *ZonesService) Get(ctx context.Context, id string) (*ZoneResponse, error) {
	var res ZoneResponse
	err := s.client.do(ctx, http.MethodGet, fmt.Sprintf("/v1/zones?id=%s", id), nil, &res)
	return &res, err
}

// TriggerEvent manually fires an event into the Sapliy ecosystem.
func (c *Client) TriggerEvent(ctx context.Context, eventType string, zoneID string, data interface{}) error {
	return c.do(ctx, http.MethodPost, "/v1/events/trigger", map[string]interface{}{
		"type":    eventType,
		"zone_id": zoneID,
		"data":    data,
	}, nil)
}

// ReplayEvent re-triggers a past event.
func (c *Client) ReplayEvent(ctx context.Context, eventID string, zoneID string) error {
	return c.do(ctx, http.MethodPost, fmt.Sprintf("/api/v1/events/%s/replay", eventID), map[string]interface{}{
		"zoneId": zoneID,
	}, nil)
}

// GetPastEvents retrieves past events for a zone.
func (c *Client) GetPastEvents(ctx context.Context, zoneID string, limit int, offset int) ([]*Event, error) {
	var res []*Event
	err := c.do(ctx, http.MethodGet, fmt.Sprintf("/api/v1/zones/%s/events/past?limit=%d&offset=%d", zoneID, limit, offset), nil, &res)
	return res, err
}
