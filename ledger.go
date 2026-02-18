package fintech

import (
	"context"

	"github.com/sapliy/fintech-sdk-go/generated"
)

// LedgerService handles high-integrity accounting and ledger accounts.
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

func (s *LedgerService) GetAccount(ctx context.Context, zoneID string, id string) (*generated.LedgerAccount, error) {
	acc, _, err := s.c.gen.LedgerAPI.GetLedgerAccount(ctx, id).
		XZoneID(zoneID).
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

func (s *LedgerService) GetTransaction(ctx context.Context, zoneID string, id string) (*generated.LedgerTransaction, error) {
	tx, _, err := s.c.gen.LedgerAPI.GetLedgerTransaction(ctx, id).
		XZoneID(zoneID).
		Execute()
	return tx, err
}
