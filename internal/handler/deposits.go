package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/NicoNex/echotron/v3"
	"github.com/makarychev13/wallet/internal/handler/button"
	"github.com/makarychev13/wallet/internal/model/deposit"
)

func (h *Handler) ListDeposits(msg echotron.Message) error {
	_, err := h.depositLister.All(context.Background())
	if errors.Is(err, deposit.ErrNotFound) {
		res, err := h.tg.SendMessage("Вкладов нет. Добавить?", msg.From.ID, &echotron.MessageOptions{
			ReplyMarkup: button.AddDeposit,
		})
		if err != nil {
			return NewErrSendMsg(res, msg.From.ID)
		}

		return nil
	}
	if err != nil {
		return fmt.Errorf("list deposits: %w", err)
	}

	return nil
}
