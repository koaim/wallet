package main

import (
	"log/slog"

	"github.com/BurntSushi/toml"
	"github.com/NicoNex/echotron/v3"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/makarychev13/wallet/internal/config"
	"github.com/makarychev13/wallet/internal/handler"
	"github.com/makarychev13/wallet/internal/handler/message"
	brokeragesdb "github.com/makarychev13/wallet/internal/storage/brokerages"
	depositsdb "github.com/makarychev13/wallet/internal/storage/deposits"
	"github.com/makarychev13/wallet/internal/usecase/brokerages"
	"github.com/makarychev13/wallet/internal/usecase/deposits"
	"github.com/makarychev13/wallet/pkg/state"
)

func main() {
	var cfg config.Config

	_, err := toml.DecodeFile("/Users/makarychev/Apps/github.com/makarychev13/wallet/config.toml", &cfg)
	if err != nil {
		slog.Error("can't parse config", "err", err)
		return
	}

	db, err := sqlx.Connect("postgres", cfg.Db.ConnectionString())
	if err != nil {
		slog.Error("can't connect ot db", "err", err)
		return
	}

	api := echotron.NewAPI(cfg.TgApiToken)

	storage := state.NewMemStorage[handler.State]()
	storage.Set(cfg.MyTgID, handler.InitState)

	sm := state.NewMachine(storage)

	accountsDb := brokeragesdb.New(db)
	depositsDb := depositsdb.New(db)

	accountsLister := brokerages.NewListUseCase(accountsDb)
	depositLister := deposits.NewListUseCase(depositsDb)

	reply := handler.New(api, accountsLister, depositLister)

	initState := state.New(handler.InitState)
	initState.On(message.Start, reply.Init)
	initState.On(message.BrokerageAccounts, reply.BrokerageAccounts)
	initState.On(message.Deposits, reply.ListDeposits)

	sm.Register(initState)

	for u := range echotron.PollingUpdates(cfg.TgApiToken) {
		if u.Message == nil {
			continue
		}
		if u.Message.From.ID != cfg.MyTgID {
			continue
		}

		err := sm.Handle(*u.Message)
		if err != nil {
			slog.Error("can't handle message", "err", err)
		}
	}
}
