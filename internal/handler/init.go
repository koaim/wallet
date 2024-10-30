package handler

import (
	"github.com/NicoNex/echotron/v3"
	"github.com/makarychev13/wallet/internal/handler/button"
	"github.com/makarychev13/wallet/internal/handler/message"
)

func (h *Handler) Init(msg echotron.Message) error {
	_, err := h.tg.SendMessage(message.Welcome, msg.From.ID, &echotron.MessageOptions{
		ReplyMarkup: button.Savings,
	})
	if err != nil {
		return err
	}

	return nil
}
