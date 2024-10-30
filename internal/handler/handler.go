package handler

import (
	"context"
	"fmt"

	"github.com/NicoNex/echotron/v3"
	"github.com/makarychev13/wallet/internal/model/brokerage"
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

type Handler struct {
	accountLister accountLister
	tg            echotron.API
}

func New(tg echotron.API, accountLister accountLister) *Handler {
	return &Handler{accountLister: accountLister, tg: tg}
}
