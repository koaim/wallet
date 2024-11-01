package deposits

import (
	"context"
	"fmt"

	"github.com/makarychev13/wallet/internal/model/deposit"
)

type depositsCreator interface {
	Create(ctx context.Context, d deposit.Deposit) error
}

type CreateUseCase struct {
	deposits depositsCreator
}

func NewCreateUseCase(d depositsCreator) *CreateUseCase {
	return &CreateUseCase{deposits: d}
}

func (u *CreateUseCase) Create(ctx context.Context, d deposit.Deposit) error {
	err := u.deposits.Create(ctx, d)
	if err != nil {
		return fmt.Errorf("create deposit: %w", err)
	}

	return nil
}
