package handler

import (
	"context"
	"fmt"

	"github.com/NicoNex/echotron/v3"
	"github.com/makarychev13/wallet/internal/model/brokerage"
	"github.com/makarychev13/wallet/internal/model/deposit"
)

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

type accountLister interface {
	All(ctx context.Context) ([]brokerage.Account, error)
}

type depositLister interface {
	All(ctx context.Context) ([]deposit.Deposit, error)
}

type Handler struct {
	accountLister accountLister
	depositLister depositLister
	tg            echotron.API
}

func New(tg echotron.API, accountLister accountLister, depositLister depositLister) *Handler {
	return &Handler{accountLister: accountLister, depositLister: depositLister, tg: tg}
}
