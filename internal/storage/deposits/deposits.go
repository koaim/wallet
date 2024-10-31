package deposits

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/makarychev13/wallet/internal/model/deposit"
)

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) All(ctx context.Context) ([]deposit.Deposit, error) {
	query, args, err :=
		squirrel.
			Select("id", "name", "balance", "rate", "created_at", "closed_at").
			From("deposits").
			OrderBy("id").
			ToSql()
	if err != nil {
		return nil, err
	}

	var deposits []deposit.Deposit
	err = r.db.SelectContext(ctx, &deposits, query, args...)
	if err != nil {
		return nil, err
	}

	if len(deposits) == 0 {
		return nil, deposit.ErrNotFound
	}

	return deposits, nil
}
