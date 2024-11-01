package handler

import (
	"context"
	"fmt"

	"github.com/NicoNex/echotron/v3"
	"github.com/makarychev13/wallet/internal/model/brokerage"
	"github.com/makarychev13/wallet/internal/model/deposit"
)

type accountLister interface {
	All(ctx context.Context) ([]brokerage.Account, error)
}

type depositLister interface {
	All(ctx context.Context) ([]deposit.Deposit, error)
}

type depositCreator interface {
	Create(ctx context.Context, d deposit.Deposit) error
}

type session interface {
	Set(id int64, key string, value interface{}) error
	Get(id int64, key string) (interface{}, error)
	Clear(id int64) error
	GetAll(id int64) (map[string]interface{}, error)
}

type states interface {
	Set(id int64, state State) error
}

type Handler struct {
	session session
	states  states
	tg      echotron.API

	accountLister  accountLister
	depositLister  depositLister
	depositCreator depositCreator
}

func New(
	tg echotron.API,
	accountLister accountLister,
	depositLister depositLister,
	depositCreator depositCreator,
	session session,
	states states,
) *Handler {
	return &Handler{
		accountLister:  accountLister,
		depositLister:  depositLister,
		depositCreator: depositCreator,
		tg:             tg,
		session:        session,
		states:         states,
	}
}

type ErrSendMsg struct {
	Code        int
	Description string
	ID          int64
}

func NewErrSendMsg(res echotron.APIResponseMessage, id int64) ErrSendMsg {
	return ErrSendMsg{
		Code:        res.ErrorCode,
		Description: res.Description,
		ID:          id,
	}
}

func (e ErrSendMsg) Error() string {
	return fmt.Sprintf("%v: id: %v, err: '%v'", e.Code, e.ID, e.Description)
}
