package handler

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/NicoNex/echotron/v3"
	"github.com/makarychev13/wallet/internal/model/brokerage"
)

func (h *Handler) BrokerageAccounts(msg echotron.Message) error {
	accounts, err := h.accountLister.All(context.Background())
	if errors.Is(err, brokerage.ErrNotFound) {
		res, err := h.tg.SendMessage("У вас нет брокерских счетов", msg.From.ID, nil)
		if err != nil {
			return NewErrSendMsg(res, msg.From.ID)
		}
	}
	if err != nil {
		return fmt.Errorf("get accounts: %w", err)
	}

	var sb strings.Builder
	for _, v := range accounts {
		sb.WriteString(v.Name)
		sb.WriteString("\n")
	}

	res, err := h.tg.SendMessage(sb.String(), msg.From.ID, &echotron.MessageOptions{
		ReplyMarkup: echotron.ReplyKeyboardRemove{
			RemoveKeyboard: true,
		},
	})
	if err != nil {
		return NewErrSendMsg(res, msg.From.ID)
	}

	return nil
}
