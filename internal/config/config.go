package config

import "fmt"

type Config struct {
	TgApiToken string `toml:"tg_api_token"`
	MyTgID     int64  `toml:"my_tg_id"`
	Db         Db     `toml:"db"`
}

type Db struct {
	User     string `toml:"user"`
	Name     string `toml:"name"`
	Password string `toml:"password"`
	SslMode  string `toml:"sslmode"`
}

func (d Db) ConnectionString() string {
	return fmt.Sprintf("user=%v dbname=%v password=%v sslmode=%v", d.User, d.Name, d.Password, d.SslMode)
}
