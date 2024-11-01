package message

import (
	"fmt"
	"strings"

	"github.com/makarychev13/wallet/internal/model/deposit"
)

const (
	BrokerageAccounts = "Брокерские счета"
	Deposits          = "Вклады"
	SavingAccounts    = "Накопительные счета"
	Cards             = "Карты"
	Welcome           = "Добро пожаловать"
	Report            = "Отчёт"
	Start             = "/start"
	AddDeposit        = "Добавить вклад"
)

func DepositsList(deposits []deposit.Deposit) string {
	var sb strings.Builder

	var sum float64
	for _, v := range deposits {
		s := fmt.Sprintf("<b>%v</b> - %vр. (%v%% - %v мес.)", v.Name, v.Balance, v.Rate, v.MonthPeriod)
		sb.WriteString(s)
		sb.WriteString("\n")

		sum += v.Balance
	}

	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("<b>Всего</b> - %vр.", sum))

	return sb.String()
}
