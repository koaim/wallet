package deposit

import (
	"fmt"
	"time"
)

var (
	ErrNotFound = fmt.Errorf("deposits not found")
)

type Deposit struct {
	ID          int
	Name        string
	Balance     float64
	Rate        float64
	MonthPeriod int        `db:"month_period"`
	CreatedAt   time.Time  `db:"created_at"`
	ClosedAt    *time.Time `db:"closed_at"`
}
