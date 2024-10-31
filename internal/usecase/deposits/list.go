package deposits

import (
	"context"
	"fmt"

	"github.com/makarychev13/wallet/internal/model/deposit"
)

type depositLister interface {
	All(ctx context.Context) ([]deposit.Deposit, error)
}

type ListUseCase struct {
	deposits depositLister
}

func NewListUseCase(deposits depositLister) *ListUseCase {
	return &ListUseCase{deposits: deposits}
}

func (u *ListUseCase) All(ctx context.Context) ([]deposit.Deposit, error) {
	res, err := u.deposits.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("get deposits: %w", err)
	}

	return res, nil
}
