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
	accountsdb "github.com/makarychev13/wallet/internal/storage/accounts"
	"github.com/makarychev13/wallet/internal/usecase/brokerages"
	"github.com/makarychev13/wallet/pkg/state"
)

func main() {
	var cfg config.Config

	_, err := toml.DecodeFile("config.toml", &cfg)
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

	accountsDb := accountsdb.New(db)
	accountsLister := brokerages.NewListUseCase(accountsDb)

	reply := handler.New(api, accountsLister)

	initState := state.New(handler.InitState)
	initState.On(message.Start, reply.Init)
	initState.On(message.BrokerageAccounts, reply.BrokerageAccounts)

	sm.Register(initState)

	for u := range echotron.PollingUpdates(cfg.TgApiToken) {
		err := sm.Handle(*u.Message)
		if err != nil {
			slog.Error("can't send message", "err", err)
		}
	}
}
