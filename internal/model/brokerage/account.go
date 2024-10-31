package brokerage

import "fmt"

var (
	ErrNotFound = fmt.Errorf("not found")
)

type Account struct {
	ID      int
	Name    string
	Balance float64
}
