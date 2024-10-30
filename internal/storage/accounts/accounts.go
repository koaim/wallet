package accounts

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/makarychev13/wallet/internal/model/brokerage"
)

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) All(ctx context.Context) ([]brokerage.Account, error) {
	query, args, err :=
		squirrel.
			Select("id", "name", "balance").
			From("accounts").
			OrderBy("id").
			ToSql()
	if err != nil {
		return nil, err
	}

	var accounts []brokerage.Account
	err = r.db.SelectContext(ctx, &accounts, query, args...)
	if err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, brokerage.ErrNotFound
	}

	return accounts, nil
}
