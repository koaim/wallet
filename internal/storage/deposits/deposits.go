package deposits

import (
	"context"
	"fmt"

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
			Select("id", "name", "balance", "rate", "month_period", "created_at", "closed_at").
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

func (r *Repo) Create(ctx context.Context, d deposit.Deposit) error {
	query, args, err :=
		squirrel.
			Insert("deposits").
			Columns("name", "balance", "rate", "month_period", "created_at").
			Values(d.Name, d.Balance, d.Rate, d.MonthPeriod, squirrel.Expr("now()")).
			PlaceholderFormat(squirrel.Dollar).
			ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec sql: %w", err)
	}

	return nil
}
