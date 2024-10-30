package button

import (
	"github.com/NicoNex/echotron/v3"
	"github.com/makarychev13/wallet/internal/handler/message"
)

var (
	Savings = echotron.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]echotron.KeyboardButton{
			{
				{Text: message.BrokerageAccounts},
				{Text: message.Deposits},
				{Text: message.SavingAccounts},
				{Text: message.Cards},
			},
			{
				{Text: message.Report},
			},
		},
	}
)
