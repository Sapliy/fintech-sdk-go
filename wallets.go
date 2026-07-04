package sapliy

import (
	"context"

	"github.com/sapliy/sapliy-sdk-go/generated"
)

// WalletsService manages user wallets and internal currency transfers.
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
	res, _, err := s.c.gen.WalletsAPI.TopupWallet(ctx).
		XZoneID(zoneID).
		TopupWalletRequest(generated.TopupWalletRequest{
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
	res, _, err := s.c.gen.WalletsAPI.TransferWallet(ctx).
		XZoneID(zoneID).
		TransferWalletRequest(generated.TransferWalletRequest{
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
