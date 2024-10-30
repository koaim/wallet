package brokerages

import (
	"context"
	"fmt"

	"github.com/makarychev13/wallet/internal/model/brokerage"
)

type accountsLister interface {
	Accounts(ctx context.Context) ([]brokerage.Account, error)
}

type ListUseCase struct {
	accounts accountsLister
}

func NewListUseCase(accounts accountsLister) *ListUseCase {
	return &ListUseCase{accounts: accounts}
}

func (u *ListUseCase) All(ctx context.Context) ([]brokerage.Account, error) {
	res, err := u.accounts.Accounts(ctx)
	if err != nil {
		return nil, fmt.Errorf("get accounts: %w", err)
	}

	return res, nil
}
