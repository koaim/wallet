package handler

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/NicoNex/echotron/v3"
	"github.com/makarychev13/wallet/internal/handler/button"
	"github.com/makarychev13/wallet/internal/handler/message"
	"github.com/makarychev13/wallet/internal/model/deposit"
)

const (
	sessionDepositKey = "deposit"
)

func (h *Handler) ListDeposits(msg echotron.Message) error {
	deposits, err := h.depositLister.All(context.Background())
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

	reply := message.DepositsList(deposits)

	res, err := h.tg.SendMessage(reply, msg.From.ID, &echotron.MessageOptions{
		ParseMode: echotron.HTML,
	})
	if err != nil {
		return NewErrSendMsg(res, msg.From.ID)
	}

	return nil
}

func (h *Handler) InitDepositCreating(msg echotron.Message) error {
	err := h.states.Set(msg.From.ID, WaitDepositName)
	if err != nil {
		return fmt.Errorf("set state %v: %w", WaitDepositName, err)
	}

	res, err := h.tg.SendMessage("Введите имя вклада", msg.From.ID, &echotron.MessageOptions{
		ReplyMarkup: echotron.ReplyKeyboardRemove{
			RemoveKeyboard: true,
		},
	})
	if err != nil {
		return NewErrSendMsg(res, msg.From.ID)
	}

	return nil
}

func (h *Handler) SetDepositName(msg echotron.Message) error {
	d := deposit.Deposit{
		Name: strings.TrimSpace(msg.Text),
	}

	err := h.session.Set(msg.From.ID, sessionDepositKey, d)
	if err != nil {
		return fmt.Errorf("set deposit name: %w", err)
	}

	err = h.states.Set(msg.From.ID, WaitDepositRate)
	if err != nil {
		return fmt.Errorf("set state %v: %w", WaitDepositName, err)
	}

	res, err := h.tg.SendMessage("Введите процентную ставку", msg.From.ID, &echotron.MessageOptions{})
	if err != nil {
		return NewErrSendMsg(res, msg.From.ID)
	}

	return nil
}

func (h *Handler) SetDepositRate(msg echotron.Message) error {
	v, err := h.session.Get(msg.From.ID, sessionDepositKey)
	if err != nil {
		return fmt.Errorf("get deposit from session: %w", err)
	}

	rate, err := strconv.ParseFloat(strings.TrimSpace(msg.Text), 64)
	if err != nil {
		res, err := h.tg.SendMessage("Неправильный формат данных. Введите процентную ставку ещё раз", msg.From.ID, nil)
		if err != nil {
			return NewErrSendMsg(res, msg.From.ID)
		}

		return nil
	}

	d := v.(deposit.Deposit)
	d.Rate = rate

	err = h.session.Set(msg.From.ID, sessionDepositKey, d)
	if err != nil {
		return fmt.Errorf("set deposit rate: %w", err)
	}

	err = h.states.Set(msg.From.ID, WaitDepositPeriod)
	if err != nil {
		return fmt.Errorf("set state %v: %w", WaitDepositPeriod, err)
	}

	res, err := h.tg.SendMessage("Введите срок вклада (в месяцах)", msg.From.ID, nil)
	if err != nil {
		return NewErrSendMsg(res, msg.From.ID)
	}

	return nil
}

func (h *Handler) SetDepositPeriod(msg echotron.Message) error {
	v, err := h.session.Get(msg.From.ID, sessionDepositKey)
	if err != nil {
		return fmt.Errorf("get deposit from session: %w", err)
	}

	period, err := strconv.Atoi(strings.TrimSpace(msg.Text))
	if err != nil {
		res, err := h.tg.SendMessage("Неправильный формат данных. Введите период ещё раз", msg.From.ID, nil)
		if err != nil {
			return NewErrSendMsg(res, msg.From.ID)
		}

		return nil
	}

	d := v.(deposit.Deposit)
	d.MonthPeriod = period

	err = h.session.Set(msg.From.ID, sessionDepositKey, d)
	if err != nil {
		return fmt.Errorf("set deposit period: %w", err)
	}

	err = h.states.Set(msg.From.ID, WaitDepositBalance)
	if err != nil {
		return fmt.Errorf("set state %v: %w", WaitDepositPeriod, err)
	}

	res, err := h.tg.SendMessage("Введите сумму вклада", msg.From.ID, nil)
	if err != nil {
		return NewErrSendMsg(res, msg.From.ID)
	}

	return nil
}

func (h *Handler) CreateDeposit(msg echotron.Message) error {
	v, err := h.session.Get(msg.From.ID, sessionDepositKey)
	if err != nil {
		return fmt.Errorf("get deposit from session: %w", err)
	}

	balance, err := strconv.ParseFloat(strings.TrimSpace(msg.Text), 64)
	if err != nil {
		res, err := h.tg.SendMessage("Неправильный формат данных. Введите процентную ставку ещё раз", msg.From.ID, nil)
		if err != nil {
			return NewErrSendMsg(res, msg.From.ID)
		}

		return nil
	}

	d := v.(deposit.Deposit)
	d.Balance = balance

	err = h.depositCreator.Create(context.Background(), d)
	if err != nil {
		return fmt.Errorf("create deposit: %w", err)
	}

	res, err := h.tg.SendMessage("Вклад создан", msg.From.ID, &echotron.MessageOptions{
		ReplyMarkup: button.Savings,
	})
	if err != nil {
		return NewErrSendMsg(res, msg.From.ID)
	}

	err = h.states.Set(msg.From.ID, InitState)
	if err != nil {
		return fmt.Errorf("set state %v: %w", InitState, err)
	}

	return nil
}
