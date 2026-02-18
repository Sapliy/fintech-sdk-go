package fintech

import (
	"context"

	"github.com/sapliy/fintech-sdk-go/generated"
)

// AuthService handles authentication and key validation.
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
